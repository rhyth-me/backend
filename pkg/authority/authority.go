package authority

import (
	"context"
	r "reflect"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/pkg/firebase"
)

func MapToStruct(mapVal map[string]interface{}, val interface{}) (ok bool) {
	structVal := r.Indirect(r.ValueOf(val))
	for name, elem := range mapVal {
		structVal.FieldByName(name).Set(r.ValueOf(elem))
	}

	return
}

// Identify - Identify the user from JWT.
func Identify(c echo.Context, jwt string) *CustomContext {
	var auth = firebase.InitAuth()
	token, err := auth.VerifyIDToken(context.Background(), jwt)
	if err != nil {
		return &CustomContext{Context: c}
	}

	claims := token.Claims
	identities := claims["firebase"].(map[string]interface{})["identities"].(map[string]interface{})

	var screenName string = ""
	if val, ok := claims["screen_name"].(string); ok {
		screenName = val
	}

	cc := &CustomContext{Context: c, User: User{
		UID:        token.UID,
		ScreenName: screenName,
		Google: Google{
			ID:    identities["google.com"].([]interface{})[0].(string),
			Email: identities["email"].([]interface{})[0].(string),
		},
	}}

	return cc
}

// GetIdentifier - Get the identity of auth user.
func GetIdentifier(c echo.Context) User {
	return c.(*CustomContext).User
}
