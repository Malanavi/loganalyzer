package analyzer

import (
	"sort"
	"strings"
)

type StatsResult struct {
	Total int
	Info  int
	Warn  int
	Error int
}

type TopLogItem struct {
	Log   string
	Times int
}

const (
	errorPrefix = "[ERROR]"
	warnPrefix  = "[WARN]"
	infoPrefix  = "[INFO]"
)

func Stats(lines []string) StatsResult {
	var s StatsResult

	for _, l := range lines {
		s.Total++

		switch {
		case strings.HasPrefix(l, errorPrefix):
			s.Error++
		case strings.HasPrefix(l, warnPrefix):
			s.Warn++
		case strings.HasPrefix(l, infoPrefix):
			s.Info++
		}
	}

	return s
}

func Errors(lines []string) []string {
	errs := make([]string, 0, len(lines)/2)

	for _, l := range lines {
		if strings.HasPrefix(l, errorPrefix) {
			errs = append(errs, l)
		}
	}

	return errs
}

func Top(lines []string) []TopLogItem {
	top := make([]TopLogItem, 0, 10)
	m := make(map[string]int, 10)

	for _, l := range lines {
		m[l]++
	}

	for l, t := range m {
		top = append(top, TopLogItem{l, t})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].Times > top[j].Times
	})

	return top
}
