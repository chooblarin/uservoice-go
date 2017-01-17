package main

import (
	"fmt"
	"log"
	"os"

	"io/ioutil"

	"encoding/json"

	"github.com/chooblarin/uservoice-go"
)

func loadCredentials() (subdomain, apiKey, apiSecret string) {
	file, err := os.Open("credentials.json")
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(file)
	cred := &struct {
		SubDomain string `json:"subdomain"`
		APIKey    string `json:"apiKey"`
		APISecret string `json:"apiSecret"`
	}{}
	decoder.Decode(cred)
	return cred.SubDomain, cred.APIKey, cred.APISecret
}

func main() {
	subdomain, apiKey, apiSecret := loadCredentials()

	client, err := uservoice.NewClient(subdomain, apiKey, apiSecret)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Request("GET", "/api/v1/tickets.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
