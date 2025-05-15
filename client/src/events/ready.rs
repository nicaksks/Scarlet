use log::info;
use serenity::{
    all::{ActivityData, Context, EventHandler, OnlineStatus, Ready},
    async_trait,
};

pub struct ReadyHandler;

#[async_trait]
impl EventHandler for ReadyHandler {
    async fn ready(&self, ctx: Context, ready: Ready) {

        info!("{} online!", ready.user.name);
        info!("{} is on {} guild(s)!", ready.user.name, ready.guilds.len());

        ctx.set_presence(
            Some(ActivityData::custom("Black Survival: Project Lumia")),
            OnlineStatus::DoNotDisturb,
        );
    }
}
