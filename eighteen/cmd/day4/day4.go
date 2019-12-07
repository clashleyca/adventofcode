package main

import (
	"eighteen/internal/parser"
	"fmt"
	"sort"
	"strings"
)

type closetGuard struct {
	id          int
	startSleeps []int
	endSleeps   []int
}

func (c *closetGuard) timeAsleep() (timeAsleep int) {
	for i, startSleep := range c.startSleeps {
		endSleep := c.endSleeps[i]
		timeAsleep = timeAsleep + (endSleep - startSleep)
	}
	return timeAsleep
}

func (c *closetGuard) addStart(newsleep int) {
	c.startSleeps = append(c.startSleeps, newsleep)
}

func (c *closetGuard) addEnd(newsleep int) {
	c.endSleeps = append(c.endSleeps, newsleep)
}

func main() {
	strnarr := parser.ParseStrings("day4.txt")
	//strnarr := parser.ParseStrings("day4test.txt")
	sliceSort(strnarr)
	fmt.Println("lazy guard result", getLazyGuard(strnarr))
}

func minuteMap(startMinute, endMinute int, minMap *map[int]int) {
	for i := startMinute; i < endMinute; i++ {
		(*minMap)[i] = (*minMap)[i] + 1
	}
}

func getLazyGuard(closetLogs []string) int {

	var allTheGuards []closetGuard
	var currentGuardIndex int
	for _, logEntry := range closetLogs {
		guard, found := getAGuard(logEntry)
		if found {
			oldGuard := false
			for i, g := range allTheGuards {
				if g.id == guard.id {
					currentGuardIndex = i
					oldGuard = true
				}
			}
			if !oldGuard {
				allTheGuards = append(allTheGuards, guard)
				currentGuardIndex = len(allTheGuards) - 1
			}
			continue // go to next log entry
		}
		startSleep := getStartSleep(logEntry)
		if startSleep >= 0 {
			allTheGuards[currentGuardIndex].addStart(startSleep)
			continue
		}
		endSleep := getEndSleep(logEntry)
		if endSleep >= 0 {
			allTheGuards[currentGuardIndex].addEnd(endSleep)
		}
	}
	//var maxSleep int
	//var lazyGuard closetGuard
	//for _, g := range allTheGuards {
	//	if g.timeAsleep() > maxSleep {
	//		lazyGuard = g
	//		maxSleep = g.timeAsleep()
	//	}
	//}

	bestMin := 0
	mostFreq := 0
	mostGID := 0
	bestistMin := 0
	for _, g := range allTheGuards {
		var freq int
		bestMin, freq = getMostMin(g)
		if freq > mostFreq {
			mostFreq = freq
			mostGID = g.id
			bestistMin = bestMin
		}
	}

	//bestMin, freq := getMostMin(lazyGuard)
	return mostGID * bestistMin
}

func getMostMin(lazyGuard closetGuard) (mostMin, frequency int) {
	minMap := make(map[int]int)
	for i, _ := range lazyGuard.startSleeps {
		minuteMap(lazyGuard.startSleeps[i], lazyGuard.endSleeps[i], &minMap)
	}
	for i, _ := range minMap {
		if minMap[i] > frequency {
			frequency = minMap[i]
			mostMin = i
		}
	}
	return
}

func getAGuard(logEntry string) (guard closetGuard, found bool) {
	if !strings.Contains(logEntry, "Guard") {
		return closetGuard{}, false
	}
	var year, month, day, hour, minute, guardID int
	fmt.Sscanf(logEntry, "[%d-%d-%d %d:%d] Guard #%d", &year, &month, &day, &hour, &minute, &guardID)
	fmt.Println(year, month, day, hour, minute, guardID)
	return closetGuard{id: guardID}, true
}

func getStartSleep(logEntry string) (sleepMinute int) {

	if !strings.Contains(logEntry, "falls asleep") {
		return -1
	}
	var year, month, day, hour, minute, guardNum int
	fmt.Sscanf(logEntry, "[%d-%d-%d %d:%d] falls asleep", &year, &month, &day, &hour, &sleepMinute)
	fmt.Println(year, month, day, hour, sleepMinute)
	fmt.Println(year, month, day, hour, minute, guardNum)
	return sleepMinute
}

func getEndSleep(logEntry string) (sleepMinute int) {

	if !strings.Contains(logEntry, "wakes up") {
		return -1
	}
	var year, month, day, hour, minute, guardNum int
	fmt.Sscanf(logEntry, "[%d-%d-%d %d:%d] wakes up", &year, &month, &day, &hour, &sleepMinute)
	fmt.Println(year, month, day, hour, sleepMinute)
	fmt.Println(year, month, day, hour, minute, guardNum)
	return sleepMinute
}

/// sorting slices

// AxisSorter sorts planets by axis.
type AxisSorter []Planet

func (a AxisSorter) Len() int           { return len(a) }
func (a AxisSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AxisSorter) Less(i, j int) bool { return a[i].Axis < a[j].Axis }

// NameSorter sorts planets by name.
type NameSorter []Planet

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// NameSorter sorts planets by name.
type DateSorter []string

func (d DateSorter) Len() int { return len(d) }
func (d DateSorter) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d DateSorter) Less(l, r int) bool {
	var lYear, lMonth, lDay, lHour, lMinute int
	fmt.Sscanf(d[l], "[%d-%d-%d %d:%d]", &lYear, &lMonth, &lDay, &lHour, &lMinute)

	var rYear, rMonth, rDay, rHour, rMinute int
	fmt.Sscanf(d[r], "[%d-%d-%d %d:%d]", &rYear, &rMonth, &rDay, &rHour, &rMinute)

	if lYear != rYear {
		return lYear < rYear
	}
	if lMonth != rMonth {
		return lMonth < rMonth
	}
	if lDay != rDay {
		return lDay < rDay
	}
	if lHour != rHour {
		return lHour < rHour
	}
	if lMinute != rMinute {
		return lMinute < rMinute
	}
	return false
}

type Planet struct {
	Name       string  `json:"name"`
	Aphelion   float64 `json:"aphelion"`   // in million km
	Perihelion float64 `json:"perihelion"` // in million km
	Axis       int64   `json:"Axis"`       // in km
	Radius     float64 `json:"radius"`
}

func sliceSort(inarr []string) (outarr []string) {
	//var mars Planet
	//mars.Name = "Mars"
	//mars.Aphelion = 249.2
	//mars.Perihelion = 206.7
	//mars.Axis = 227939100
	//mars.Radius = 3389.5
	//
	//var earth Planet
	//earth.Name = "Earth"
	//earth.Aphelion = 151.930
	//earth.Perihelion = 147.095
	//earth.Axis = 149598261
	//earth.Radius = 6371.0
	//
	//var venus Planet
	//venus.Name = "Venus"
	//venus.Aphelion = 108.939
	//venus.Perihelion = 107.477
	//venus.Axis = 108208000
	//venus.Radius = 6051.8

	//log.Println("unsorted:", inarr)

	outarr = inarr
	sort.Sort(DateSorter(outarr))
	//log.Println("by axis:", outarr)
	return outarr

}
