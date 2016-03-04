package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func normalizeHour(hour float64) float64 {
	if hour > 12 {
		return hour - 12
	} else if hour == 12 {
		return 0
	}

	return hour
}

func normalizePositiveValues(hour, minute int) (hourFloat, minuteFloat float64) {
	hourFloat = float64(hour)
	minuteFloat = float64(minute)
	if hourFloat < 0.0 {
		hourFloat = math.Abs(hourFloat)
	}
	if minuteFloat < 0.0 {
		minuteFloat = math.Abs(minuteFloat)
	}
	return
}

func calculateAngle(h, m int) float64 {
	hour, minute := normalizePositiveValues(h, m)
	hour = normalizeHour(hour)
	hourAngle := (30.0*hour + 0.5*minute)
	minuteAngle := 6.0 * minute
	angle := hourAngle - minuteAngle
	return math.Abs(float64(angle))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	hr, min, _ := t.Clock()
	fmt.Fprintf(w, "Current Time: %d:%d \n", hr, min)
	fmt.Fprintf(w, "Angle between clockwise: %fº\n", calculateAngle(hr, min))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
