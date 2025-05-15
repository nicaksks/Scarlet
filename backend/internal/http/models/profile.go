package models

import "scarlet/internal/http/enum"

type Profile struct {
	ActivatedPotentialSkillId int         `json:"activatedPotentialSkillId"`
	UserNum                   int         `json:"unm"`
	ReceiverPushMsg           bool        `json:"rpm"`
	NewPostArrived            bool        `json:"npa"`
	TermsAgree                bool        `json:"tag"`
	NickName                  string      `json:"nnm"`
	BackGroundMusic           bool        `json:"bgm"`
	SoundEffect               bool        `json:"sef"`
	LastWord                  string      `json:"lwd"`
	WatchWord                 string      `json:"wwd"`
	ActiveCharacterNum        int         `json:"acn"`
	AppLanguageCode           string      `json:"alc"`
	AdViewCount               int         `json:"rvn"`
	League                    enum.League `json:"leg"`
	Aps                       int         `json:"aps"`
	MaxAdViewCount            int         `json:"mrn"`
	Rtn                       string      `json:"rtn"`
	BearPoint                 int         `json:"bpt"`
	Gold                      int         `json:"gld"`
	Gem                       int         `json:"gem"`
	Credit                    int         `json:"cd"`
	Mileage                   int         `json:"ma"`
	ExperiementMemory         int         `json:"em"`
	TournamentPoint           int         `json:"tp"`
	TournamentTicket          int         `json:"tt"`
	VoteTicket                int         `json:"vt"`
	VoteStamp                 int         `json:"vs"`
	LabyrinthPoint            int         `json:"lp"`
	EventCoin                 int         `json:"hc"`
}
