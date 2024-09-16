package algorithms

// Page structure for the Second Chance algorithm
type Page struct {
	pageID     string
	referenced bool
}

// The Second Chance page replacement algorithm is a modified version of the FIFO (First-In-First-Out) algorithm;
// It gives each page a "second chance" before it is replaced, using a reference bit;
// If the reference bit is 1, the page is given a second chance and its bit is reset to 0;
// If the bit is 0, the page is replaced. This algorithm aims to balance simplicity and efficiency;
func SecondChance(accesses []string, numFrames int) (int, map[string]int) {
	// Initialize memory and related structures
	memory := make([]Page, numFrames) // Memory frames (numFrames)
	pageMap := make(map[string]int)   // To track if a page is in memory

	faults := 0                   // Page fault counter
	loads := make(map[string]int) // To track how many times each page is loaded

	pos := 0 // The pointer for replacement

	for _, page := range accesses {
		// Check if the page is already in memory
		if index, found := pageMap[page]; found {
			// Page is found in memory, update reference bit to true (second chance)
			memory[index].referenced = true
			continue
		}

		// Page fault occurs (page not found in memory)
		faults++
		loads[page]++ // Count how many times the page has been loaded

		// If memory is full, replace a page using the second chance algorithm
		for {
			// Check the reference bit at the current pointer position
			if !memory[pos].referenced {
				// Remove the old page from pageMap
				oldPage := memory[pos].pageID
				delete(pageMap, oldPage)

				// Add the new page to memory
				memory[pos] = Page{pageID: page, referenced: false}
				pageMap[page] = pos // Track the position of the new page in memory

				// Move pointer to the next position (circularly)
				pos = (pos + 1) % numFrames
				break
			} else {
				// If the reference bit is 1, reset it to 0 and give it a second chance
				memory[pos].referenced = false
				// Move pointer to the next position (circularly)
				pos = (pos + 1) % numFrames
			}
		}
	}

	// Return the number of page faults and the loads of each page
	return faults, loads
}
