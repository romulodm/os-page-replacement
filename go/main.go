package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"replacement/algorithms"
)

// Função para rodar o algoritmo Segunda Chance em uma goroutine
func runSecondChance(acessos []string, numFrames int, result chan<- int) {
	faltas := algorithms.SecondChance(acessos, numFrames)
	result <- faltas
}

// Função para rodar o algoritmo Ótimo em uma goroutine
func runOptimal(acessos []string, numFrames int, mapaFuturo map[string][]int, result chan<- int) {
	faltas := algorithms.Optimal(acessos, numFrames, mapaFuturo)
	result <- faltas
}

func lerAcessosComFuturo(filename string) ([]string, map[string][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var acessos []string
	mapaFuturo := make(map[string][]int)
	scanner := bufio.NewScanner(file)

	// Compile uma expressão regular para capturar o formato da página (I ou D seguido de números)
	re := regexp.MustCompile(`[A-Z]\d+`)

	// Itera sobre cada linha do arquivo
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Extrai a página no formato "I0", "D1", etc.
		match := re.FindString(line)
		if match != "" {
			// Armazena o acesso
			acessos = append(acessos, match)

			// Preenche o mapa de acessos futuros
			mapaFuturo[match] = append(mapaFuturo[match], i)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	fmt.Printf("Quantidade de registros no arquivo %s: %d\n\n", filepath.Base(filename), len(acessos))

	return acessos, mapaFuturo, nil
}

func main() {
	// Caminho para a pasta "files"
	dir := "../files"

	// Lê os arquivos dentro da pasta
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Erro ao ler a pasta:", err)
		return
	}

	// Lista os arquivos disponíveis
	fmt.Println("Arquivos disponíveis:")
	for i, file := range files {
		fmt.Printf("%d - %s\n", i+1, file.Name())
	}

	// Solicita ao usuário que escolha um arquivo
	var fileChoice int
	fmt.Print("Escolha o arquivo (digite o número): ")
	_, err = fmt.Scanln(&fileChoice)
	if err != nil || fileChoice < 1 || fileChoice > len(files) {
		fmt.Println("Escolha inválida.")
		return
	}

	// Caminho completo para o arquivo selecionado
	selectedFile := filepath.Join(dir, files[fileChoice-1].Name())
	fmt.Printf("Você selecionou: %s\n", selectedFile)

	// Lê acessos e constrói o mapa de acessos futuros ao mesmo tempo
	acessos, mapaFuturo, err := lerAcessosComFuturo(selectedFile)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Opções de memórias físicas em KB (1 página = 4KB)
	memorias := []int{1 * 1024 * 1024, 128 * 1024, 16 * 1024, 8 * 1024}

	// Lista as opções de memórias físicas
	fmt.Println("Memórias disponíveis (em KB):")
	for i, memoria := range memorias {
		fmt.Printf("%d - %d KB\n", i+1, memoria)
	}
	fmt.Println("5 - Escolher manualmente o tamanho da memória")

	// Solicita ao usuário que escolha o tamanho da memória
	var memoryChoice int
	fmt.Print("Escolha a memória (digite o número): ")
	_, err = fmt.Scanln(&memoryChoice)
	if err != nil || memoryChoice < 1 || memoryChoice > 5 {
		fmt.Println("Escolha inválida.")
		return
	}

	var numFrames int

	if memoryChoice == 5 {
		// Se o usuário escolher a opção de definir manualmente
		fmt.Print("Digite o número de frames (tamanho da memória): ")
		_, err = fmt.Scanln(&numFrames)
		if err != nil || numFrames <= 0 {
			fmt.Println("Escolha inválida.")
			return
		}
	} else {
		// Converte de KB para número de frames (1 frame = 4 KB)
		selectedMemory := memorias[memoryChoice-1]
		numFrames = selectedMemory / 4
		fmt.Printf("Você selecionou: %d KB de memória, que equivale a %d frames\n\n", selectedMemory, numFrames)
	}

	// Canais para receber os resultados dos algoritmos
	chanSecondChance := make(chan int)
	chanOptimal := make(chan int)

	// Executa os algoritmos em paralelo usando goroutines
	go runSecondChance(acessos, numFrames, chanSecondChance)
	go runOptimal(acessos, numFrames, mapaFuturo, chanOptimal)

	// Recebe os resultados
	faltasSC := <-chanSecondChance
	faltasOptimal := <-chanOptimal

	// Imprime os resultados
	fmt.Printf("Faltas de página (Segunda Chance): %d\n", faltasSC)
	fmt.Printf("Faltas de página (Ótimo): %d\n", faltasOptimal)
}
