package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	HOST = "http://localhost:8080"
)

func buildHeader(request *http.Request, header map[string]string) {
	h := make(map[string]string)
	h["Content-Type"] = "application/json"
	if header != nil {
		for k, v := range header {
			h[k] = v
		}
	}
	for k, v := range h {
		request.Header.Add(k, v)
	}
}

func request[TBody any, TRes any](method string, url string, body *TBody, header map[string]string) (TRes, error) {
	var t TRes
	var bodyReader io.Reader

	if body != nil {
		bs, err := json.Marshal(body)
		if err != nil {
			return t, err
		}
		bodyReader = bytes.NewReader(bs)
	}

	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return t, err
	}

	buildHeader(request, header)

	for k, v := range header {
		request.Header.Add(k, v)
	}

	res, err := new(http.Client).Do(request)
	if err != nil {
		return t, err
	}

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		if err != io.EOF {
			return t, err
		}
	}
	defer res.Body.Close()

	if err := json.Unmarshal(bs, &t); err != nil {
		return t, err
	}
	return t, err
}

//GET
func GET[T any](url string, header map[string]string) (T, error) {
	return request[any, T]("GET", url, nil, header)
}

//POST
func Post[TBody any, TRes any](url string, param *TBody, header map[string]string) (TRes, error) {
	return request[TBody, TRes]("POST", url, nil, header)
}
