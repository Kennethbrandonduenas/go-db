package invoiceheader

import "time"

//Model of Invoiceheader
type Model struct {
	ID uint
	Client string
	CreatedAt time.Time
	UpdatedAt time.Time
}