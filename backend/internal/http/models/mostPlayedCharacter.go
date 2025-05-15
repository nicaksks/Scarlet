package models

type MostPlayedCharacter struct {
	TotalPlayCount           int `json:"t"`
	MostPlayedCharacterCount int `json:"m"`
	MostPlayedCharacter      int `json:"c"`
	WinCount                 int `json:"w"`
}
