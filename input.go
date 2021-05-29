package main

//multiple module imports
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the birth year: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d in this year.", 2021 - input)
}