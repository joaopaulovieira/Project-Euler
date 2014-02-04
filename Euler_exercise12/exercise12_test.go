package main

import(
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrimeFactor(c *C) {
	factor := factorPrime(6)
	c.Assert(factor, Equals, 4)
}

func (s *MySuite) TestMMC(c *C) {
	test_true := MMC(10, 2)
	c.Assert(test_true, Equals, true)
	test_false := MMC(3,2)
	c.Assert(test_false, Equals, false)
}

func (s *MySuite) TestMakeTriangularNumbers(c *C) {
	result := makeTriangularNumberReturnDivisors(400)
	c.Assert(result, Equals, result)
}