package models

import (
	"scarlet/internal/http/enum"

	"github.com/google/uuid"
)

type Token struct {
	ShardNum           int                `json:"snm"`
	AuthProvider       enum.AuthProvider  `json:"atp"`
	BuildPlatform      enum.BuildPlatform `json:"bpf"`
	OsType             enum.OsType        `json:"otp"`
	OsVersion          string             `json:"ovs"`
	AppVersion         string             `json:"avs"`
	Guest              bool               `json:"gst"`
	SessionKey         uuid.UUID          `json:"snk"`
	DeviceLangCode     string             `json:"dlc"`
	AppLangCode        string             `json:"alc"`
	Market             enum.Market        `json:"mkt"`
	Attribution        string             `json:"atb"`
	DeviceLocationCode string             `json:"dloc"`
	UserNum            int                `json:"unm"` // DiscordId
}
