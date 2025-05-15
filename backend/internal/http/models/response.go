package models

type BSResponse[T any] struct {
	Code           int    `json:"Cod"`
	Message        string `json:"Msg"`
	CacheExpiresAt int64  `json:"Cea"`
	Result         T      `json:"Rst"`
}

type User struct {
	User Profile `json:"User"`
}

type Histories struct {
	NewRequestArrived bool              `json:"newRequestArrived"`
	BattleHistories   []BattleHistories `json:"battleHistories"`
}

type PlayCount struct {
	CharacterPlayCount MostPlayedCharacter `json:"characterPlayCount"`
}

type CharactersInfo struct {
	Character              Character            `json:"character"`
	CharacterSkins         []CharacterSkin      `json:"characterSkins"`
	CharacterSignatureList []CharacterSignature `json:"characterSignatureList"`
	CharacterDocument      CharacterDocument    `json:"characterDocument"`
}
