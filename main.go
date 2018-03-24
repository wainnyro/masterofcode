package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paramQArray := (r.URL.Query()["q"])
		if len(paramQArray) > 0 {

			paramQ := paramQArray[0]
			fmt.Fprintf(w, "Symnonyms for '%s' : \n", paramQ)

			url := fmt.Sprintf("http://words.bighugelabs.com/api/2/6d1907189a886a851caba6c96ea972db/%s/json", paramQ)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal("NewRequest: ", err)
				return
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal("Do: ", err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				bodyBytes, err2 := ioutil.ReadAll(resp.Body)
				if err2 != nil {
					log.Fatal(err2)
				}
				bodyString := string(bodyBytes)
				fmt.Println(bodyString)
				fmt.Fprintf(w, bodyString)
			}
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
