package main

import (
	"sync"
	"time"
)

var lock sync.Mutex

// isuUUID=>[]condition
var currentHourConditions = map[string][]IsuCondition{}
var currentHour time.Time

// Returns rows to insert now
func addIsuConditionToPool(cond IsuCondition) []IsuCondition {
	lock.Lock()
	defer lock.Unlock()
	hour := cond.Timestamp.Truncate(time.Hour)
	if hour == currentHour {
		currentHourConditions[cond.JIAIsuUUID] = append(currentHourConditions[cond.JIAIsuUUID], cond)
		// fmt.Printf("was same len %d\n", len(currentHourConditions[cond.JIAIsuUUID]))
		return nil
	} else {
		rowsToInsert := append([]IsuCondition{}, currentHourConditions[cond.JIAIsuUUID]...)

		currentHour = hour
		currentHourConditions[cond.JIAIsuUUID] = []IsuCondition{cond}

		// fmt.Printf("was different len %d\n", len(rowsToInsert))

		return rowsToInsert
	}
}

func getLatestCoditionForEachIsu() []IsuCondition {
	return nil
}
