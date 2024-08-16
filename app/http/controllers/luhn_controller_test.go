package controllers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LuhnControllerTestSuite struct {
	suite.Suite
}

func TestLuhnControllerTestSuite(t *testing.T) {
	suite.Run(t, &LuhnControllerTestSuite{})

}
