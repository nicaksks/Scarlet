use serenity::all::{Client, GatewayIntents};
use std::env;

use crate::events::{interaction, ready};

fn token() -> String {
    if let Ok(token) = env::var("DISCORD_TOKEN") {
        return token;
    }

    panic!("[ERROR] -> Missing Discord Token in ENV file.");
}

fn intents() -> GatewayIntents {
    GatewayIntents::GUILD_MESSAGES | GatewayIntents::MESSAGE_CONTENT
}

pub async fn scarlet() -> Client {
    Client::builder(&token(), intents())
        .framework(interaction::register())
        .event_handler(ready::ReadyHandler)
        .await
        .expect("[ERROR] -> creating client")
}
