package q2

import (
	"errors"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Title      string
	BaseSalary float64
	Bonuses    []float64
}

func CalculateTotalSalary(emp *Employee) (float64, error) {
	if emp == nil {
		return 0, errors.New("ponteiro para empregado é nulo")
	}

	total := emp.BaseSalary
	for _, bonus := range emp.Bonuses {
		total += bonus
	}

	if total > 1500.0 {
		if !strings.HasPrefix(emp.Title, "Senior") {
			emp.Title = "Senior " + emp.Title
		}
	}

	return total, nil
}
