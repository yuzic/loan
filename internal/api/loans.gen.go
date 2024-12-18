// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// LoanRequest defines model for LoanRequest.
type LoanRequest struct {
	Amount     *float64 `json:"amount,omitempty"`
	Borrower   *string  `json:"borrower,omitempty"`
	Collateral *float64 `json:"collateral,omitempty"`
}

// LoanResponse defines model for LoanResponse.
type LoanResponse struct {
	Amount     *float64 `json:"amount,omitempty"`
	Collateral *float64 `json:"collateral,omitempty"`
	DueDate    *string  `json:"dueDate,omitempty"`
	LoanId     *int64   `json:"loanId,omitempty"`
}

// RepayRequest defines model for RepayRequest.
type RepayRequest struct {
	Amount *float64 `json:"amount,omitempty"`
}

// CreateLoanJSONRequestBody defines body for CreateLoan for application/json ContentType.
type CreateLoanJSONRequestBody = LoanRequest

// RepayLoanJSONRequestBody defines body for RepayLoan for application/json ContentType.
type RepayLoanJSONRequestBody = RepayRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a loan
	// (POST /loans)
	CreateLoan(ctx echo.Context) error
	// Repay a loan
	// (POST /loans/{loanId}/repay)
	RepayLoan(ctx echo.Context, loanId uint64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateLoan converts echo context to params.
func (w *ServerInterfaceWrapper) CreateLoan(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateLoan(ctx)
	return err
}

// RepayLoan converts echo context to params.
func (w *ServerInterfaceWrapper) RepayLoan(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "loanId" -------------
	var loanId uint64

	err = runtime.BindStyledParameterWithOptions("simple", "loanId", ctx.Param("loanId"), &loanId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter loanId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RepayLoan(ctx, loanId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/loans", wrapper.CreateLoan)
	router.POST(baseURL+"/loans/:loanId/repay", wrapper.RepayLoan)

}
