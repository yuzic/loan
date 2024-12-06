package public

import (
	"net/http"

	"loan/internal/api"
	"loan/internal/entity"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type LoanHandler struct {
	api.ServerInterface
	loanService LoanServiceInterface
	logger      *zap.Logger
}

type LoanServiceInterface interface {
	CreateLoan(borrower string, amount, collateral float64) (*entity.Loan, error)
	RepayLoan(loanID int64, amount float64) error
}

func NewLoanHandler(
	loanService LoanServiceInterface,
	logger *zap.Logger,
) *LoanHandler {
	return &LoanHandler{
		logger:      logger,
		loanService: loanService,
	}
}

func (h *LoanHandler) CreateLoan(ctx echo.Context) error {
	var req api.LoanRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loan, err := h.loanService.CreateLoan(*req.Borrower, *req.Amount, *req.Collateral)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, toResponse(loan))
}
