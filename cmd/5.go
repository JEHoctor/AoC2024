package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var _5Cmd = &cobra.Command{
	Use:   "5",
	Short: "Print Queue",
	Long:  `Advent of Code day five: Print Queue`,
	Run:   day_five,
}

func init() {
	rootCmd.AddCommand(_5Cmd)
}

func day_five(cmd *cobra.Command, args []string) {
	f, err := os.Open("inputs/5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile(`^([0-9]+)\|([0-9]+)$`)

	rules := [][2]int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		match := r.FindStringSubmatch(scanner.Text())
		if len(match) != 3 {
			log.Fatal("Invalid input")
		}
		a, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, [2]int{a, b})
	}

	updates := [][]int{}
	for scanner.Scan() {
		update := []int{}
		for _, s := range strings.Split(scanner.Text(), ",") {
			a, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			update = append(update, a)
		}
		updates = append(updates, update)
	}

	total := 0
	total_two := 0
	for _, update := range updates {
		if obeys_all_rules(update, rules) {
			total += middle_number(update)
		} else {
			total_two += middle_number(fix_order(update, rules))
		}
	}
	fmt.Println("Part one:", total)
	fmt.Println("Part two:", total_two)
}

func fix_order(update []int, rules [][2]int) []int {
	if len(update) < 2 {
		return update
	}
	new_rules := filter_rules(update, rules)
	return fix_order_helper(update, new_rules)
}

func filter_rules(update []int, rules [][2]int) [][2]int {
	new_rules := [][2]int{}
	for _, rule := range rules {
		if index_of(update, rule[0], -1) > -1 && index_of(update, rule[1], -1) > -1 {
			new_rules = append(new_rules, rule)
		}
	}
	return new_rules
}

func fix_order_helper(update []int, rules [][2]int) []int {
	// We assume that len(update) >= 2.
	pivot := update[0]
	prefix := []int{pivot}
	suffix := update[1:]
	prefix_growing := true
	for prefix_growing {
		new_prefix := prefix
		new_suffix := []int{}
		for _, s := range suffix {
			goes_to_new_prefix := false
			for _, rule := range rules {
				if rule[0] == s && index_of(prefix, rule[1], -1) > -1 {
					new_prefix = append(new_prefix, s)
					goes_to_new_prefix = true
					break
				}
			}
			if !goes_to_new_prefix {
				new_suffix = append(new_suffix, s)
			}
		}
		prefix_growing = len(new_prefix) > len(prefix)
		prefix, suffix = new_prefix, new_suffix
	}
	return append(append(fix_order(prefix[1:], rules), pivot), fix_order(suffix, rules)...)
}

func obeys_rule(update []int, rule [2]int) bool {
	return index_of(update, rule[0], -1) < index_of(update, rule[1], len(update))
}

func obeys_all_rules(update []int, rules [][2]int) bool {
	for _, rule := range rules {
		if !obeys_rule(update, rule) {
			return false
		}
	}
	return true
}

func index_of(slice []int, target int, sentinel int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return sentinel
}

func middle_number(slice []int) int {
	return slice[len(slice)/2]
}
