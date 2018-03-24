package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// API to get symnonym of words

func main() {
	http.HandleFunc("/", getWordSymnonym{})
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

// Function to handle / endpoint
func getWordSymnonym(ww http.ResponseWriter, r *http.Request) {
	BIG_HUGE_LABS_URL = "http://words.bighugelabs.com/api/2/6d1907189a886a851caba6c96ea972db/%s/json"
	paramQArray := (r.URL.Query()["q"])

	if len(paramQArray) > 0 {

		// Inject paramQ into API url
		paramQ := paramQArray[0]
		symnonymUrl := fmt.Sprintf(BIG_HUGE_LABS_URL, paramQ)

		// Send and get response from Big Huge Labs (word symnonyms)
		req, err := http.NewRequest("GET", symnonymUrl, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
		}

		defer resp.Body.Close()

		// Return to client list of symnonym of asked word
		// In the full version, this should have the ability
		// to further parse word from raw input (example.com => [example, com])
		// Plus the format should have been better prepared
		// (I am just returning the response from Big Huge Labs
		// for now instead of marshalling it)
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				log.Fatal(err2)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Symnonyms for '%s' : \n", paramQ)
			fmt.Fprintf(w, bodyString)
		}
	}
}

// Server port
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
