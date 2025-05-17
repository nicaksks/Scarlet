use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Games {
    #[serde(rename = "newRequestArrived")]
    new_request_arrived: bool,
    #[serde(rename = "battleHistories")]
    battle_histories: Vec<BattleHistories>,
}

#[derive(Serialize, Deserialize, Clone, Debug)]
pub enum UnitType {
    Character = 1,
    Monster = 2,
    EventObject = 3,
}

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct BattleHistories {
    #[serde(rename = "knn")]
    pub killer_nickname: String,
    #[serde(rename = "rcn")]
    pub retained_char_num: i64,
    #[serde(rename = "skn")]
    pub character_skin_type: i32,
    #[serde(rename = "pkc")]
    pub player_kill: i32,
    #[serde(rename = "pakc")]
    pub player_add_kill: i32,
    #[serde(rename = "tdtc")]
    pub player_hit_damage: f32,
    #[serde(rename = "rak")]
    pub rank: i32,
    #[serde(rename = "pts")]
    pub play_time_seconds: i64,
    #[serde(rename = "bdt")]
    pub battle_dtm: i64,
    #[serde(rename = "cht")]
    pub chest: i32,
    #[serde(rename = "leg")]
    pub leg: i32,
    #[serde(rename = "hed")]
    pub head: i32,
    #[serde(rename = "wep")]
    pub weapon: i32,
    #[serde(rename = "trk")]
    pub trinket: i32,
    #[serde(rename = "arm")]
    pub arm: i32,
    #[serde(rename = "pla")]
    pub players: i32,
    #[serde(rename = "mkc")]
    pub monster_kill: i32,
    #[serde(rename = "tdtm")]
    pub monster_hit_damage: f32,
    #[serde(rename = "hea")]
    pub health: f32,
    #[serde(rename = "sta")]
    pub stamina: f32,
    #[serde(rename = "off")]
    pub offence: f32,
    #[serde(rename = "def")]
    pub defence: f32,
    #[serde(rename = "gfa")]
    pub gun_familiarity: f32,
    #[serde(rename = "bfa")]
    pub blade_familiarity: f32,
    #[serde(rename = "tfa")]
    pub throw_familiarity: f32,
    #[serde(rename = "pfa")]
    pub punch_familiarity: f32,
    #[serde(rename = "wfa")]
    pub bow_familiarity: f32,
    #[serde(rename = "lfa")]
    pub blunt_familiarity: f32,
    #[serde(rename = "sfa")]
    pub stab_familiarity: f32,
    #[serde(rename = "clv")]
    pub character_level: i32,
    #[serde(rename = "inventoryItems")]
    pub inventory_items: Vec<i32>,
    #[serde(rename = "gmd")]
    pub game_mode: i32,
    #[serde(rename = "psid")]
    pub potential_skill: i32,
    #[serde(rename = "kut")]
    pub killer_unit_type: i32, //UnitType
    #[serde(rename = "asc")]
    pub assist_count: i32,
    #[serde(rename = "ddc")]
    pub dead_count: i32,
    #[serde(rename = "rts")]
    pub red_team_score: i32,
    #[serde(rename = "bts")]
    pub blue_team_score: i32,
    #[serde(rename = "tnm")]
    pub team_number: i32,
    #[serde(rename = "mpk")]
    pub main_perk: i32,
    #[serde(rename = "fpk")]
    pub first_perk: i32,
    #[serde(rename = "spk")]
    pub second_perk: i32,
}
