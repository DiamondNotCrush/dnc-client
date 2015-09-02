package router_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dnc/dnc-client/router"
)

var (
	server     *httptest.Server
	reader     io.Reader
	libraryUrl string
	sharedUrl  string
	nestedUrl  string
	badUrl     string
)

func init() {
	server = httptest.NewServer(router.Routes())

	libraryUrl = fmt.Sprintf("%s/library", server.URL)
	sharedUrl = fmt.Sprintf("%s/shared/blank.mp3", server.URL)
	nestedUrl = fmt.Sprintf("%s/shared/folder/nested.mp3", server.URL)
	badUrl = fmt.Sprintf("%s/shared/notype", server.URL)
}

func TestLibrary(t *testing.T) {
	request, err := http.NewRequest("GET", libraryUrl, nil)

	res, err := http.DefaultClient.Do(request)

	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	if string(contents) != `{"blank.mp3":true,"folder/nested.mp3":true}` {
		t.Errorf("Expected: " + string(contents) + " to be " + `{"blank.mp3":true}`)
	}
}

func TestShared(t *testing.T) {
	request, err := http.NewRequest("GET", sharedUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected: ", res.StatusCode, " to be ", 200)
	}
}

func TestNested(t *testing.T) {
	request, err := http.NewRequest("GET", nestedUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected: ", res.StatusCode, " to be ", 200)
	}
}

func TestBad(t *testing.T) {
	request, err := http.NewRequest("GET", badUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 404 {
		t.Errorf("Expected: ", res.StatusCode, " to be ", 404)
	}
}
