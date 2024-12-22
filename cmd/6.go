package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var _6Cmd = &cobra.Command{
	Use:   "6",
	Short: "Guard Gallivant",
	Long:  `Advent of Code day six: Guard Gallivant`,
	Run:   day_six,
}

func init() {
	rootCmd.AddCommand(_6Cmd)
}

func day_six(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	guard_pos := [2]int{-1, -1}
	guard_direction := [2]int{-1, 0}
	obstructions := [][]bool{}
	row_number := 0
	for scanner.Scan() {
		obstruction_row := []bool{}
		for ic, c := range scanner.Text() {
			obstruction_row = append(obstruction_row, c == '#')
			if c == '^' {
				guard_pos = [2]int{row_number, ic}
			}
		}
		obstructions = append(obstructions, obstruction_row)
		row_number++
	}

	visited, looped := guard_movement(guard_pos, guard_direction, obstructions)
	if looped {
		log.Fatal("Looped unexpectedly before adding another obstruction")
	}

	total := 0
	for _, visited_row := range visited {
		for _, visited_cell := range visited_row {
			for _, visited_direction := range visited_cell {
				if visited_direction {
					total++
					break
				}
			}
		}
	}
	fmt.Println("Part one:", total)

	total_two := 0
	for i := 0; i < len(obstructions); i++ {
		for j := 0; j < len(obstructions[i]); j++ {
			if !obstructions[i][j] && !(i == guard_pos[0] && j == guard_pos[1]) && any4(visited[i][j]) {
				obstructions[i][j] = true
				_, looped := guard_movement(guard_pos, guard_direction, obstructions)
				if looped {
					total_two++
				}
				obstructions[i][j] = false
			}
		}
	}
	fmt.Println("Part two:", total_two)
}

func any4(visited [4]bool) bool {
	for _, v := range visited {
		if v {
			return true
		}
	}
	return false
}

func guard_movement(guard_pos [2]int, guard_direction [2]int, obstructions [][]bool) ([][][4]bool, bool) {
	visited := [][][4]bool{}
	for i := 0; i < len(obstructions); i++ {
		visited_row := [][4]bool{}
		for j := 0; j < len(obstructions[i]); j++ {
			visited_row = append(visited_row, [4]bool{false, false, false, false})
		}
		visited = append(visited, visited_row)
	}

	direction_index := 0
	for {
		if visited[guard_pos[0]][guard_pos[1]][direction_index] {
			return visited, true
		}
		visited[guard_pos[0]][guard_pos[1]][direction_index] = true
		next_guard_position := [2]int{guard_pos[0] + guard_direction[0], guard_pos[1] + guard_direction[1]}
		if next_guard_position[0] < 0 || next_guard_position[0] >= len(obstructions) || next_guard_position[1] < 0 || next_guard_position[1] >= len(obstructions[0]) {
			break
		}
		if obstructions[next_guard_position[0]][next_guard_position[1]] {
			guard_direction = [2]int{guard_direction[1], -guard_direction[0]}
			direction_index = (direction_index + 1) % 4
		} else {
			guard_pos = next_guard_position
		}
	}

	return visited, false
}
