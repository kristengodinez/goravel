package controllers

import (
	"testing"

	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/goravel/gin"
	"github.com/stretchr/testify/suite"
)

type LuhnControllerTestSuite struct {
	suite.Suite
}

func TestLuhnControllerTestSuite(t *testing.T) {
	suite.Run(t, &LuhnControllerTestSuite{})

}

func (s *LuhnControllerTestSuite) SetupTest() {
}

func (s *LuhnControllerTestSuite) TearDownTest() {
}

func (s *LuhnControllerTestSuite) TestJson() {
	mockFactory := testingmock.Factory()
	mockContext := mockFactory.Context()
	mockRequest := mockFactory.ContextRequest()
	mockResponse := mockFactory.ContextResponse()
	mockValidator := mockFactory.ValidationValidator()
	mockContext.EXPECT().Request().Return(mockRequest).Once()
	mockRequest.EXPECT().Validate(map[string]string{
		"numbers": "required",
	}).Return(mockValidator, nil).Once()
	mockValidator.EXPECT().Fails().Return(false).Once()

	var creditCard CreditCard
	mockValidator.EXPECT().Bind(&creditCard).Run(func(creditCard any) {
		var numbers []CreditCardNumber
		numbers = []CreditCardNumber{
			{
				Number: "123",
			},
			{
				Number: "3379 5135 6110 8795",
			},
		}
		creditCard.(*CreditCard).Numbers = numbers
	}).Return(nil).Once()

	mockContext.EXPECT().Response().Return(mockResponse).Once()
	mockResponseStatus := mockFactory.ResponseStatus()
	mockResponse.EXPECT().Success().Return(mockResponseStatus).Once()

	resp := &gin.JsonResponse{}
	expected_results := CreditCardResponse{
		[]*CreditCardResult{
			{
				CreditCardNumber: "123",
				IsValid:          false,
			},
			{
				CreditCardNumber: "3379 5135 6110 8795",
				IsValid:          true,
			},
		},
	}

	mockResponseStatus.EXPECT().Json(expected_results).Return(resp).Once()

	s.Equal(resp, NewLuhnController().Json(mockContext))

	mockContext.AssertExpectations(s.T())
	mockRequest.AssertExpectations(s.T())
	mockResponse.AssertExpectations(s.T())
	mockValidator.AssertExpectations(s.T())
	mockResponseStatus.AssertExpectations(s.T())
}
