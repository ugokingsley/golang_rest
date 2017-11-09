package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


// A Response struct to map the Entire Response
type Response struct {
	Data    string    `json:"data"`
	//Newsletters []Newsletters `json:"newsletters"`
}

type Data struct {
	Newsletters []Newsletters `json:"newsletters"`
}


type Newsletters struct {
	//Newsletters []Newsletters `json:"newsletters"`
	Id string `json:"_id"`
	Issue string `json:"issue"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Type string `json:"type"`
	Label string `json:"label"`
	Url string `json:"url"`

}


func main() {
	response, err := http.Get("https://hakaselabs-newsletters.herokuapp.com/api/v1/st")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	
	fmt.Println(len(responseObject.Data))

	for i := 0; i < len(responseObject.Data); i++ {
		fmt.Println(responseObject.Data[i])
	}

}