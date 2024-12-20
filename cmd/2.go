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

var _2Cmd = &cobra.Command{
	Use:   "2",
	Short: "Red-Nosed Reports",
	Long:  `Advent of Code day two: Red-Nosed Reports`,
	Run:   day_two,
}

func init() {
	rootCmd.AddCommand(_2Cmd)
}

func is_safe(levels []int) bool {
	if len(levels) == 1 {
		return true
	}
	increasing := 0
	if levels[1] < levels[0] {
		increasing = -1
	} else {
		increasing = 1
	}
	for i := 0; i < len(levels)-1; i++ {
		if levels[i+1] == levels[i] {
			return false
		}
		q := increasing * (levels[i+1] - levels[i])
		if q < 0 || q > 3 {
			return false
		}
	}
	return true
}

func day_two(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	levels_slice := [][]int{}
	for scanner.Scan() {
		levels := []int{}
		fields := strings.Fields(scanner.Text())
		for _, field := range fields {
			level, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}
		levels_slice = append(levels_slice, levels)

		if is_safe(levels) {
			total++
		}
	}
	fmt.Println("Part one:", total)

	total_two := 0
	for _, levels := range levels_slice {
		if is_safe(levels) {
			total_two++
		} else {
			for i := 0; i < len(levels); i++ {
				modified_levels := append(append([]int{}, levels[:i]...), levels[i+1:]...)
				if is_safe(modified_levels) {
					total_two++
					break
				}
			}
		}
	}
	fmt.Println("Part two:", total_two)
}
