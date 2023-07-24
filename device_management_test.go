package weemesh

import (
	"flag"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func TestDeviceManagement(t *testing.T) {
	TestingT(t)
}

type DeviceManagementSuite struct{}

var _ = Suite(&DeviceManagementSuite{})

var sn = flag.String("sn", "", "设备sn")

func (s *DeviceManagementSuite) TestGetDeviceState(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" || *sn == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetDeviceState(*sn)
	if err != nil {
		c.Fatal(err)
	}

	c.Assert(result.Code, Equals, http.StatusOK)
}

func (s *DeviceManagementSuite) TestGetDeviceLatestProperty(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" || *sn == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetDeviceLatestProperty(*sn)
	if err != nil {
		c.Fatal(err)
	}

	c.Assert(result.Code, Equals, http.StatusOK)
}
