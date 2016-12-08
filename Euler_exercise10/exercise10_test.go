package main

import(
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrimeNumber(c *C) {
	prime := is_prime(2)
	c.Assert(prime, Equals, true)
	not_prime := is_prime(6)
	c.Assert(not_prime, Equals, false)
}

func (s *MySuite) TestSumPrimeNumbersBelowParamater(c *C) {
	result := sumPrimeNumbersBelowsParamater(10)
	c.Assert(result, Equals, float64(17))
}