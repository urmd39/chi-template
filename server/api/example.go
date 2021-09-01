package api

import (
	"log"
	"nutrition/controller"
	"nutrition/middleware"

	"github.com/go-chi/chi"
)

func NewExampleRouter(subR chi.Router) {
	exampleController := controller.NewExampleController()
	log.Println(exampleController) // Log Test

	// Private routes
	subR.Group(func(r chi.Router) {
		r.Use(middleware.Auth())

	})

	// Public routes
	subR.Group(func(r chi.Router) {

	})
}
