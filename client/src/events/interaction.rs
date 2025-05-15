use crate::commands::profile;

#[derive(Debug)]
pub struct Data {}

pub fn register() -> poise::Framework<Data, Box<dyn std::error::Error + Send + Sync>> {
    let commands = vec![profile::profile()];

    poise::Framework::builder()
        .options(poise::FrameworkOptions {
            commands: commands,
            ..Default::default()
        })
        .setup(|ctx, _, framework| {
            Box::pin(async move {
                poise::builtins::register_globally(ctx, &framework.options().commands).await?;
                Ok(Data {})
            })
        })
        .build()
}
