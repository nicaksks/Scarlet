use dotenv::dotenv;
use log::debug;
use scarlet::scarlet;

mod commands;
mod components;
mod constants;
mod events;
mod scarlet;

#[tokio::main]
async fn main() {
    env_logger::init();
    dotenv().ok();

    let mut client = scarlet().await;

    if let Err(why) = client.start().await {
        debug!("[ERROR] -> {:?}", why)
    }
}
