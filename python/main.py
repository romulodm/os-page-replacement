import os
from algorithms.optimal import optimal
from algorithms.second_chance import second_chance

# Função para ler os acessos de um arquivo
def ler_acessos(arquivo):
    acessos = []
    with open(arquivo, 'r') as f:
        for linha in f:
            pagina = linha.strip()
            acessos.append(pagina)
    return acessos

# Função para converter o tamanho da memória física em número de páginas
def calcular_paginas(memoria_fisica_kb):
    return memoria_fisica_kb // 4  # Cada página tem 4 KB

# Função para converter o tamanho da memória física em KB
def converter_tamanho_em_kb(tamanho):
    unidades = {'GB': 1024**2, 'MB': 1024, 'KB': 1}
    return tamanho * unidades.get('KB', 1)

# Função principal para rodar os algoritmos e calcular faltas de página
def rodar_simulacao(arquivo_acessos):
    acessos = ler_acessos(arquivo_acessos)

    tamanhos_memoria = {
        '1 GB': converter_tamanho_em_kb(1024**2),
        '128 MB': converter_tamanho_em_kb(128),
        '16 MB': converter_tamanho_em_kb(16),
        '8 KB': 8
    }

    for descricao, tamanho_memoria_kb in tamanhos_memoria.items():
        tamanho_memoria = calcular_paginas(tamanho_memoria_kb)

        print(f"\nSimulação com memória física de {descricao} ({tamanho_memoria_kb} KB):")

        # Algoritmo Ótimo
        faltas_otimo, carregamentos_otimo = optimal(acessos, tamanho_memoria)

        # Algoritmo Second Chance
        faltas_second_chance, carregamentos_second_chance = second_chance(acessos, tamanho_memoria)

        # Calculando eficiência
        eficiencia = (faltas_otimo / faltas_second_chance) * 100

        # Exibindo os resultados
        print(f"Com o algoritmo Ótimo ocorrem {faltas_otimo} faltas de página.")
        print(f"Com o algoritmo Second Chance ocorrem {faltas_second_chance} faltas de página, atingindo {eficiencia:.2f}% do desempenho do Ótimo.")

        # Perguntar se deseja listar carregamentos
        listar = input("Deseja listar o número de carregamentos (s/n)? ").lower()
        if listar == 's':
            print(f"\nPágina\tÓtimo\tSecond Chance")
            todas_paginas = set(carregamentos_otimo.keys()).union(carregamentos_second_chance.keys())
            for pagina in todas_paginas:
                print(f"{pagina}\t{carregamentos_otimo.get(pagina, 0)}\t{carregamentos_second_chance.get(pagina, 0)}")

if __name__ == "__main__":
    # Lista de arquivos disponíveis
    arquivos = ['A.txt', 'B.txt', 'C.txt', 'Z.txt']

    # Exibindo as opções de arquivos
    print("Escolha um arquivo para simular:")
    for i, arquivo in enumerate(arquivos):
        print(f"[{i}] {arquivo}")

    # Solicitando que o usuário escolha um arquivo
    escolha = int(input("Digite o número do arquivo escolhido (0-3): "))

    # Verificando se a escolha é válida
    if 0 <= escolha < len(arquivos):
        arquivo_escolhido = arquivos[escolha]
        print(f"\nVocê escolheu o arquivo: {arquivo_escolhido}")
    else:
        print("Escolha inválida.")
        exit(1)

    # Rodando a simulação para cada tamanho de memória
    rodar_simulacao(f"../files/{arquivo_escolhido}")
