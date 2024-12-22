package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var _4Cmd = &cobra.Command{
	Use:   "4",
	Short: "Ceres Search",
	Long:  `Advent of Code day four: Ceres Search`,
	Run:   day_four,
}

func init() {
	rootCmd.AddCommand(_4Cmd)
}

func day_four(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	remainder := "MAS"
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	total := 0
	for i, line := range lines {
		for j, letter := range line {
			if letter == 'X' {
				for _, direction := range directions {
					match := true
					s, t := i, j
					ds, dt := direction[0], direction[1]
					for _, rc := range remainder {
						s += ds
						t += dt
						if s < 0 || s >= len(lines) || t < 0 || t >= len(line) {
							match = false
							break
						}
						if rune(lines[s][t]) != rc {
							match = false
							break
						}
					}
					if match {
						total++
					}
				}
			}
		}
	}
	fmt.Println("Part one:", total)

	total_two := 0
	directions_two := [][]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	for i, line := range lines {
		for j, letter := range line {
			if letter == 'A' {
				is_xmas := x_mas(lines, i, j, directions_two)
				total_two += is_xmas
			}
		}
	}
	fmt.Println("Part two:", total_two)
}

func check_with_bounds(lines []string, i int, j int, c rune) bool {
	if i < 0 || i >= len(lines) {
		return false
	}
	if j < 0 || j >= len(lines[i]) {
		return false
	}
	return rune(lines[i][j]) == c
}

func x_mas(lines []string, i int, j int, directions [][]int) int {
	mas_match_count := 0
	for _, direction := range directions {
		di, dj := direction[0], direction[1]
		if check_with_bounds(lines, i+di, j+dj, 'M') && check_with_bounds(lines, i-di, j-dj, 'S') {
			mas_match_count++
			if mas_match_count == 2 {
				return 1
			}
		}
	}
	return 0
}
