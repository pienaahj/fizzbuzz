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
