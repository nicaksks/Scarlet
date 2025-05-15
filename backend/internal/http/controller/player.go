package controller

import (
	"fmt"
	scarlet "scarlet/internal/http"
	"scarlet/internal/http/enum"
	"scarlet/internal/http/helpers"
	"scarlet/internal/http/models"
	"scarlet/internal/redis"
)

func Profile(region enum.Region, userNum string) (models.BSResponse[models.User], error) {

	cacheKey := fmt.Sprintf("%s_%s", "profile", userNum)
	cached, err := redis.Get[models.User](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.User]{
			Region:   region,
			UserNum:  userNum,
			Endpoint: helpers.FormatEndpoint(enum.PROFILE, userNum),
			CacheKey: cacheKey,
		}

		response, err := client.Post(helpers.DefaultBody())
		return response, err
	}

	return cached, nil
}

func Histories(region enum.Region, userNum string) (models.BSResponse[models.Histories], error) {

	cacheKey := fmt.Sprintf("%s_%s", "histories", userNum)
	cached, err := redis.Get[models.Histories](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.Histories]{
			Region:   region,
			UserNum:  userNum,
			Endpoint: helpers.FormatEndpoint(enum.HISTORIES, userNum),
			CacheKey: cacheKey,
		}

		response, err := client.Get()
		return response, err
	}

	return cached, nil
}

func MostPlayedCharacter(region enum.Region, userNum string) (models.BSResponse[models.PlayCount], error) {

	cacheKey := fmt.Sprintf("%s_%s", "most_played_character", userNum)
	cached, err := redis.Get[models.PlayCount](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.PlayCount]{
			Region:   region,
			UserNum:  userNum,
			Endpoint: string(enum.PLAYCOUNT),
			CacheKey: cacheKey,
		}

		response, err := client.Get()
		return response, err
	}

	return cached, nil
}

func CharacterInfo(region enum.Region, userNum string, characterNum string) (models.BSResponse[models.CharactersInfo], error) {

	cacheKey := fmt.Sprintf("%s_%s", "character", characterNum)
	cached, err := redis.Get[models.CharactersInfo](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.CharactersInfo]{
			Region:   region,
			UserNum:  userNum,
			Endpoint: helpers.FormatEndpoint(enum.CHARACTERSINFO, characterNum),
			CacheKey: cacheKey,
		}

		response, err := client.Get()
		return response, err
	}

	return cached, nil
}
