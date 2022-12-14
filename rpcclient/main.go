package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	req, err := http.NewRequest(`POST`, `http://localhost:8081/delivery`, nil)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	body := []byte(`{
		"jsonprc": 2.0,
		"method": "sms.SendSMS",
		"params": [{}]
	})`)
	req.Header.Add(`Content-Type`, `application/json`)
	neededBody := io.NopCloser(bytes.NewReader(body))
	req.Body = neededBody
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error: %s\\n", err)
		return
	}
	defer resp.Body.Close()
}
