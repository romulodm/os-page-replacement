package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"replacement/algorithms"
	"strings"
)

// Function to read page accesses from the file while tracking future accesses
func readAccesses(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var accesses []string
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			i++
			continue
		}

		if len(line) > 1 {
			page := line
			accesses = append(accesses, page)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Number of entries in file %s: %d\n\n", filepath.Base(filename), len(accesses))
	return accesses, nil
}

// Function to run the Second Chance algorithm in a goroutine
func runSecondChance(accesses []string, numFrames int, result chan<- int, loads chan<- map[string]int) {
	faults, loadsMap := algorithms.SecondChance(accesses, numFrames)
	result <- faults
	loads <- loadsMap
}

// Function to run the Optimal algorithm in a goroutine
func runOptimal(accesses []string, numFrames int, result chan<- int, loads chan<- map[string]int) {
	faults, loadsMap := algorithms.Optimal(accesses, numFrames)
	result <- faults
	loads <- loadsMap

}

func main() {
	// Path to the "files" folder
	dir := "../files"

	// Reads the files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading the folder:", err)
		return
	}

	// Lists the available files
	fmt.Println("Available files:")
	for i, file := range files {
		fmt.Printf("%d - %s\n", i+1, file.Name())
	}

	// Ask the user to select a file
	var fileChoice int
	fmt.Print("Choose the file (enter the number): ")
	_, err = fmt.Scanln(&fileChoice)
	if err != nil || fileChoice < 1 || fileChoice > len(files) {
		fmt.Println("Invalid choice.")
		return
	}

	// Full path to the selected file
	selectedFile := filepath.Join(dir, files[fileChoice-1].Name())
	fmt.Printf("You selected: %s\n", selectedFile)

	// Read accesses
	accesses, err := readAccesses(selectedFile)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Options for physical memory in KB (1 page = 4KB)
	memories := []int{1 * 1024 * 1024, 128 * 1024, 16 * 1024, 8 * 1024}

	// Lists the available memory options
	fmt.Println("Available memories (in KB):")
	for i, memory := range memories {
		fmt.Printf("%d - %d KB\n", i+1, memory)
	}
	fmt.Println("5 - Choose memory size manually")

	// Ask the user to choose the memory size
	var memoryChoice int
	fmt.Print("Choose the memory (enter the number): ")
	_, err = fmt.Scanln(&memoryChoice)
	if err != nil || memoryChoice < 1 || memoryChoice > 5 {
		fmt.Println("Invalid choice.")
		return
	}

	var numFrames int

	if memoryChoice == 5 {
		// If the user chooses to define manually
		fmt.Print("Enter the number of frames (memory size): ")
		_, err = fmt.Scanln(&numFrames)
		if err != nil || numFrames <= 0 {
			fmt.Println("Invalid choice.")
			return
		}
	} else {
		// Convert from KB to number of frames (1 frame = 4 KB)
		selectedMemory := memories[memoryChoice-1]
		numFrames = selectedMemory / 4
		fmt.Printf("You selected: %d KB of memory, which equals %d frames\n\n", selectedMemory, numFrames)
	}

	// Channels to receive the results of the algorithms
	chanSecondChance := make(chan int)
	chanOptimal := make(chan int)

	// Channels to receive the map of the page loads
	chanOptimalLoads := make(chan map[string]int)
	chanSecondChanceLoads := make(chan map[string]int)

	// Run the algorithms in parallel using goroutines
	go runSecondChance(accesses, numFrames, chanSecondChance, chanSecondChanceLoads)
	go runOptimal(accesses, numFrames, chanOptimal, chanOptimalLoads)

	// Receive the results
	faultsOptimal := <-chanOptimal
	fmt.Printf("Page faults (Optimal): %d\n", faultsOptimal)

	faultsSC := <-chanSecondChance
	fmt.Printf("Page faults (Second Chance): %d\n", faultsSC)

	// Receive the rpage loads
	loadsOptimal := <-chanOptimalLoads
	loadsSC := <-chanSecondChanceLoads

	// Ask if the user wants to see the number of loads per page
	var showLoads string
	fmt.Print("\nDo you want to see the number of loads per page? (y/n): ")
	_, err = fmt.Scanln(&showLoads)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	// If user wants to see the loads
	if strings.ToLower(showLoads) == "y" {
		fmt.Printf("\n%-10s %-10s %-10s\n", "Page", "Optimal", "Second Chance")
		fmt.Println(strings.Repeat("-", 30))
		for page, optimalLoads := range loadsOptimal {
			secondChanceLoads := loadsSC[page]
			fmt.Printf("%-10s %-10d %-10d\n", page, optimalLoads, secondChanceLoads)
		}
	}
}
