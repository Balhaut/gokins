package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type test_struct struct {
	Push struct {
		Changes []struct {
			New struct {
				Branch string `json:"name"`
			} `json:"new"`
		} `json:"changes"`
	} `json:"push"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I GOED %s WEE!",
		html.EscapeString(r.URL.Path))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
	}
	log.Println(string(body))
	fmt.Fprintf(w, "RESPONSE BODY:\n%s", body)

	var t test_struct
	err = json.Unmarshal(body, &t)

	if err != nil {
		panic(err.Error())
		log.Fatal(err)
	}
	fmt.Fprintf(w, "BRANCH NAME = '%s'", t.Push.Changes[0].New.Branch)
	fmt.Fprintf(w, "REPO NAME = '%s'", t.Repository.Name)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	var out io.Writer = ioutil.Discard
	out = os.Stdout
	logger := log.New(out, "[gokins]", log.Lshortfile)
	logger.Fatal(http.ListenAndServe(":9001", router))
}
