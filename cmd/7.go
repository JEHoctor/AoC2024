package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var _7Cmd = &cobra.Command{
	Use:   "7",
	Short: "Bridge Repair",
	Long:  `Advent of Code day seven: Bridge Repair`,
	Run:   day_seven,
}

func init() {
	rootCmd.AddCommand(_7Cmd)
}

type Equation struct {
	target int
	values []int
}

func day_seven(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	equations := []Equation{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		if len(parts) != 2 {
			log.Fatal("Invalid input")
		}
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		values := []int{}
		for _, value := range strings.Fields(parts[1]) {
			v, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, v)
		}
		equations = append(equations, Equation{target, values})
	}

	total := 0
	for _, e := range equations {
		if could_be_true(e, false) {
			total += e.target
		}
	}
	fmt.Println("Part one:", total)

	total_two := 0
	for _, e := range equations {
		if could_be_true(e, true) {
			total_two += e.target
		}
	}
	fmt.Println("Part two:", total_two)
}

func could_be_true(eq Equation, use_concat bool) bool {
	if len(eq.values) == 0 {
		return false
	}
	if len(eq.values) == 1 {
		return eq.values[0] == eq.target
	}
	return (could_be_true(Equation{eq.target, append([]int{eq.values[0] * eq.values[1]}, eq.values[2:]...)}, use_concat) ||
		could_be_true(Equation{eq.target, append([]int{eq.values[0] + eq.values[1]}, eq.values[2:]...)}, use_concat) ||
		(use_concat &&
			could_be_true(Equation{eq.target, append([]int{concat_numbers(eq.values[0], eq.values[1])}, eq.values[2:]...)}, use_concat)))
}

func concat_numbers(x, y int) int {
	x_str := strconv.Itoa(x)
	y_str := strconv.Itoa(y)
	concat_str := x_str + y_str
	concat_values, err := strconv.Atoi(concat_str)
	if err != nil {
		log.Fatal(err)
	}
	return concat_values
}
