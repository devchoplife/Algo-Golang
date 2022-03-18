package main

import "fmt"

const ErrNotFound = "Search Query is not contained in the source"

func Linear(array []int, query int) (int, error) {
	for i, item := range array {
		if item == query {
			return i, nil
		}
	}
	return -1, fmt.Errorf(ErrNotFound)
}
