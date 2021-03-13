package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhyth-me/backend/api/apigen/props"
	"github.com/rhyth-me/backend/pkg/firebase/auth"
	"github.com/rhyth-me/backend/pkg/validator"
)

// initControllerProps - setup controller props
func initControllerProps() *props.ControllerProps {
	p := new(props.ControllerProps)

	return p
}

// initEchoSetting - setup echo
func initEchoSetting(e *echo.Echo) *echo.Echo {
	e.Debug = true
	e.HideBanner = true
	e.Validator = validator.NewValidator()
	e.Use(middleware.Recover())

	// Custom context for identifying users
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := auth.Identify(c)
			return next(cc)
		}
	})

	// CORS config
	origins := []string{"https://rhyth.me"}
	if os.Getenv("STAGING") == "true" {
		origins = []string{"*"}
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: origins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		MaxAge:       3600,
	}))

	return e
}
