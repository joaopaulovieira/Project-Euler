package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestFibonacciNumber(c *C) {
	fibonacci := fibonacci_sequence()
	c.Assert(fibonacci(), Equals, 0)
}

func (s *MySuite) TestSumFibonacciSequence(c *C) {
	sum_test := sum_sequence(3)
	c.Assert(sum_test, Equals, 2)
}

func (s *MySuite) TestBelowFourMillion(c *C) {
	test_positive := below_four_million(5000000)
	c.Assert(test_positive, Equals, true)
	test_negative := below_four_million(10)
	c.Assert(test_negative, Equals, false)
}

func (s *MySuite) TestEvenNumber(c *C) {
	test_positive := return_even_number(2)
	c.Assert(test_positive, Equals, true)
	test_negative := return_even_number(3)
	c.Assert(test_negative, Equals, false)
}
