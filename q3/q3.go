package q3

import "errors"

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	if product == nil {
		return errors.New("ponteiro para produto é nulo")
	}

	for code, quantity := range sales {
		if code == product.Code {
			newQuantity := product.Quantity - quantity
			if newQuantity < 0 {
				return errors.New("quantidade em estoque insuficiente")
			}
			product.Quantity = newQuantity
			break
		}
	}

	return nil
}
