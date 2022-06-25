package main

import (
	"sync"
	"time"
)

var lock sync.Mutex

// isuUUID=>[]condition
var currentHourConditions = map[string][]IsuCondition{}
var latestConditions = map[string]IsuCondition{}

// isuUUID=>current hour
var currentHour = map[string]time.Time{}

var rowsToInsert []IsuCondition

// Returns rows to insert now
func addIsuConditionToPool(cond IsuCondition) []IsuCondition {
	lock.Lock()
	defer lock.Unlock()

	if cond.Timestamp.After(latestConditions[cond.JIAIsuUUID].Timestamp) {
		latestConditions[cond.JIAIsuUUID] = cond
	}

	hour := cond.Timestamp.Truncate(time.Hour)
	if hour == currentHour[cond.JIAIsuUUID] {
		if len(currentHourConditions[cond.JIAIsuUUID]) > 5 {
			return nil
		}
		currentHourConditions[cond.JIAIsuUUID] = append(currentHourConditions[cond.JIAIsuUUID], cond)
		// fmt.Printf("was same len %d\n", len(currentHourConditions[cond.JIAIsuUUID]))
		return nil
	} else {
		rows := append([]IsuCondition{}, currentHourConditions[cond.JIAIsuUUID]...)

		currentHour[cond.JIAIsuUUID] = hour
		currentHourConditions[cond.JIAIsuUUID] = []IsuCondition{cond}

		// fmt.Printf("was different len %d\n", len(rowsToInsert))
		rowsToInsert = append(rowsToInsert, rows...)
		if len(rowsToInsert) > 1000 {
			copy := rowsToInsert
			rowsToInsert = []IsuCondition{}
			return copy
		} else {
			return nil
		}

	}
}

func getLatestConditions() []IsuCondition {
	lock.Lock()
	defer lock.Unlock()

	conditions := []IsuCondition{}
	for _, cond := range latestConditions {
		conditions = append(conditions, cond)
	}
	// fmt.Printf("trend len %d\n", len(latestConditions))
	return conditions
}

func getLatestConditionsAsMap(userId string) map[string]IsuCondition {
	lock.Lock()
	defer lock.Unlock()

	result := map[string]IsuCondition{}
	for isuUUID, cond := range latestConditions {
		// if cond.UserId == userId {
		result[isuUUID] = cond
		// }
	}
	return result
}

var lock2 sync.Mutex

var isus = map[string]Isu{}

func getIsu(uuid string) *Isu {
	lock.Lock()
	defer lock.Unlock()
	result := isus[uuid]
	return &result
}

func getIsusForUser(userId string) []Isu {
	lock.Lock()
	defer lock.Unlock()
	isuList := []Isu{}
	for _, isu := range isus {
		if isu.JIAUserID == userId {
			isuList = append(isuList, isu)
		}
	}
	return isuList
}

func loadLatestConditionFromDb() error {
	// fmt.Println("loadLatestConditionFromDb")
	latestConditions = map[string]IsuCondition{}

	conds := []IsuCondition{}

	err := db.Select(&conds,
		"SELECT a.character, a.id as isu_id, a.jia_user_id, b.id, b.message, b.timestamp, b.condition, b.jia_isu_uuid, b.is_sitting, b.created_at FROM isu a INNER JOIN isu_condition b ON a.jia_isu_uuid = b.jia_isu_uuid WHERE b.timestamp = (SELECT timestamp FROM isu_condition WHERE jia_isu_uuid = a.jia_isu_uuid ORDER BY timestamp DESC limit 1)")
	if err != nil {
		return err
	}

	for _, cond := range conds {
		latestConditions[cond.JIAIsuUUID] = cond
		// fmt.Printf("loading cond: %v", cond)
	}
	// fmt.Printf("loaded # : %d\n", len(latestConditions))
	return nil
}

func addIsu(newIsu Isu) {
	lock.Lock()
	defer lock.Unlock()
	isus[newIsu.JIAIsuUUID] = newIsu
}
