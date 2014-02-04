package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCollatzSequence(c *C) {
	result := makeCollatzSequence(13)
	c.Assert(result, Equals, 40)
}

func (s *MySuite) TestLargestCollatzSequence(c *C) {
	result := largestCollatzSequence(1000000)
	c.Assert(result, Equals, 837799)
}