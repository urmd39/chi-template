package controller

import (
	"nutrition/service"
)

type exampleController interface {
}

type exampleControllerIml struct {
	exampleService service.ExampleService
}

func NewExampleController() exampleController {
	return &exampleControllerIml{
		exampleService: service.NewExampleService(),
	}
}

// @tags
// @Summary
// @Description
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string false "Authorization"
// @Success 200 {object}  response.Response
// @Router  / [get]
