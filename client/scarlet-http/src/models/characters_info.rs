pub use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct CharactersInfo {
    #[serde(rename = "character")]
    pub character: Character,
    #[serde(rename = "characterSkins")]
    pub character_skins: Vec<CharacterSkin>,
    #[serde(rename = "characterSignatureList")]
    pub character_signature_list: Vec<CharacterSignature>,
    #[serde(rename = "characterDocument")]
    pub character_document: CharacterDocument,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Character {
    #[serde(rename = "cnm")]
    pub character_num: i32,
    #[serde(rename = "unm")]
    pub user_num: i64,
    #[serde(rename = "cls")]
    pub character_class: i32,
    #[serde(rename = "crd")]
    pub character_grade: i32,
    #[serde(rename = "ast")]
    pub active_character_skin_type: i32,
    #[serde(rename = "cpt")]
    pub character_purchase_type: i32,
    #[serde(rename = "unn")]
    pub user_nick_name: String,
    #[serde(rename = "ctt")]
    pub character_status: i32,
    #[serde(rename = "npc")]
    pub normal_play_count: i32,
    #[serde(rename = "nwc")]
    pub normal_win_count: i32,
    #[serde(rename = "psi")]
    pub potential_skill_id: i32,
    #[serde(rename = "l2d")]
    pub active_live2d: bool,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CharacterSkin {
    #[serde(rename = "unm")]
    pub user_num: i64,
    #[serde(rename = "cls")]
    pub character_class: i32,
    #[serde(rename = "ckt")]
    pub character_skin_type: i32,
    #[serde(rename = "live")]
    pub active_live2d: bool,
    #[serde(rename = "own")]
    pub owned: bool,
    #[serde(rename = "setp")]
    pub parent_skin: i32,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CharacterSignature {
    #[serde(rename = "cc")]
    pub character_class: i32,
    #[serde(rename = "ats")]
    pub signature_state_type: i32,
    #[serde(rename = "att")]
    pub signature_type: i32,
    #[serde(rename = "lv")]
    pub level: i32,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CharacterDocument {
    #[serde(rename = "unm")]
    pub user_num: i64,
    #[serde(rename = "cls")]
    pub character_class: i32,
    #[serde(rename = "ptn")]
    pub pattern: bool,
    #[serde(rename = "wks")]
    pub weakness: bool,
    #[serde(rename = "lik")]
    pub character_like_reward: bool,
    #[serde(rename = "itr")]
    pub interview_reward: bool,
}
