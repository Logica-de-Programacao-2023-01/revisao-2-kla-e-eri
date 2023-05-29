package bonus

import "fmt"

func Destino(caminhos [][2]string) (string, error) {
	cidades := make(map[string]string)

	for _, caminho := range caminhos {
		origem := caminho[0]
		destino := caminho[1]

		cidades[origem] = destino
	}

	for _, destino := range cidades {
		if _, ok := cidades[destino]; !ok {
			return destino, nil
		}
	}

	return "", fmt.Errorf("erro: nenhum destino encontrado")
}
