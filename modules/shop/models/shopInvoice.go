package models

type ShopInvoice struct {
	ID                 uint
	TitleType          int
	Type               int
	Title              string
	TaxNo              string
	Tel                string
	Bank               string
	Account            string
	Address            string
	UserID             int
	Money              float32
	Status             int
	InvoiceType        int
	ReceiverEmail      string
	ReceiverName       string
	ReceiverTel        string
	ReceiverRegionID   int
	ReceiverRegionName string
	ReceiverAddress    string
	CreatedAt          int
	UpdatedAt          int
}
