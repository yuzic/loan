package service

import (
	"time"

	"loan/internal/entity"
)

type LoanService struct{}

func NewLoanService() *LoanService {
	return &LoanService{}
}

func (*LoanService) CreateLoan(_ string, _, _ float64) (*entity.Loan, error) {
	return &entity.Loan{
		ID:         8888,
		Amount:     100000,
		Collateral: 67777,
		DueDate:    time.Now().Unix(),
	}, nil
}

func (*LoanService) RepayLoan(_ int64, _ float64) error {
	return nil
}
