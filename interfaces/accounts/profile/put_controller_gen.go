// Package profile ...
// generated version: 1.8.0
package profile

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/interfaces/wrapper"
	"github.com/rhyth-me/backend/pkg/authority"
)

// PutController ...
type PutController struct {
	*props.ControllerProps
}

// NewPutController ...
func NewPutController(cp *props.ControllerProps) *PutController {
	p := &PutController{
		ControllerProps: cp,
	}
	return p
}

// Put ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param profile body model.SocialProfile true ""
// @Success 200 {object} PutResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /accounts/profile [PUT]
func (p *PutController) Put(
	c echo.Context, req *PutRequest,
) (res *PutResponse, err error) {

	user := authority.GetIdentifier(c)
	if user.UID == "" {
		body := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "You need to log in.",
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	// Fetch user by uid.
	ctx := context.Background()
	_, err = p.ControllerProps.Firestore.Collection("users").Doc(user.Google.ID).Update(ctx, []firestore.Update{
		{
			Path:  "profile",
			Value: req.Profile,
		},
	})
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  req.Profile,
	}

	return res, nil
}
