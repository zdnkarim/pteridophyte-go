package server

import (
	"pteridophyte-app/pkg/response"

	cRoute "pteridophyte-app/internal/class/route"
	fRoute "pteridophyte-app/internal/family/route"
	gRoute "pteridophyte-app/internal/genus/route"
	kRoute "pteridophyte-app/internal/kingdom/route"
	oRoute "pteridophyte-app/internal/order/route"
	phRoute "pteridophyte-app/internal/phylum/route"
	plRoute "pteridophyte-app/internal/plant/route"
	sRoute "pteridophyte-app/internal/species/route"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	// Group the API
	api := e.Group("/api")

	// Route for not found
	e.RouteNotFound("/*", func(c echo.Context) error {
		return response.GenerateResponse(c, 404, response.WithMessage("path not found"))
	})

	// Route for ping
	api.GET("", func(c echo.Context) error {
		return response.GenerateResponse(c, 200, response.WithMessage("Welcome to the API Backend"))
	})

	initializeHandler(api, db)
}

func initializeHandler(e *echo.Group, db *gorm.DB) {

	masterGroup := e.Group("/master")

	// Initialize Class Route
	cRoute.NewClassRoute(masterGroup, db)

	// Initialize Route Route
	fRoute.NewFamilyRoute(masterGroup, db)

	// Initialize Genus Route
	gRoute.NewGenusRoute(masterGroup, db)

	// Initialize Kingdom Route
	kRoute.NewKingdomRoute(masterGroup, db)

	// Initialize Order Route
	oRoute.NewOrderRoute(masterGroup, db)

	// Initialize Phylum Route
	phRoute.NewPhylumRoute(masterGroup, db)

	// Initialize Plant Route
	plRoute.NewPlantRoute(masterGroup, db)

	// Initialize Species Route
	sRoute.NewSpeciesRoute(masterGroup, db)

	// Initialize Plant Image Route
	// pIRoute.NewPlantImageRoute(masterGroup, db)
}
