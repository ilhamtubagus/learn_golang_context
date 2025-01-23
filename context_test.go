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
