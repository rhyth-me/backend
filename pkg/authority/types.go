package authority

import "github.com/labstack/echo/v4"

// CustomContext - echoContext expansion
type CustomContext struct {
	echo.Context
	User User
}

// User - The identifier of the authenticated user.
type User struct {
	UID        string
	ScreenName string
	Google     Google
}

// Google - Authenticated user's account.
type Google struct {
	ID    string `firestore:"id" json:"-"`
	Email string `firestore:"email"  json:"-"`
}
