package models

import "scarlet/internal/http/enum"

type Character struct {
	CharacterNum            int                        `json:"cnm"`
	UserNum                 int64                      `json:"unm"`
	CharacterClass          int                        `json:"cls"`
	CharacterGrade          enum.CharacterGrade        `json:"crd"`
	ActiveCharacterSkinType int                        `json:"ast"`
	CharacterPurchaseType   enum.CharacterPurchaseType `json:"cpt"`
	UserNickName            string                     `json:"unn"`
	CharacterStatus         enum.CharacterStatus       `json:"ctt"`
	NormalPlayCount         int                        `json:"npc"`
	NormalWinCount          int                        `json:"nwc"`
	PotentialSkillId        int                        `json:"psi"`
	ActiveLive2D            bool                       `json:"l2d"`
}

type CharacterSkin struct {
	UserNum           int64 `json:"unm"`
	CharacterClass    int   `json:"cls"`
	CharacterSkinType int   `json:"ckt"`
	ActiveLive2D      bool  `json:"live"`
	Owned             bool  `json:"own"`
	ParentSkin        int   `json:"setp"`
}

type CharacterSignature struct {
	CharacterClass     int                     `json:"cc"`
	SignatureStateType enum.SignatureStateType `json:"ats"`
	SignatureType      enum.SignatureType      `json:"att"`
	Level              int                     `json:"lv"`
}

type CharacterDocument struct {
	UserNum             int64 `json:"unm"`
	CharacterClass      int   `json:"cls"`
	Pattern             bool  `json:"ptn"`
	Weakness            bool  `json:"wks"`
	CharacterLikeReward bool  `json:"lik"`
	InterviewReward     bool  `json:"itr"`
}
