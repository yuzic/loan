package public

import (
	"strconv"

	"loan/internal/api"
	"loan/internal/entity"
)

func toResponse(e *entity.Loan) *api.LoanResponse {
	return &api.LoanResponse{
		LoanId:     intPtr(e.ID),
		Amount:     float64Ptr(e.Amount),
		Collateral: float64Ptr(e.Collateral),
		DueDate:    stringPtr(strconv.FormatInt(e.DueDate, 10)),
	}
}

func intPtr(i int64) *int64 {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func stringPtr(s string) *string {
	return &s
}
