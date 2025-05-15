package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"scarlet/internal/http/enum"
	"scarlet/internal/http/helpers"
	"scarlet/internal/http/models"
	"scarlet/internal/redis"
	"time"
)

var clientSettings = &http.Client{}

type HTTP[T any] struct {
	Region   enum.Region
	CacheKey string
	Endpoint string
	UserNum  string
}

func (h *HTTP[T]) url() string {
	return fmt.Sprintf("%s://%s.%s:%d/%s%s", enum.PROTOCOL, h.Region, enum.DOMAIN, enum.REST, enum.PATH, h.Endpoint)
}

func (h *HTTP[T]) payload(body []io.Reader) io.Reader {
	if len(body) > 0 {
		return body[0]
	}

	return nil
}

func (h *HTTP[T]) setCacheTime(minutes uint8) int64 {
	timeNow := time.Now().Local()
	return timeNow.Add(time.Minute * time.Duration(minutes)).UnixMilli()
}

func (h *HTTP[T]) client(method enum.Methods, body ...io.Reader) *http.Response {

	log.Println("[CLIENT - URL] -> ", h.url())

	request, err := http.NewRequest(string(method), h.url(), h.payload(body))
	if err != nil {
		// TODO: Add error handling
		fmt.Println("TODO")
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Bs-Session", helpers.Token(h.UserNum))

	response, err := clientSettings.Do(request)
	if err != nil {
		// TODO: Add error handling
		fmt.Println("TODO")
		fmt.Println(err)
	}

	return response
}

func (h *HTTP[T]) getJSON(response *http.Response) (models.BSResponse[T], error) {

	var data models.BSResponse[T]

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println("[HTTP Client] -> StatusCode: ", response.StatusCode)
		log.Println("[Error] -> Player not found.")
		return data, helpers.ErrNotFound
	}

	err := json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		log.Println("[Decoder - Error] -> ", err)
		return data, helpers.ErrUnkown
	}

	if data.Code != 200 {
		log.Printf("[BSPL] Error -> %s\n", data.Message)
		return data, errors.New(data.Message)
	}

	log.Println("[Response] -> ", data)

	data.CacheExpiresAt = h.setCacheTime(30)
	redis.Set(h.CacheKey, data)
	return data, nil
}

func (h *HTTP[T]) Get() (models.BSResponse[T], error) {
	request := h.client(enum.GET)
	return h.getJSON(request)
}

func (h *HTTP[T]) Post(body io.Reader) (models.BSResponse[T], error) {
	request := h.client(enum.POST, body)
	return h.getJSON(request)
}
