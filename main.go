package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

var sequence[] int

func fibonacci(num int) (fib[] int) {
	sequence = append(sequence, 0)
	sequence = append(sequence, 1)
	if (num == 0) {
		return
	} else if (num == 1) {
		return sequence[:len(sequence) - 1]
	} else{
		for i := 2; ; i++ {	
			sequence = append(sequence, sequence[i-2] + sequence [i-1])
			if sequence[i] > num {
				sequence = sequence[:len(sequence) - 1]
				break
			}
		}
	}
	return sequence
}

func main() {
	fmt.Println("The program generates Fibonacci sequence, up to the number, read from console\nEnter the number to generate Fibonacci sequence up to, pls")
	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("error reading string",  err)
	}

	num = strings.Trim(num, "\n")
	number, err := strconv.Atoi(num)

	if err != nil {
		log.Fatalln("error converting str to int32", err)
	}

	for  _, v := range fibonacci(number) {
		fmt.Println(strconv.Itoa(v))
	}
}