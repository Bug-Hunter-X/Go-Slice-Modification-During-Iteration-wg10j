# Go Slice Modification During Iteration Bug

This repository demonstrates a common bug in Go where modifying a slice during iteration leads to unexpected behavior. The `bug.go` file contains the erroneous code, while `bugSolution.go` provides a corrected version.

The bug occurs because modifying the slice's length while iterating with a `for range` loop can cause the index to become out of bounds or skip elements.

The solution demonstrates a correct approach to removing elements from a slice, ensuring that the iteration is not affected by slice modifications.  Consider using a new slice to build the result.