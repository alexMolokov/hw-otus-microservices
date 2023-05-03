package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexMolokov/hw-otus-microservices/internal/model"
	"github.com/bxcodec/faker/v3"
)

var (
	host              string
	port              string
	duration          int
	rps               int
	layout            = "2006-01-02 15:04:05"
	hostAPI           string
	httpAPI           string
	httpAPICreateUser string
	httpAPIError      string
)

func init() {
	flag.StringVar(&host, "host", "arch.homework", "Host for request")
	flag.IntVar(&duration, "duration", 60, "Load duration in seconds")
	flag.IntVar(&rps, "rps", 20, "rps")
	flag.StringVar(&port, "port", "80", "port")
}

// Нагрузка для api ...
func main() {
	flag.Parse()
	tCh := time.After(time.Duration(duration) * time.Second)
	ids := make(chan int)
	defer func() {
		close(ids)
	}()

	hostAPI = fmt.Sprintf("http://%s:%s", host, port)
	httpAPICreateUser = hostAPI + "/api/v1/user"
	httpAPI = hostAPI + "/api/v1/user/{id}"
	httpAPIError = hostAPI + "/sometimes/error"

	rand.Seed(time.Now().UnixNano())

	go func() {
		client := &http.Client{}
		timeout := 1000 / rps
		for {
			r := userCreateRequest()
			jsonData, _ := json.Marshal(r)
			req, _ := http.NewRequestWithContext(context.Background(), "POST",
				httpAPICreateUser, bytes.NewBuffer(jsonData))
			req.Header.Set("contentType", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			if resp.StatusCode == http.StatusOK {
				body, _ := ioutil.ReadAll(resp.Body)
				response := model.UserCreateResponse{}
				err := json.Unmarshal(body, &response)
				if err == nil {
					ids <- int(response.ID)
				}
			}
			_ = resp.Body.Close()
			time.Sleep(time.Duration(timeout) * time.Millisecond)
		}
	}()

	go func() {
		client := &http.Client{}
		for id := range ids {
			r := userUpdateRequest(id)
			jsonData, _ := json.Marshal(r)
			req, _ := http.NewRequestWithContext(context.Background(), "PUT",
				getURL(id), bytes.NewBuffer(jsonData))
			req.Header.Set("contentType", "application/json")
			resp, _ := client.Do(req)
			_ = resp.Body.Close()
			s := rand.Intn(500) // nolint
			time.Sleep(time.Duration(s) * time.Millisecond)
		}
	}()

	go func() {
		client := &http.Client{}
		for id := range ids {
			req, _ := http.NewRequestWithContext(context.Background(), "GET",
				getURL(id), nil)
			req.Header.Set("contentType", "application/json")
			resp, _ := client.Do(req)
			_ = resp.Body.Close()
		}
	}()

	go func() {
		client := &http.Client{}
		for id := range ids {
			req, _ := http.NewRequestWithContext(context.Background(), "DELETE",
				getURL(id), nil)
			req.Header.Set("contentType", "application/json")
			resp, _ := client.Do(req)
			_ = resp.Body.Close()
			s := rand.Intn(500) // nolint
			time.Sleep(time.Duration(s) * time.Millisecond)
		}
	}()

	go func() {
		client := &http.Client{}
		for id := range ids {
			req, _ := http.NewRequestWithContext(context.Background(), "DELETE",
				getURL(id), nil)
			req.Header.Set("contentType", "application/json")
			resp, _ := client.Do(req)
			_ = resp.Body.Close()
			s := rand.Intn(300) // nolint
			time.Sleep(time.Duration(s) * time.Millisecond)
		}
	}()

	go func() {
		client := &http.Client{}
		for range ids {
			req, _ := http.NewRequestWithContext(context.Background(), "GET",
				httpAPIError, nil)
			req.Header.Set("contentType", "application/json")
			resp, _ := client.Do(req)
			_ = resp.Body.Close()
			s := rand.Intn(900) // nolint
			time.Sleep(time.Duration(s) * time.Millisecond)
		}
	}()

	<-tCh
}

func userCreateRequest() model.UserCreateRequest {
	return model.UserCreateRequest{
		UserName: faker.Name() + time.Now().Format(layout),
		UserCommon: model.UserCommon{
			FirstName: getStringPointer(faker.Name() + time.Now().Format(layout)),
			LastName:  getStringPointer(faker.Name() + time.Now().Format(layout)),
			Email:     getStringPointer(faker.Email()),
			Phone:     getStringPointer(faker.Phonenumber()),
		},
	}
}

func userUpdateRequest(id int) model.UserUpdateRequest {
	return model.UserUpdateRequest{
		UserID: int64(id),
		UserCommon: model.UserCommon{
			FirstName: getStringPointer(faker.Name() + time.Now().Format(layout)),
			LastName:  getStringPointer(faker.Name() + time.Now().Format(layout)),
			Email:     getStringPointer(faker.Email()),
			Phone:     getStringPointer(faker.Phonenumber()),
		},
	}
}

func getStringPointer(str string) *string {
	return &str
}

func getURL(id int) string {
	sID := strconv.Itoa(id)
	return strings.ReplaceAll(httpAPI, "{id}", sID)
}
