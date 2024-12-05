package main

import (
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := os.ReadFile("./part01_input.txt")
	rules, updates := parse(string(b))
	ok, nok := satisfied(updates, rules)
	log.Printf("Part 1: %d", sum(ok))
	log.Printf("Part 2: %d", sum(correct(nok, rules)))
}

type Order struct {
	first  int
	second int
}

type Update = []int

// Sums the middle element of each update
func sum(updates []Update) int {
	total := 0
	for _, update := range updates {
		total += update[len(update)/2]
	}
	return total
}

// Corrects the updates according to the rules
func correct(updates []Update, rules []Order) []Update {
	corrected := make([]Update, 0)
	for _, update := range updates {
		corrected = append(corrected, order(update, rules))
	}
	return corrected
}

// Splits the updates into satisfied and not satisfied according to the rules
func satisfied(updates []Update, rules []Order) ([]Update, []Update) {
	ok, nok := make([]Update, 0), make([]Update, 0)
	for _, update := range updates {
		sorted := order(update, rules)
		if reflect.DeepEqual(update, sorted) {
			ok = append(ok, update)
		} else {
			nok = append(nok, update)
		}
	}
	return ok, nok
}

// Sort the update according to the rules
func order(update Update, rules []Order) Update {
	sorted := make([]int, len(update))
	copy(sorted, update)
	sort.Slice(sorted, func(i, j int) bool {
		for _, rule := range rules {
			if sorted[i] == rule.first && sorted[j] == rule.second {
				return true
			}
		}
		return false
	})
	return sorted
}

func parse(input string) (rules []Order, updates []Update) {
	rules = make([]Order, 0)
	updates = make([]Update, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Contains(line, "|") {
			split := strings.Split(strings.TrimSpace(line), "|")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			rules = append(rules, Order{x, y})
		}
		if strings.Contains(line, ",") {
			split := strings.Split(line, ",")
			update := make(Update, len(split))
			for i, s := range strings.Split(line, ",") {
				x, _ := strconv.Atoi(strings.TrimSpace(s))
				update[i] = x
			}
			updates = append(updates, update)
		}
	}

	return
}
