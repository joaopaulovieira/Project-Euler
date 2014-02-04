package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrimeFactor(c *C) {
	factor := factor_prime(13195)
	c.Assert(factor, Equals, 29)
}

func (s *MySuite) TestMMC(c *C) {
	test_number, test_minimum := MMC(10, 2)
	c.Assert(test_number, Equals, 5)
	c.Assert(test_minimum, Equals, 2)
}
