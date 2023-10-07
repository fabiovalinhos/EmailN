package endpoints

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignResponse := contract.CampaignResponse{
		ID:      "4321",
		Name:    "teste",
		Content: "hi!",
		Status:  "Peding",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	// usando o type assertion do golang abaixo
	assert.Equal(campaignResponse.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignGetById(rr, req)

	// usando o type assertion do golang abaixo
	assert.Equal(errExpected.Error(), err.Error())
}
