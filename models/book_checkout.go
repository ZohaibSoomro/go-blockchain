package models

type BookCheckout struct {
	Book         *Book  `json:"book"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

func NewBookCheckout(b *Book, cd string) *BookCheckout {
	return &BookCheckout{
		Book:         b,
		CheckoutDate: cd,
	}
}
