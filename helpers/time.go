package helpers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/madjiebimaa/fcc-exercise-tracker-ms/models"
)

func StripedDateToTime(date string) (time.Time, error) {
	d := strings.Split(date, "-")
	var dateArr []int
	for _, v := range d {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
			return time.Time{}, models.ErrInternalServerError
		}

		dateArr = append(dateArr, val)
	}

	ti := time.Date(dateArr[0], time.Month(dateArr[1]), dateArr[2], 0, 0, 0, 0, time.UTC)
	return ti, nil
}
