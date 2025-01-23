package learn_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	backgroundContext := context.Background()
	fmt.Println(backgroundContext)

	todoContext := context.TODO()
	fmt.Println(todoContext)
}

func TestWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) // Output: F
	fmt.Println(contextF.Value("c")) // Output: C
	fmt.Println(contextF.Value("b")) // Output: nil, different parent A -> C -> F
	fmt.Println(contextA.Value("b")) // Output: nil, parent cant invoke a child value
}
