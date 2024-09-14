package algorithms

func Optimal(acessos []string, numFrames int, mapaFuturo map[string][]int) int {
	memoria := make([]string, 0, numFrames) // Armazena as páginas na memória
	faltas := 0

	for _, pagina := range acessos {
		// Atualiza o mapa de acessos futuros, removendo o acesso atual
		if len(mapaFuturo[pagina]) > 0 {
			mapaFuturo[pagina] = mapaFuturo[pagina][1:]
		}

		// Verifica se a página já está na memória
		encontrado := false
		for _, p := range memoria {
			if p == pagina {
				encontrado = true
				break
			}
		}

		// Se a página não está na memória, ocorre uma falta de página
		if !encontrado {
			faltas++

			// Se a memória está cheia, encontrar uma página para substituir
			if len(memoria) >= numFrames {
				indiceParaSubstituir := -1
				maisDistante := -1

				// Verificar quando cada página será usada no futuro
				for j, memPage := range memoria {
					if len(mapaFuturo[memPage]) == 0 {
						// A página memPage não será mais usada
						indiceParaSubstituir = j
						break
					} else if mapaFuturo[memPage][0] > maisDistante {
						// A página memPage será usada, mas mais tarde do que as outras
						maisDistante = mapaFuturo[memPage][0]
						indiceParaSubstituir = j
					}
				}

				// Substituir a página na posição encontrada
				memoria[indiceParaSubstituir] = pagina
			} else {
				// Se há espaço, simplesmente adicione a página
				memoria = append(memoria, pagina)
			}
		}
	}

	return faltas
}
