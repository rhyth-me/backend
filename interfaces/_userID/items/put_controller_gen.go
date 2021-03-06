// Package items ...
// generated version: 1.8.0
package items

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rhyth-me/backend/domain/model"
	"github.com/rhyth-me/backend/interfaces/props"
	"github.com/rhyth-me/backend/interfaces/wrapper"
	"github.com/rhyth-me/backend/pkg/authority"
	"github.com/rhyth-me/backend/pkg/random"
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
// @Param userID path string true ""
// @Param item body string false ""
// @Success 200 {object} PutResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /_userID/items [PUT]
func (p *PutController) Put(
	c echo.Context, req *PutRequest,
) (res *PutResponse, err error) {

	author := authority.GetIdentifier(c)
	if author.UID == "" || author.ScreenName != req.UserID {
		body := map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "You need to log in.",
		}
		return nil, wrapper.NewAPIError(http.StatusUnauthorized, body)
	}

	ctx := context.Background()

	// Generate itemID
	id := random.String(8)
	recode := model.Item{
		ID:      id,
		Snippet: req.Details,
		Author:  author.Google.ID,
	}

	// Add recode
	_, _, err = p.ControllerProps.Firestore.Collection(os.Getenv("ITEMS_COLLECTION")).Add(ctx, recode)
	if err != nil {
		return nil, wrapper.NewAPIError(http.StatusInternalServerError)
	}

	res = &PutResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Result:  recode,
	}

	return res, nil
}
