package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", func(c echo.Context) error {
		m, err := unmarshal(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		data := string(m.Message.Data)
		log.Printf("Subscriver_Received_before:%s", data)
		time.Sleep(time.Second * 5)
		log.Printf("Subscriver_Received_after:%s", data)
		return c.String(http.StatusOK, "OK")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func unmarshal(r io.ReadCloser) (*PubSubMessage, error) {
	var m PubSubMessage
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		return nil, err
	}
	return &m, nil
}
