package helpers

import (
	"scarlet/internal/http/enum"
	"scarlet/internal/http/models"
	"strconv"

	"encoding/base64"
	"encoding/json"

	"github.com/google/uuid"
)

func payload(userNum string) models.Token {
	sessionKey := uuid.New()
	return models.Token{
		ShardNum:           0,
		AuthProvider:       enum.GUEST,
		BuildPlatform:      enum.NN,
		OsType:             enum.ETC,
		OsVersion:          "Scarlet",
		AppVersion:         "0.1.0",
		Guest:              true,
		SessionKey:         sessionKey,
		DeviceLangCode:     "en",
		AppLangCode:        "en",
		Market:             enum.STEAM,
		Attribution:        "DISCORD",
		DeviceLocationCode: "us",
		UserNum:            validateUserNum(userNum),
	}
}

func validateUserNum(userNum string) int {
	if un, err := strconv.Atoi(userNum); err == nil {
		return un
	}

	return int(enum.DEFAULT)
}

func Token(userNum string) string {
	if serialize, err := json.Marshal(payload(userNum)); err == nil {
		jwt := base64.StdEncoding.EncodeToString(serialize)
		return jwt
	}
	return ""
}
