package puck_test

import (
	"github.com/KAllan357/puck"
	. "launchpad.net/gocheck"
)

type ResourcesSuite struct {
	configJSON string
}

func (r *ResourcesSuite) SetUpSuite(c *C) {
	r.configJSON = LoadJSON("fixtures/sample-request.json")
}

var _ = Suite(&ResourcesSuite{})

func (r *ResourcesSuite) Test_ParseResourcesJSON(c *C) {
	resources := puck.ParseResourcesJSON([]byte(r.configJSON))
	resource := resources.Resources[0]
	c.Check(resource.Type, Equals, "directory")
	c.Check(resource.NameAttribute, Equals, "/foo/bar")
	c.Check(resource.Attributes["recursive"], Equals, true)
	c.Check(resource.Attributes["action"], Equals, "create")
}
