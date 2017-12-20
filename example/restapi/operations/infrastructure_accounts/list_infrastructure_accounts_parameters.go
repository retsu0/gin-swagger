package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/retsu0/gin-swagger/api"
)

// ListInfrastructureAccountsEndpoint executes the core logic of the related
// route endpoint.
func ListInfrastructureAccountsEndpoint(handler func(ctx *gin.Context) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		resp := handler(ctx)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// ListInfrastructureAccountsParams contains all the bound params for the list infrastructure accounts operation
// typically these are obtained from a http.Request
//
// swagger:parameters listInfrastructureAccounts
type ListInfrastructureAccountsParams struct {
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ListInfrastructureAccountsParams) readRequest(ctx *gin.Context) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// vim: ft=go
