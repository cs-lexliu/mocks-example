package handlers

import (
	"fmt"

	"github.com/cs-lexliu/mocks-example/services"
)

type doHandler struct {
	doService services.DoService
}

func (h *doHandler) Do() {
	res, err := h.doService.Do()
	if err != nil {
		err := fmt.Errorf("doService.Do: %w", err)
		fmt.Println(err)
		return
	}
	fmt.Println("success:", res)
}
