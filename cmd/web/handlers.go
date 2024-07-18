package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

var ProxyCounter = 0
var ProxyData = &sync.Map{}

type ProxyRequest struct {
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
}
type ProxyResponse struct {
	ID      int                 `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server for proxy"))
}

// @Summary Sends request to outer service
// @Tags Proxy
// @Accept json
// @Param request body ProxyRequest true "Request to outer service"
// @Produce json
// @Success 200 {object} ProxyResponse
// @Router /proxy [post]
func (app *application) proxy(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &ProxyRequest{}

	bytes, err := io.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		app.errorLog.Fatal(err)
	}

	err = json.Unmarshal(bytes, &bodyRequest)
	if err != nil && err != io.EOF {
		app.errorLog.Fatal(err)
	}
	app.infoLog.Printf("Proxy call method:%s, url:%s\n", bodyRequest.Method, bodyRequest.URL)

	resToOuter, err := http.NewRequest(bodyRequest.Method, bodyRequest.URL, nil)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	for k, values := range bodyRequest.Headers {
		for _, v := range values {
			resToOuter.Header.Add(k, v)
		}
	}
	resFromOuter, err := http.DefaultClient.Do(resToOuter)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	bodyFromOuter, err := io.ReadAll(resFromOuter.Body)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	ProxyCounter++
	proxyResponse := &ProxyResponse{
		ID:      ProxyCounter,
		Status:  resFromOuter.Status,
		Length:  len(bodyFromOuter),
		Headers: resFromOuter.Header,
	}

	// saving locally in sync.Map named ProxyData
	ProxyData.Store(proxyResponse.ID, proxyResponse)

	jsonResponse, err := json.MarshalIndent(proxyResponse, "", "\t")
	if err != nil {
		app.errorLog.Fatal(err)
	}
	w.Write(jsonResponse)
}
