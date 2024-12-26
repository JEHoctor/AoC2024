package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var _9Cmd = &cobra.Command{
	Use:   "9",
	Short: "Disk Fragmenter",
	Long:  `Advent of Code day nine: Disk Fragmenter`,
	Run:   day_nine,
}

func init() {
	rootCmd.AddCommand(_9Cmd)
}

func day_nine(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/9.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 1 {
		log.Fatal("Invalid input")
	}

	next_file_id := 0
	next_run_is_empty_space := false
	block := 0
	file_runs := []FileRun{}
	for _, c := range lines[0] {
		n := int(c) - int('0')
		if next_run_is_empty_space {
			file_runs = append(file_runs, FileRun{size: n, id: -1, start_block: block})
			next_run_is_empty_space = false
		} else {
			file_runs = append(file_runs, FileRun{size: n, id: next_file_id, start_block: block})
			next_file_id++
			next_run_is_empty_space = true
		}
		block += n
	}

	new_file_runs := compact_files(file_runs)
	file_system_checksum := checksum(new_file_runs)

	fmt.Println("Part one:", file_system_checksum)

	new_file_runs2 := compact_files2(file_runs)
	file_system_checksum2 := checksum(new_file_runs2)

	fmt.Println("Part two:", file_system_checksum2)
}

type FileRun struct {
	size        int
	id          int // set to -1 for empty space
	start_block int
}

func append_file_run(runs []FileRun, size int, id int) []FileRun {
	if size == 0 {
		return runs
	}
	if len(runs) == 0 {
		return []FileRun{{size: size, id: id, start_block: 0}}
	}
	last_run := &runs[len(runs)-1]
	if last_run.id == id {
		last_run.size += size
		return runs
	} else {
		return append(runs, FileRun{size: size, id: id, start_block: last_run.start_block + last_run.size})
	}
}

func compact_files(runs []FileRun) []FileRun {
	right_hand := len(runs) - 1
	empty_blocks_at_end := 0
	for right_hand >= 0 && runs[right_hand].id == -1 {
		empty_blocks_at_end += runs[right_hand].size
		right_hand--
	}
	if right_hand == -1 {
		return []FileRun{{size: empty_blocks_at_end, id: -1, start_block: 0}}
	}
	new_runs := []FileRun{}
	right_hand_blocks_remaining := runs[right_hand].size

	for left_hand, run := range runs {
		if left_hand == right_hand {
			new_runs = append_file_run(new_runs, right_hand_blocks_remaining, run.id)
			break
		}
		if run.id != -1 {
			new_runs = append_file_run(new_runs, run.size, run.id)
		} else {
			for blocks_to_fill := run.size; blocks_to_fill > 0; {
				if right_hand_blocks_remaining > blocks_to_fill {
					new_runs = append_file_run(new_runs, blocks_to_fill, runs[right_hand].id)
					right_hand_blocks_remaining -= blocks_to_fill
					empty_blocks_at_end += blocks_to_fill
					blocks_to_fill = 0
				} else {
					new_runs = append_file_run(new_runs, right_hand_blocks_remaining, runs[right_hand].id)
					blocks_to_fill -= right_hand_blocks_remaining
					empty_blocks_at_end += right_hand_blocks_remaining
					right_hand_blocks_remaining = 0
					for right_hand > left_hand+1 && (runs[right_hand].id == -1 || right_hand_blocks_remaining == 0) {
						empty_blocks_at_end += right_hand_blocks_remaining
						right_hand--
						right_hand_blocks_remaining = runs[right_hand].size
					}
					if right_hand == left_hand+1 && (runs[right_hand].id == -1 || right_hand_blocks_remaining == 0) {
						blocks_to_fill = 0
					}
				}
			}
		}
	}

	new_runs = append_file_run(new_runs, empty_blocks_at_end, -1)
	return new_runs
}

func checksum(runs []FileRun) int {
	sum := 0
	for _, file_run := range runs {
		if file_run.id != -1 {
			total_block_positions := (2*file_run.start_block + file_run.size - 1) * file_run.size / 2
			sum += file_run.id * total_block_positions
		}
	}
	return sum
}

func compact_files2(runs []FileRun) []FileRun {
	left_runs := []FileRun{}
	right_runs := []FileRun{}
	left_hand := 0
	right_hand := len(runs) - 1
	for left_hand <= right_hand {
		if runs[left_hand].id != -1 {
			left_runs = append(left_runs, runs[left_hand])
			left_hand++
		}
		if runs[right_hand].id == -1 {
			right_runs = append(right_runs, runs[right_hand])
			right_hand--
		}
		if runs[left_hand].size >= runs[right_hand].size {

		} else {
			right_runs = append(right_runs, runs[right_hand])
			right_hand--
		}
	}
	for i := len(right_runs) - 1; i >= 0; i-- {
		left_runs = append(left_runs, right_runs[i])
	}
	return left_runs
}
