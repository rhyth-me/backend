package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/domain/model"
)

// CustomContext - echoContext expansion
type CustomContext struct {
	echo.Context
	User LoginUser
}

// LoginUser - The identifier of the authenticated user.
type LoginUser struct {
	UID        string
	ScreenName string
	Google     model.Google
}
