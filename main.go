package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	libraKnife()
	fmt.Println("initialized")
	go func() {
		for {
			fmt.Println("Entrou dentro do for ")
			time.Sleep(5 * time.Second)
			value := getLibraCurrency()
			print(value)
			if value < 6.80 {
				fmt.Println("awaiting the time")
				sendtext("valor diario libra", fmt.Sprintf("%.2f", value))
				fmt.Println("valor enviado")
			}
		}
	}()

	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getLibraCurrency() float64 {
	const site = "https://economia.awesomeapi.com.br/json/last/GBP-BRL"
	resp, err := http.Get(site)
	if err != nil {
		log.Fatalf("erro")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal((body), &result)
	//fmt.Println(result)

	exampleMap := result.(map[string]interface{})

	tickerMap := exampleMap["GBPBRL"].(map[string]interface{})

	value := tickerMap["bid"].(string)

	valuefloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println(err)
	}
	return valuefloat
}

func getUrl() string {
	Token := "5790558519:AAHVftlm--noZFITVCID9c1JywGSJNKxT1M"
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func sendtext(message string, text interface{}) error {
	ChatId := "-751952371"
	var err error
	var response *http.Response

	//Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, err := json.Marshal(map[string]string{
		"chat_id": ChatId,
		"text":    message + " " + text.(string),
	})
	if err != nil {
		return err
	}

	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	fmt.Println(response)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}

func libraKnife() {
	libraByte, _ := ioutil.ReadFile("libra")
	fmt.Println(string(libraByte))
}

//func getVariable() {
//
//	log.Println("[VIPER] Loading context")
//
//	configName := "application"
//
//	viperSetup := viper.GetViper()
//	viperSetup.SetConfigType("yml")
//	viperSetup.SetConfigName(configName)
//	viper.AddConfigPath(".")
//	viperSetup.AllowEmptyEnv(true)
//	viperSetup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
//	viperSetup.AutomaticEnv()
//	err := viper.ReadInConfig()
//
//	if err != nil {
//		panic(err)
//	}
//}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
