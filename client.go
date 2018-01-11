package pubgopgg

import (
	"net/http"
)

const (
	API_ROOT = "https://pubg.op.gg/"
)

type Client struct {
	*http.Client
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}
