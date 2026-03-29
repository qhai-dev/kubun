package feishu

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

type Client struct {
	sdk *lark.Client
}

func NewClient() *Client {
	client := lark.NewClient("cli_a8fe2ed0fe71100b", "GUNBl45fn2u1LdyAlxFJWbhzyaBNIFYS")

	return &Client{
		sdk: client,
	}
}
