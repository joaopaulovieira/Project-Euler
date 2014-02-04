package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSumNumbersThreeAndFive(c *C) {
	result := sum_numbers(1000)
	c.Assert(233168, Equals, result)
}

func (s *MySuite) TestVerifyMultiple(c *C) {
	negative_test := is_multiple(6, 3)
	c.Assert(negative_test, Equals, true)
	positive_test := bool(is_multiple(8, 7))
	c.Assert(positive_test, Equals, false)
}
