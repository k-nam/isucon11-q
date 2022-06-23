package main

import (
	"sync"
	"time"
)

var lock sync.Mutex

// isuUUID=>[]condition
var currentHourConditions = map[string][]IsuCondition{}

// isuUUID=>current hour
var currentHour = map[string]time.Time{}

// Returns rows to insert now
func addIsuConditionToPool(cond IsuCondition) []IsuCondition {
	lock.Lock()
	defer lock.Unlock()
	hour := cond.Timestamp.Truncate(time.Hour)
	if hour == currentHour[cond.JIAIsuUUID] {
		if len(currentHourConditions[cond.JIAIsuUUID]) > 10 {
			return nil
		}
		currentHourConditions[cond.JIAIsuUUID] = append(currentHourConditions[cond.JIAIsuUUID], cond)
		// fmt.Printf("was same len %d\n", len(currentHourConditions[cond.JIAIsuUUID]))
		return nil
	} else {
		rowsToInsert := append([]IsuCondition{}, currentHourConditions[cond.JIAIsuUUID]...)

		currentHour[cond.JIAIsuUUID] = hour
		currentHourConditions[cond.JIAIsuUUID] = []IsuCondition{cond}

		// fmt.Printf("was different len %d\n", len(rowsToInsert))

		return rowsToInsert
	}
}

func getLatestCoditionForEachIsu() []IsuCondition {
	return nil
}

var lock2 sync.Mutex

var isus = map[string]*Isu{}

func getIsu(uuid string) *Isu {
	lock.Lock()
	defer lock.Unlock()
	result := isus[uuid]
	return result
}

func addIsu(newIsu Isu) {
	lock.Lock()
	defer lock.Unlock()
	isus[newIsu.JIAIsuUUID] = &newIsu
}