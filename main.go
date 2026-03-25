package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func (m *Measurement) addTemp(temp float64) {
	m.Sum += temp
	m.Min = min(m.Min, temp)
	m.Max = max(m.Max, temp)
	m.Count++
}

func (m* Measurement) Avg() float64 {
	if m.Count == 0 {
		return 0
	}
	return m.Sum / float64(m.Count)
}

func main() {
	locations := make(map[string]*Measurement)
	file, err := os.Open("measurements.txt")
	if err != nil {
		fmt.Println("Measurements not found")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		index := strings.Index(line, ";")
		if index == -1 {
			continue
		}

		location := line[:index]
		rawTemp := line[index+1:]
		temp, err := strconv.ParseFloat(rawTemp, 64)
		if err != nil {
			continue
		}

		existisMeasurement, ok := locations[location]
		if !ok {
			locations[location] = &Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
			continue
		}
		existisMeasurement.addTemp(temp)
		
	}

	locationsSorted := make([]string, 0, len(locations))
	for key := range locations {
		locationsSorted = append(locationsSorted, key)
	}
	sort.Strings(locationsSorted)

	for _, key := range locationsSorted {
		elem := locations[key]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f\n", 
			key, 
			elem.Min, 
			elem.Avg(), 
			elem.Max,
		)
	}
}