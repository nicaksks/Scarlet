mod client;
pub mod models;

use crate::client::Client;
use crate::models::histories::Histories;
use crate::models::profile::User;
use crate::models::most_played_character::MostPlayedCharacter;
use crate::models::characters_info::CharactersInfo;
pub struct Scarlet;

impl Scarlet {
    pub async fn profile(&self, user_num: String) -> Result<User, String> {
        Client.get::<User>(format!("/profile/{}", user_num)).await
    }

    pub async fn histories(&self, user_num: String) -> Result<Histories, String> {
        Client
            .get::<Histories>(format!("/profile/{}/histories", user_num))
            .await
    }

    pub async fn most_played_character(&self, user_num: String) -> Result<MostPlayedCharacter, String> {
        Client
            .get::<MostPlayedCharacter>(format!("/profile/{}/mostPlayedCharacter", user_num))
            .await
    }

    pub async fn character_info(
        &self,
        user_num: String,
        character_num: String,
    ) -> Result<CharactersInfo, String> {
        Client
            .get::<CharactersInfo>(format!("/profile/{}/character/{}", user_num, character_num))
            .await
    }
}
