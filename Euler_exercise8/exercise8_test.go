package main

import(
	."launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestProductFiveNumbers(c *C) {
	result := fiveDigitProduct(0)
	c.Assert(result, Equals, 882)
}

func (s *MySuite) TestLargestProductInSeries(c *C) {
	result := largestProductInSeries()
	c.Assert(result, Equals, 40824)
}