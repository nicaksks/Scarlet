package controller

import (
	"fmt"
	scarlet "scarlet/internal/http"
	"scarlet/internal/http/enum"
	"scarlet/internal/http/helpers"
	"scarlet/internal/http/models"
	"scarlet/internal/redis"
)

func User(region enum.Region, userNum string) (models.BSResponse[models.User], error) {

	cacheKey := fmt.Sprintf("%s_%s", "user", userNum)
	cached, err := redis.Get[models.User](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.User]{
			Region:   region,
			UserNum:  userNum,
			Endpoint: helpers.FormatEndpoint(enum.USERPROFILE, userNum),
			CacheKey: cacheKey,
		}

		response, err := client.Post(helpers.DefaultBody())
		return response, err
	}

	return cached, nil
}

func Games(region enum.Region, userNum string) (models.BSResponse[models.Games], error) {

	cacheKey := fmt.Sprintf("%s_%s", "games", userNum)
	cached, err := redis.Get[models.Games](cacheKey)

	if err != nil {

		client := scarlet.HTTP[models.Games]{
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
