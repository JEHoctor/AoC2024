package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var _3Cmd = &cobra.Command{
	Use:   "3",
	Short: "Mull It Over",
	Long:  `Advent of Code day three: Mull It Over`,
	Run:   day_three,
}

func init() {
	rootCmd.AddCommand(_3Cmd)
}

func day_three(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	total := 0
	for _, line := range lines {
		for _, match := range r.FindAllStringSubmatch(line, -1) {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			total += a * b
		}
	}
	fmt.Println("Part one:", total)

	r2 := regexp.MustCompile(`(do)\(\)|(don't)\(\)|(mul)\(([0-9]+),([0-9]+)\)`)
	total_two := 0
	enabled := true
	for _, line := range lines {
		for _, match := range r2.FindAllStringSubmatch(line, -1) {
			if match[1] == "do" {
				enabled = true
			} else if match[2] == "don't" {
				enabled = false
			} else if match[3] == "mul" && enabled {
				a, err := strconv.Atoi(match[4])
				if err != nil {
					log.Fatal(err)
				}
				b, err := strconv.Atoi(match[5])
				if err != nil {
					log.Fatal(err)
				}
				total_two += a * b
			}
		}
	}
	fmt.Println("Part two:", total_two)
}
