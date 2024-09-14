package algorithms

type Page struct {
	pageNumber string
	referenced bool
}

// SecondChance implementa o algoritmo de substituição de página com segunda chance.
func SecondChance(acessos []string, numFrames int) int {
	// Lista de páginas na memória
	memoria := make([]Page, 0, numFrames)

	faltas := 0
	pos := 0 // Posição do ponteiro (circular) para a página a ser substituída

	for _, pagina := range acessos {
		// Verifica se a página já está na memória
		encontrada := false
		for i := range memoria {
			if memoria[i].pageNumber == pagina {
				// Página já está na memória, damos uma segunda chance
				memoria[i].referenced = true
				encontrada = true
				break
			}
		}

		// Se a página não está na memória, ocorre uma falta de página
		if !encontrada {
			faltas++

			// Se a memória está cheia, precisamos substituir uma página
			if len(memoria) >= numFrames {
				for {
					// Se a página não tem segunda chance, substituímos
					if !memoria[pos].referenced {
						// Substitui a página e ajusta o bit de referência
						memoria[pos] = Page{pageNumber: pagina, referenced: true}
						pos = (pos + 1) % numFrames // Avança o ponteiro circular
						break
					} else {
						// Dá uma segunda chance, reseta o bit e avança o ponteiro
						memoria[pos].referenced = false
						pos = (pos + 1) % numFrames
					}
				}
			} else {
				// Se há espaço, apenas adiciona a página na memória
				memoria = append(memoria, Page{pageNumber: pagina, referenced: true})
			}
		}
	}

	return faltas
}
