mod client;
pub mod models;

use crate::client::Client;
use crate::models::games::Games;
use crate::models::user::User;
use crate::models::most_played_character::MostPlayedCharacter;
use crate::models::characters_info::CharactersInfo;
pub struct Scarlet;

impl Scarlet {
    pub async fn profile(&self, user_num: String) -> Result<User, String> {
        Client.get::<User>(format!("/user/{}", user_num)).await
    }

    pub async fn games(&self, user_num: String) -> Result<Games, String> {
        Client
            .get::<Games>(format!("/user/{}/games", user_num))
            .await
    }

    pub async fn most_played_character(&self, user_num: String) -> Result<MostPlayedCharacter, String> {
        Client
            .get::<MostPlayedCharacter>(format!("/user/{}/mostPlayedCharacter", user_num))
            .await
    }

    pub async fn character_info(
        &self,
        user_num: String,
        character_num: String,
    ) -> Result<CharactersInfo, String> {
        Client
            .get::<CharactersInfo>(format!("/user/{}/character/{}", user_num, character_num))
            .await
    }
}
