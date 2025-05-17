package models

import "scarlet/internal/http/enum"

type BattleHistories struct {
	KillerNickname    string        `json:"knn"`
	RetainedCharNum   int64         `json:"rcn"`
	CharacterSkinType int           `json:"skn"`
	PlayerKill        int           `json:"pkc"`
	PlayerAddKill     int           `json:"pakc"`
	PlayerHitDamage   float32       `json:"tdtc"`
	Rank              int           `json:"rak"`
	PlayTimeSeconds   int64         `json:"pts"`
	BattleDtm         int64         `json:"bdt"`
	Chest             int           `json:"cht"`
	Leg               int           `json:"leg"`
	Head              int           `json:"hed"`
	Weapon            int           `json:"wep"`
	Trinket           int           `json:"trk"`
	Arm               int           `json:"arm"`
	Players           int           `json:"pla"`
	MonsterKill       int           `json:"mkc"`
	MonsterHitDamage  float32       `json:"tdtm"`
	Health            float32       `json:"hea"`
	Stamina           float32       `json:"sta"`
	Offence           float32       `json:"off"`
	Defence           float32       `json:"def"`
	GunFamiliarity    float32       `json:"gfa"`
	BladeFamiliarity  float32       `json:"bfa"`
	ThrowFamiliarity  float32       `json:"tfa"`
	PunchFamiliarity  float32       `json:"pfa"`
	BowFamiliarity    float32       `json:"wfa"`
	BluntFamiliarity  float32       `json:"lfa"`
	StabFamiliarity   float32       `json:"sfa"`
	CharacterLevel    int           `json:"clv"`
	InventoryItems    []int         `json:"inventoryItems"`
	GameMode          int           `json:"gmd"`
	PotentialSkill    int           `json:"psid"`
	KillerUnitType    enum.UnitType `json:"kut"`
	AssistCount       int           `json:"asc"`
	DeadCount         int           `json:"ddc"`
	RedTeamScore      int           `json:"rts"`
	BlueTeamScore     int           `json:"bts"`
	TeamNumber        int           `json:"tnm"`
	MainPerk          int           `json:"mpk"`
	FirstPerk         int           `json:"fpk"`
	SecondPerk        int           `json:"spk"`
}
