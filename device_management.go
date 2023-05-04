package weemesh

import (
	"net/http"
	"net/url"
)

func (c *Client) GetDeviceState(sn string) (*CommonResult[bool], error) {

	params := url.Values{}
	params.Set("sn", sn)

	request, err := http.NewRequest(http.MethodGet, c.address+"/"+APIVersion+"/device/state?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	return newRequest[bool](request, c)
}
