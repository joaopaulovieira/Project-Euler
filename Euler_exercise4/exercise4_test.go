package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPalindromicNumber(c *C) {
	palindromic := palindromic_number(11)
	not_palindromic := palindromic_number(1000)
	c.Assert(palindromic, Equals, true)
	c.Assert(not_palindromic, Equals, false)
}

func (s *MySuite) TestDigitsNumber(c *C) {
	largest_palindromic := largest_palindromic_number()
	c.Assert(largest_palindromic, Equals, 906609)
}