package handlers

import (
	"testing"

	. "github.com/cs-lexliu/mocks-example/services/mocks"
)

func TestDoHandler(t *testing.T) {
	mockDoService := &DoService{}
	mockDoService.On("Do").Return(true, nil)
	doHandler := &doHandler{
		doService: mockDoService,
	}
	doHandler.Do()
}
