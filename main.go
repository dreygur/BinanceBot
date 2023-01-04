package main

import (
	"binancebot/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Stop printing error stack
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(0)
		}
	}()

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		rawString := scanner.Text()
		start := time.Now()
		utils.ProcessCommand(rawString)
		fmt.Printf("Time taken %v\n", time.Since(start))
	}
}
