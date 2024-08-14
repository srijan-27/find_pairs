package handlers

import (
	"gofr.dev/pkg/gofr"
	gofrHTTP "gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"

	"find_pairs/models"
	"find_pairs/services"
)

// FindPairs handles the request to take input and respond with the pairs of indices whose sum is equal to provided target
func FindPairs(ctx *gofr.Context) (interface{}, error) {
	var i models.Input

	// Bind the input body to the expected struct
	err := ctx.Bind(&i)
	if err != nil {
		return nil, gofrHTTP.ErrorInvalidParam{Params: []string{"body"}}
	}

	resp, err := services.FindPairs(i)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: models.Response{Solutions: resp}}, nil
}
