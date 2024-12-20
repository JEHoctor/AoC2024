package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var _1Cmd = &cobra.Command{
	Use:   "1",
	Short: "Historian Hysteria",
	Long:  `Advent of Code day one: Historian Hysteria`,
	Run:   day_one,
}

func init() {
	rootCmd.AddCommand(_1Cmd)
}

func day_one(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/1p1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	left_slice := []int{}
	right_slice := []int{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			log.Fatal("Invalid input")
		}
		a, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		left_slice = append(left_slice, a)
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		right_slice = append(right_slice, b)
	}

	sort.Ints(left_slice)
	sort.Ints(right_slice)

	total := 0
	for i, a := range left_slice {
		b := right_slice[i]
		// add the absolute value of b - a
		if b > a {
			total += b - a
		} else {
			total += a - b
		}
	}
	fmt.Println("Part one:", total)

	// Part 2 solution
	new_total := 0
	a_index, b_index := 0, 0
	for a_index < len(left_slice) {
		a := left_slice[a_index]
		count_a_left := 0
		for a_index < len(left_slice) && left_slice[a_index] == a {
			count_a_left++
			a_index++
		}
		for b_index < len(right_slice) && right_slice[b_index] < a {
			b_index++
		}
		count_a_right := 0
		for b_index < len(right_slice) && right_slice[b_index] == a {
			count_a_right++
			b_index++
		}
		new_total += a * count_a_left * count_a_right
	}
	fmt.Println("Part two:", new_total)
}
