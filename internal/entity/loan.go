package entity

type Loan struct {
	ID         int64
	Borrower   string
	Amount     float64
	Collateral float64
	Interest   float64
	DueDate    int64
	IsRepaid   bool
}
