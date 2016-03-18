package webdav

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//publish to webdav endpoint

func Test_check(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	check(errors.New("foobar"))
}

func Test_putFile_success(t *testing.T) {
	var file string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output, _ := ioutil.ReadAll(r.Body)
		file = string(output)
	}))
	defer ts.Close()

	status, _ := putFile("testdata/test.txt", ts.URL, "foo", "bar")
	if !strings.Contains(status, "200") {
		t.Errorf("Expected: Status to be 200, got [%s]\n", status)
	}

	if !strings.Contains(file, "foo bar") {
		t.Errorf("Expected: Contents to contain Foo Bar, got:\n\n%s\n", file)
	}
}
