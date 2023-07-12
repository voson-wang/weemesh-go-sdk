package weemesh

import (
	"flag"
	"fmt"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func TestAppManagement(t *testing.T) {
	TestingT(t)
}

type AppManagementSuite struct{}

var _ = Suite(&AppManagementSuite{})

var appName = flag.String("appName", "", "应用名称")

func (s *AppManagementSuite) TestGetDeviceState(c *C) {
	if *addr == "" || *accessKey == "" || *accessSecret == "" || *appName == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetAppProductKeys(*appName)
	if err != nil {
		c.Fatal(err)
	}

	fmt.Println(result)

	c.Assert(result.Code, Equals, http.StatusOK)
}
