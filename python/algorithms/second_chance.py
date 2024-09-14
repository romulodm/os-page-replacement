from collections import deque

def second_chance(acessos, num_quadros):
    memoria = deque()
    faltas = 0
    bit_referencia = {}
    contador_acessos = {}

    for pagina in acessos:
        # Contador de acessos por página
        if pagina in contador_acessos:
            contador_acessos[pagina] += 1
        else:
            contador_acessos[pagina] = 1
        
        if pagina in bit_referencia:
            bit_referencia[pagina] = 1  # Segunda chance
        else:
            faltas += 1
            if len(memoria) >= num_quadros:
                while True:
                    pag, _ = memoria[0]
                    if bit_referencia[pag] == 0:
                        memoria.popleft()
                        del bit_referencia[pag]
                        break
                    else:
                        bit_referencia[pag] = 0
                        memoria.rotate(-1)

            memoria.append((pagina, 0))
            bit_referencia[pagina] = 0

    print("Número de acessos por página:")
    for pagina, count in contador_acessos.items():
        print(f"Página {pagina}: {count} acessos")

    return faltas
