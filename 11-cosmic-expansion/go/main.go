package main

import (
	"fmt"
	"os"

	"github.com/andrewwdye/adventofcode2023/11-cosmic-expansion/go/galaxy"
)

func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var result int
	switch os.Args[1] {
	case "1":
		result, err = galaxy.Solve1(f)
	default:
		err = fmt.Errorf("invalid argument")
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
