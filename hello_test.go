package helloworld

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloGet_サーバーとしてHelloWorldを返す(t *testing.T) {
	// Arrange
	reqBody := bytes.NewBufferString("request body")
	req := httptest.NewRequest(http.MethodGet, "/", reqBody)

	// レスポンスを受け止める*httptest.ResponseRecorder
	got := httptest.NewRecorder()

	// Act
	helloGet(got, req)

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if got.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, got.Code)
	}
	// Bodyは*bytes.Buffer型なので文字列の比較は少しラク
	if got := got.Body.String(); got != "Hello, World!" {
		t.Errorf("want %s, but %s", "Hello, World!", got)
	}
}

func Test_ClientHello_クライアントとして指定したURLからレスポンスを受け取る(t *testing.T) {
	// Arrange
	// ServeMuxオブジェクトなどを用意してルーティングしてもよい
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{ \"hoge\": \"fuga\" }")
	})
	// 別goroutine上でリッスンが開始される
	ts := httptest.NewServer(h)
	defer ts.Close()

	// Act
	resp := ClientHello(ts.URL)

	// Assertion
	if resp["hoge"] != "fuga" {
		t.Fatal("error")
	}
}

func Test_ServerClientHello_指定したURLのレスポンスからhogeの値を返す(t *testing.T) {
	// Arrange
	got := httptest.NewRecorder()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{ \"hoge\": \"fuga\" }")
	})
	// 別goroutine上でリッスンが開始される
	ts := httptest.NewServer(h)
	defer ts.Close()

	// Act
	ServerClientHello(got, nil, ts.URL)
	// Assertion
	if got := got.Body.String(); got != "fuga" {
		t.Errorf("want %s, but %s", "fuga", got)
	}
}
