package helloworld

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleChatWorkAPIResponse(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	cli := &http.Client{}
	req, err := http.NewRequestWithContext(context.TODO(), "GET", ts.URL, strings.NewReader(""))
	if err != nil {
		t.Errorf("NewRequest failed: %v", err)
	}

	// Act
	resp, err := cli.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	want := "Hello, World!"
	if string(got) != want {
		t.Errorf("want %q, but %q", want, got)
	}
}
