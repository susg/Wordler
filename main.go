package main

import (
	"fmt"
	"github.com/susg/wordler/internal/orchestrator"
)

func main() {
	game := orchestrator.NewWordleManager()
	for {
		var inp []string
		for i := 0; i < 5; i++ {
			var str string
			fmt.Scanf("%s", &str)
			inp = append(inp, str)
		}
		fmt.Println(game.Recommend(inp))

		fmt.Printf("\n\n1. Continue\n2. Start new game\n0. Exit\n")
		var x int
		fmt.Scanf("%d", &x)
		if x == 0 {
			break
		} else if x == 2 {
			game = orchestrator.NewWordleManager()
		}
	}
}
