package puck_test

import (
	"io/ioutil"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func LoadJSON(fixture string) string {
	bytes, _ := ioutil.ReadFile(fixture)
	return string(bytes)
}
