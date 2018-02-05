package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestGetBookList(t *testing.T) {
	ts := createTestEnv(t)

	res, err := http.Get(ts.URL + "/book")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if len(body) == 0 {
		t.Error("body cannot be empty")
	}
}
