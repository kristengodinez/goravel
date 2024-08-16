package controllers

import (
	"testing"

	"github.com/goravel/framework/contracts/http"
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
		"creditCardNumber": "required",
	}).Return(mockValidator, nil).Once()
	mockValidator.EXPECT().Fails().Return(false).Once()

	var creditCard CreditCard
	mockValidator.EXPECT().Bind(&creditCard).Run(func(creditCard any) {
		creditCard.(*CreditCard).CreditCardNumber = "123"
	}).Return(nil).Once()

	mockContext.EXPECT().Response().Return(mockResponse).Once()
	mockResponseStatus := mockFactory.ResponseStatus()
	mockResponse.EXPECT().Success().Return(mockResponseStatus).Once()

	resp := &gin.JsonResponse{}
	mockResponseStatus.EXPECT().Json(http.Json{
		"creditCardNumber": "123",
	}).Return(resp).Once()

	s.Equal(resp, NewLuhnController().Json(mockContext))

	mockContext.AssertExpectations(s.T())
	mockRequest.AssertExpectations(s.T())
	mockResponse.AssertExpectations(s.T())
	mockValidator.AssertExpectations(s.T())
	mockResponseStatus.AssertExpectations(s.T())
}
