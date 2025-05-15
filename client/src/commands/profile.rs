use crate::{commands::Context, commands::Error, components::embeds::profile};

use poise::CreateReply;
use scarlet::Scarlet;

/// Show user Profile
#[poise::command(slash_command)]
pub async fn profile(
    ctx: Context<'_>,
    #[description = "Account ID (Discord ID)"] user: Option<String>,
) -> Result<(), Error> {
    ctx.defer().await?;

    let user_num = user.unwrap_or(ctx.author().id.to_string());
    let data = Scarlet.profile(user_num.clone()).await;

    if let Err(err) = data {
        ctx.reply(err.clone()).await?;
        return Err(err.into());
    }

    let most_played_character = Scarlet.most_played_character(user_num).await;

    ctx.send(CreateReply::default().embed(profile::EmbedProfile::new(
        data.unwrap(),
        most_played_character.unwrap(),
    )))
    .await?;

    Ok(())
}
