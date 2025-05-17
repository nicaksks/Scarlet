pub mod games;
pub mod user;
pub mod most_played_character;
pub mod characters_info;

pub mod response {

    use serde::{Deserialize, Serialize};

    #[derive(Serialize, Deserialize, Debug)]
    pub struct BsResponse<T> {
        #[serde(rename = "Cod")]
        pub code: i32,

        #[serde(rename = "Msg")]
        pub message: String,

        #[serde(rename = "Cea")]
        pub cache_expires_at: Option<i64>,

        #[serde(rename = "Rst")]
        pub result: Option<T>,
    }
}
