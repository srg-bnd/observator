package client

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendMetrics() error {
	return nil
}
