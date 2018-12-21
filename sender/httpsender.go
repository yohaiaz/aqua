package sender

import (
	"bytes"
	"fmt"
	"net/http"
)

type Http struct {
	endpoint string
}

func NewHttp(url string) *Http {
	return &Http{
		endpoint: url,
	}
}

func (h *Http) Send(body string) {
	fmt.Printf("sending %s to %s \n", body, h.endpoint)

	resp, err := http.Post(h.endpoint, "application/json", bytes.NewBuffer([]byte(body)))

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp.Status)
	}
}
