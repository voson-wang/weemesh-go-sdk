package weemesh

import (
	"flag"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

var addr = flag.String("address", "", "api服务地址，例如：http://localhost:2245")
var accessKey = flag.String("accessKey", "", "accessKey")
var accessSecret = flag.String("accessSecret", "", "accessSecret")
var productKey = flag.String("productKey", "", "productKey")
var orgID = flag.String("orgID", "", "orgID")

func TestThingModel(t *testing.T) {
	TestingT(t)
}

type ThingModelSuite struct{}

var _ = Suite(&ThingModelSuite{})

func (s *ThingModelSuite) TestGetThingModel(c *C) {

	if *addr == "" || *accessKey == "" || *accessSecret == "" || *productKey == "" || *orgID == "" {
		c.Skip("参数不完整")
	}

	client := NewClient(*accessKey, *accessSecret, *addr)

	result, err := client.GetThingModel(*productKey, *orgID)
	if err != nil {
		c.Fatal(err)
	}

	c.Assert(result.Code, Equals, http.StatusOK)
}
