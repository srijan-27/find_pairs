package models

// Input represents the input data for a problem.
//
// Numbers is a slice of integers to be processed.
// Target is the value of sum of the pair of numbers.
type Input struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}
