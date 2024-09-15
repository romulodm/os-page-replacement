## Page Replacement Algorithms
This project implements page replacement algorithms (Second Chance and Optimal) in Go. It simulates page faults using memory access patterns from input files.

## Requirements
Go version 1.16 or higher

## Project Structure
```
files/                    # Directory containing input files (access patterns)
├── A.txt                 # Example input file with memory accesses
├── B.txt                 # Another example input file
│
go/                       # Directory containing the Go code
├── algorithms/           # Directory containing the page replacement algorithms
│   ├── optimal.go        # Optimal page replacement algorithm
│   └── second_chance.go  # Second Chance page replacement algorithm
│
├── main.go               # Main entry point for the project
├── go.mod                # Go module file, defines the module and Go version
└── README.md             # Project instructions and how to run the project
```
## How to Run
in the go directory use `go run main.go`