package main

import(
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestSumSquare(c *C) {
	sum_square := sum_squares(10)
	c.Assert(sum_square, Equals, float64(385))
}

func (s *MySuite) TestSquareSum(c *C) {
	square_sum := square_sum(10)
	c.Assert(square_sum, Equals, float64(3025))
}

func (s *MySuite) TestDiferenceBetweenSquareSumAndSumSquare(c *C) {
	test_diference := diference_between_square_sum_and_sum_squares(10)
	c.Assert(test_diference, Equals, float64(2640))
}