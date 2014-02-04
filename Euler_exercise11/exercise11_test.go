package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestMaximum(c *C) {
	result := Maximum(6,5)
	c.Assert(result, Equals, 6)
}

func (s *MySuite) TestMatrix20x20(c *C) {
	result := Matrix20x20()
	c.Assert(result, Equals, result)
}