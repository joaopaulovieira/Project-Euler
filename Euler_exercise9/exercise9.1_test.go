package main

import(
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSumPythagoreanTripletEquals1000(c *C) {
	test_false := sumPythagoreanTripletEquals1000(2, 3, 4)
	c.Assert(test_false, Equals, false)
}

func (s *MySuite) TestPythagoreanTriplet(c *C) {
	result := pythagoreanTriplet()
	c.Assert(result, Equals, 31875000)
}

func (s *MySuite) TestIsPythagoreanTripletNumbers(c *C) {
	test_false := isPythagoreanTriplet(3, 9, 2)
	c.Assert(test_false, Equals, false)
}