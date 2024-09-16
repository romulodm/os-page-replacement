package algorithms

// Page structure for the Second Chance algorithm
type Page struct {
	pageNumber string
	referenced bool
}

// The Second Chance page replacement algorithm is a modified version of the FIFO (First-In-First-Out) algorithm;
// It gives each page a "second chance" before it is replaced, using a reference bit;
// If the reference bit is 1, the page is given a second chance and its bit is reset to 0;
// If the bit is 0, the page is replaced. This algorithm aims to balance simplicity and efficiency;
func SecondChance(accesses []string, numFrames int) (int, map[string]int) {
	memory := make([]Page, numFrames)  // Initialize memory with a fixed size, tracking the pages and their reference bits
	pageMap := make(map[string]int)    // Map to track page positions in memory
	queue := make([]int, 0, numFrames) // Queue to manage page order

	faults := 0
	loads := make(map[string]int) // Counter to count the loads of each page

	pos := 0 // Pointer for the page to be replaced

	for _, page := range accesses {
		// Check if the page is already in memory
		if index, found := pageMap[page]; found {
			// If the page is already in memory, give it a second chance by setting the reference bit to true
			memory[index].referenced = true
			continue
		}

		// Page fault occurs if the page is not in memory
		faults++
		loads[page]++ // Increment the page counter

		// If memory is full, we need to replace a page
		if len(queue) >= numFrames {
			for {
				// Check the reference bit of the page at the pointer
				if !memory[queue[pos]].referenced {
					// If the reference bit is 0, replace this page
					pageToRemove := memory[queue[pos]].pageNumber
					delete(pageMap, pageToRemove)
					memory[queue[pos]] = Page{pageNumber: page, referenced: true}
					pageMap[page] = queue[pos]
					queue[pos] = (queue[pos] + 1) % numFrames // Move the pointer forward
					break
				} else {
					// If the reference bit is 1, reset it to 0 and give the page a second chance
					memory[queue[pos]].referenced = false
					queue[pos] = (queue[pos] + 1) % numFrames
				}
			}
		} else {
			// If there is space in memory, simply add the new page
			memory[len(queue)] = Page{pageNumber: page, referenced: true}
			pageMap[page] = len(queue)
			queue = append(queue, len(queue))
		}
	}

	return faults, loads
}
