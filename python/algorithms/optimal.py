import collections

def optimal(acessos, tamanho_memoria):
    memoria = []
    faltas = 0
    carregamentos = collections.defaultdict(int)

    for i, pagina in enumerate(acessos):
        if pagina not in memoria:
            faltas += 1
            if len(memoria) < tamanho_memoria:
                memoria.append(pagina)
            else:
                # Encontrar a página que será usada mais tarde ou nunca
                proximo_uso = {p: float('inf') for p in memoria}
                for j in range(i + 1, len(acessos)):
                    if acessos[j] in proximo_uso and proximo_uso[acessos[j]] == float('inf'):
                        proximo_uso[acessos[j]] = j
                # Escolher a página que será usada mais tarde ou nunca
                pagina_a_substituir = max(proximo_uso, key=proximo_uso.get)
                memoria[memoria.index(pagina_a_substituir)] = pagina
            carregamentos[pagina] += 1
    return faltas, carregamentos
