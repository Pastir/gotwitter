package gotwitter

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type (
	ClientReq struct {
		Method string
		Url    string
		Body   io.Reader
	}
)

func Client(url string, body io.Reader) ClientReq {
	c := ClientReq{
		Url:  url,
		Body: body,
	}
	return c
}

func (c *ClientReq) Request(event *Events) error {
	// create new request
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, (*event).EventMethod(), (*event).EventEndpoint(), c.Body)
	if err != nil {
		return err
	}

	// set headers
	for k, v := range (*event).EventHeader() {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response:", string(respBody))
	return nil
}
