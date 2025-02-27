package main

import (
	"context"
	"fmt"

	"github.com/ziemianek/go-cards/services/cards/app"
)

func main() {
	a := app.New()
	err := a.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start app: %v", err)
	}
}
