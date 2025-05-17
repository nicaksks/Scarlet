use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct User {
    #[serde(rename = "User")]
    pub user: Profile,
}

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct Profile {
    #[serde(rename = "activatedPotentialSkillId")]
    pub activated_potential_skill_id: i64,

    #[serde(rename = "unm")]
    pub user_num: i64,

    #[serde(rename = "rpm")]
    pub receiver_push_msg: bool,

    #[serde(rename = "npa")]
    pub new_post_arrived: bool,

    #[serde(rename = "tag")]
    pub terms_agree: bool,

    #[serde(rename = "nnm")]
    pub nick_name: String,

    #[serde(rename = "bgm")]
    pub back_ground_music: bool,

    #[serde(rename = "sef")]
    pub sound_effect: bool,

    #[serde(rename = "lwd")]
    pub last_word: String,

    #[serde(rename = "wwd")]
    pub watch_word: String,

    #[serde(rename = "acn")]
    pub active_character_num: i64,

    #[serde(rename = "alc")]
    pub app_language_code: String,

    #[serde(rename = "rvn")]
    pub ad_view_count: i64,

    #[serde(rename = "leg")]
    pub league: i64,

    #[serde(rename = "aps")]
    pub aps: i64,

    #[serde(rename = "mrn")]
    pub max_ad_view_count: i64,

    #[serde(rename = "rtn")]
    pub rtn: String,

    #[serde(rename = "bpt")]
    pub bear_point: i64,

    #[serde(rename = "gld")]
    pub gold: i64,

    #[serde(rename = "gem")]
    pub gem: i64,

    #[serde(rename = "cd")]
    pub credit: i64,

    #[serde(rename = "ma")]
    pub mileage: i64,

    #[serde(rename = "em")]
    pub experiement_memory: i64,

    #[serde(rename = "tp")]
    pub tournament_point: i64,

    #[serde(rename = "tt")]
    pub tournament_ticket: i64,

    #[serde(rename = "vt")]
    pub vote_ticket: i64,

    #[serde(rename = "vs")]
    pub vote_stamp: i64,

    #[serde(rename = "lp")]
    pub labyrinth_point: i64,

    #[serde(rename = "hc")]
    pub event_coin: i64,
}
