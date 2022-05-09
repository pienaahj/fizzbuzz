package fizzbuzz_test

import (
	"testing"

	"github.com/pienaahj/fizzbuzz"
	"github.com/stretchr/testify/suite"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (s *FizzBuzzSuite) TestOneUnchanged() {
	r := fizzbuzz.Run([]int{1})
	s.Equal([]string{"1"}, r)
}

func (s *FizzBuzzSuite) TestTwoUnchanged() {
	r := fizzbuzz.Run([]int{2})
	s.Equal([]string{"2"}, r)
}

func (s *FizzBuzzSuite) TestThreeChangedToFizz() {
	r := fizzbuzz.Run([]int{3})
	s.Equal([]string{"Fizz"}, r)
}
