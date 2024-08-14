package models

// Response represents a response containing a set of integer solutions.
type Response struct {
	Solutions `json:"solutions"`
}

// Solutions is a two-dimensional integer slice representing the solutions.
// Each inner slice represents a single solution.
type Solutions [][]int
