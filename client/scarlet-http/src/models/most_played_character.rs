use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct MostPlayedCharacter {
    #[serde(rename = "characterPlayCount")]
    pub character_play_count: CharacterPlayCount,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CharacterPlayCount {
    #[serde(rename = "t")]
    pub total_play_count: i32,
    #[serde(rename = "m")]
    pub most_played_character_count: i32,
    #[serde(rename = "c")]
    pub most_played_character: i32,
    #[serde(rename = "w")]
    pub win_count: i32,
}
