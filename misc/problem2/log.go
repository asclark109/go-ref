package problem2

import (
	"errors"
	"strconv"
	"strings"
)

// ProcessLog returns the total number of users from the log entries that are greater than the maxDuration time
// Note: Please make sure to read the assignment description for more details.
func ProcessLog(log []string, maxDuration int) int {
	// HIGH LEVEL ALGORITHM: (O(N) expected time via hashing)
	// (1) iterate thru log entries.
	//     hash OUT records to a map; hash IN records to a separate map
	// (2) iterate thru keys in one of the maps, if key also in other map,
	//     then person came in and out; increment counter if duration of
	//     stay is under maxDuration.

	// (1)
	inTimes := make(map[string]int) // key: visitor number string, val: timestamp
	outTimes := make(map[string]int)
	for _, line := range log {
		if visitorID, timestamp, action, err := interpretEntry(line); err == nil {
			if action == "IN" {
				inTimes[visitorID] = timestamp
			} else {
				outTimes[visitorID] = timestamp
			}
		}
	}

	// (2)
	count := 0
	for visitorID, inTimestamp := range inTimes {
		if outTimestamp, ok := outTimes[visitorID]; ok {
			// person came in and out
			if (outTimestamp - inTimestamp) <= maxDuration {
				count++
			}
		}
	}

	return count
}

// interpretEntry accepts a string of format: "**<visitor_num>**<timestamp>**<action>**",
// where ** represents any amount of whitespace,
// <visitor_num>, <timestamp>, <action> represent a supposed string, int, string,
// and returns the visitor_num, timestamp, action, error as a string, int, string, nil.
// On error, returns "", -1, "", error.
func interpretEntry(logLine string) (string, int, string, error) {
	strElems := strings.Split(logLine, ",")
	visitorNum := strings.TrimSpace(strElems[0])   // e.g. "199"
	timestampStr := strings.TrimSpace(strElems[1]) // e.g. "10"
	action := strings.TrimSpace(strElems[2])       // e.g. "OUT"

	if timestamp, err := strconv.Atoi(timestampStr); err == nil {
		return visitorNum, timestamp, action, nil
	}
	return "", -1, "", errors.New("could not interpret timestamp as an int")
}
