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

func (s *MySuite) TestsSquencePrimeNumber(c *C) {
	test_prime := find_prime_number(6)
	c.Assert(test_prime, Equals, 13)
}