package q5

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Sale struct {
	Code     string
	Date     string
	Products []*Product
}

func CalculateTotalSales(sales map[string]*Sale) float64 {
	total := 0.0

	for _, sale := range sales {
		for _, product := range sale.Products {
			total += product.Price
		}
	}

	return total
}
