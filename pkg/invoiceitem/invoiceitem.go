package invoiceitem

import "time"

//Model of Invoiceitem
type Model struct {
	ID uint
	InvoiceHeaderID uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}