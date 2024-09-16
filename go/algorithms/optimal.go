package algorithms

// The Optimal page replacement algorithm predicts future page accesses
// and always replaces the page that won't be used for the longest time in the future
func Optimal(accesses []string, numFrames int) (int, map[string]int) {
	faults := 0
	loads := make(map[string]int) // Counter to count the loads of each page

	// Pre-process the next uses of each page
	// This step builds a map of where each page will be used next, to simulate future knowledge
	nextUsage := make(map[string][]int)
	for i, page := range accesses {
		nextUsage[page] = append(nextUsage[page], i)
	}

	// The 'memory' map tracks which pages are currently in memory.
	// 'memoryList' keeps the order of pages loaded into memory.
	memory := make(map[string]bool)
	memoryList := make([]string, 0, numFrames)

	for _, page := range accesses {
		// Remove the current index from the page's next usage list
		nextUsage[page] = nextUsage[page][1:]

		// If the page is already in memory, skip it
		if memory[page] {
			continue
		}

		// Page fault occurs if the page is not in memory
		faults++
		loads[page]++ // Increment the page counter

		// If memory is full, we need to replace a page
		if len(memoryList) >= numFrames {
			// Find the page that will not be used for the longest time in the future
			pageToRemove := ""
			farthestIndex := -1

			for _, p := range memoryList {
				if len(nextUsage[p]) == 0 {
					// This page will no longer be used, so we can safely replace it
					pageToRemove = p
					break
				} else if nextUsage[p][0] > farthestIndex {
					// This page will be used the farthest in the future, so it's the best candidate for replacement
					farthestIndex = nextUsage[p][0]
					pageToRemove = p
				}
			}

			// Remove the selected page from memory
			delete(memory, pageToRemove)
			// Remove the page from the list of pages in memory
			for idx, val := range memoryList {
				if val == pageToRemove {
					memoryList = append(memoryList[:idx], memoryList[idx+1:]...)
					break
				}
			}
		}

		// Add the new page to memory
		memory[page] = true
		memoryList = append(memoryList, page)
	}

	return faults, loads
}
