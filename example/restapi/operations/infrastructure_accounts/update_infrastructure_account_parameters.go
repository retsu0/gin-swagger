package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
	"github.com/retsu0/gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/retsu0/gin-swagger/example/models"
)

// EndpointUpdateInfrastructureAccount executes the core logic of the related
// route endpoint.
func EndpointUpdateInfrastructureAccount(handler func(ctx *gin.Context, params *UpdateInfrastructureAccountParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &UpdateInfrastructureAccountParams{}
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// UpdateInfrastructureAccountParams contains all the bound params for the update infrastructure account operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateInfrastructureAccount
type UpdateInfrastructureAccountParams struct {

	/*ID of the infrastructure account.
	  Required: true
	  Pattern: ^[a-z][a-z0-9-:]*[a-z0-9]$
	  In: path
	*/
	AccountID string
	/*Infrastructure Account that will be updated.
	  Required: true
	  In: body
	*/
	InfrastructureAccount *models.InfrastructureAccountUpdate
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateInfrastructureAccountParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rAccountID := []string{ctx.Param("account_id")}
	if err := o.bindAccountID(rAccountID, true, formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(ctx.Request) {
		var body models.InfrastructureAccountUpdate
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("infrastructureAccount", "body"))
			} else {
				res = append(res, errors.NewParseError("infrastructureAccount", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.InfrastructureAccount = &body
			}
		}

	} else {
		res = append(res, errors.Required("infrastructureAccount", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateInfrastructureAccountParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.AccountID = raw

	if err := o.validateAccountID(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateInfrastructureAccountParams) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Pattern("account_id", "path", o.AccountID, `^[a-z][a-z0-9-:]*[a-z0-9]$`); err != nil {
		return err
	}

	return nil
}

// vim: ft=go
