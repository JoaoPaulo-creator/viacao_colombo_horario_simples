package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Config struct {
	SID   string
	Token string
	From  string
	To    string
}

func (c *Config) setToken(token string) {
	c.Token = token
}

func (c *Config) setSID(sid string) {
	c.SID = sid
}

func (c *Config) getToken() string {
	return c.Token
}

func (c *Config) getSID() string {
	return c.SID
}

var c Config

func load() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

	sid := os.Getenv("SID")
	authToken := os.Getenv("AUTH_TOKEN")
	from := os.Getenv("FROM_NUMBER")
	to := os.Getenv("TO_NUMBER")

	c = Config{
		From: from,
		To:   to,
	}

	c.setSID(sid)
	c.setToken(authToken)
	c.setSID(sid)
	c.setToken(authToken)

	return c, nil
}

func start(content string) *twilioApi.CreateMessageParams {
	cfg, err := load()
	if err != nil {
		log.Fatalf("ERRO: %v", err)
		// TODO: ver uma forma de retornar uma mensagem erro decente
		// return fmt.Sprintf("Ocorreu um erro ao tentar carregar variáveis de ambiente")
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetFrom(cfg.From)
	params.SetTo(cfg.To)
	params.SetBody(content)

	return params
}

var errorCode = []uint{}

func PerformRequest(content string) string {
	p := start(content)
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.getSID(),
		Password: c.getToken(),
	})

	resp, err := client.Api.CreateMessage(p)
	if err != nil {
		fmt.Println(err)
		return fmt.Sprint("Ocorreu um eror ao enviar mensagem para cliente\n")
	}

	response, _ := json.Marshal(*resp)
	return string(response)
}
