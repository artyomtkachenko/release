package webdav

import (
	"errors"
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

func Test_postFile_success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.
			fmt.Fprintln(w, "Hello Obama")
	}))
	defer ts.Close()

	body, _ := sendRequest(ts.URL)
	if !strings.Contains(string(body), "Hello Obama") {
		t.Errorf("Expected: body to be Hello, got [%s]\n", body)
	}
}
