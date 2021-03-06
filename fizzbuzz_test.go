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
func (s *FizzBuzzSuite) TestThreeChangedToBuzz() {
	r := fizzbuzz.Run([]int{5})
	s.Equal([]string{"Buzz"}, r)
}

func (s *FizzBuzzSuite) TestThreeAndFiveChangedToBuzz() {
	r := fizzbuzz.Run([]int{15})
	s.Equal([]string{"FizzBuzz"}, r)
}

func (s *FizzBuzzSuite) TestManyNumbers() {
	r := fizzbuzz.Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 15})
	s.Equal([]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "FizzBuzz"}, r)
}
