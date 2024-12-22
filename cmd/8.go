package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var _8Cmd = &cobra.Command{
	Use:   "8",
	Short: "Resonant Collinearity",
	Long:  `Advent of Code day eight: Resonant Collinearity`,
	Run:   day_eight,
}

func init() {
	rootCmd.AddCommand(_8Cmd)
}

func day_eight(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	arena_width := -1
	arena_height := 0
	antennae := map[rune][][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if arena_width == -1 {
			arena_width = len(line)
		}
		for ci, c := range line {
			if c != '.' {
				antenna_location := [2]int{arena_height, ci}
				antennae[c] = append(antennae[c], antenna_location)
			}
		}
		arena_height++
	}

	antinode_set := map[[2]int]bool{}
	for _, frequency_antennae := range antennae {
		for i, antenna := range frequency_antennae {
			for j := i + 1; j < len(frequency_antennae); j++ {
				other_antenna := frequency_antennae[j]
				dx, dy := other_antenna[0]-antenna[0], other_antenna[1]-antenna[1]
				antinode1 := [2]int{other_antenna[0] + dx, other_antenna[1] + dy}
				antinode2 := [2]int{antenna[0] - dx, antenna[1] - dy}
				if antinode1[0] >= 0 && antinode1[0] < arena_height && antinode1[1] >= 0 && antinode1[1] < arena_width {
					antinode_set[antinode1] = true
				}
				if antinode2[0] >= 0 && antinode2[0] < arena_height && antinode2[1] >= 0 && antinode2[1] < arena_width {
					antinode_set[antinode2] = true
				}
			}
		}
	}
	fmt.Println("Part one:", len(antinode_set))

	antinode_set2 := map[[2]int]bool{}
	for _, frequency_antennae := range antennae {
		for i, antenna := range frequency_antennae {
			for j := i + 1; j < len(frequency_antennae); j++ {
				other_antenna := frequency_antennae[j]
				dx, dy := other_antenna[0]-antenna[0], other_antenna[1]-antenna[1]
				dx, dy = remove_gcd(dx, dy)

				walking_antinode := antenna
				for walking_antinode[0] >= 0 && walking_antinode[0] < arena_height && walking_antinode[1] >= 0 && walking_antinode[1] < arena_width {
					antinode_set2[walking_antinode] = true
					walking_antinode = [2]int{walking_antinode[0] + dx, walking_antinode[1] + dy}
				}
				walking_antinode = [2]int{antenna[0] - dx, antenna[1] - dy}
				for walking_antinode[0] >= 0 && walking_antinode[0] < arena_height && walking_antinode[1] >= 0 && walking_antinode[1] < arena_width {
					antinode_set2[walking_antinode] = true
					walking_antinode = [2]int{walking_antinode[0] - dx, walking_antinode[1] - dy}
				}
			}
		}
	}
	fmt.Println("Part two:", len(antinode_set2))
}

func remove_gcd(a, b int) (int, int) {
	c, d := a, b
	for d != 0 {
		c, d = d, c%d
	}
	return a / c, b / c
}
