package weemesh

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func newRequest[T any](request *http.Request, c *Client) (*CommonResult[T], error) {
	v := url.Values{}

	signature, err := c.signature()
	if err != nil {
		return nil, err
	}

	et := strconv.Itoa(int(c.timestamp.Unix()))

	v.Set("version", Version230503)
	v.Set("method", MethodSha1)
	v.Set("signature", signature)
	v.Set("access_key", c.accessKey)
	v.Set("timestamp", et)

	request.Header.Set("authentication", v.Encode())

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CommonResult[T]

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
