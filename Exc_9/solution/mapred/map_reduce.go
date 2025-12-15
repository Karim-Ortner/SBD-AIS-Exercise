package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

var cleanRe = regexp.MustCompile(`[^a-zA-Z]+`)

// Run executes the whole MapReduce pipeline
func (m *MapReduce) Run(input []string) map[string]int {
	mapCh := make(chan []KeyValue)
	var wg sync.WaitGroup

	// MAP PHASE (concurrent)
	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapCh <- m.wordCountMapper(text)
		}(line)
	}

	go func() {
		wg.Wait()
		close(mapCh)
	}()

	// GROUP BY KEY
	groups := make(map[string][]int)
	for kvPairs := range mapCh {
		for _, kv := range kvPairs {
			groups[kv.Key] = append(groups[kv.Key], kv.Value)
		}
	}

	// REDUCE PHASE
	result := make(map[string]int)
	for key, vals := range groups {
		reduced := m.wordCountReducer(key, vals)
		result[reduced.Key] = reduced.Value
	}

	return result
}

// Mapper: cleans text, splits to words, returns KeyValue pairs
func (m *MapReduce) wordCountMapper(text string) []KeyValue {
	clean := cleanRe.ReplaceAllString(strings.ToLower(text), " ")
	words := strings.Fields(clean)

	out := make([]KeyValue, 0, len(words))
	for _, w := range words {
		out = append(out, KeyValue{Key: w, Value: 1})
	}
	return out
}

// Reducer: sums up the counts
func (m *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
