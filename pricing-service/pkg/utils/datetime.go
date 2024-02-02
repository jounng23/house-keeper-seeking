package utils

import (
	"fmt"
	"pricing-svc/pkg/constant"
	"time"
)

func IsSpecialDay(date time.Time) bool {
	if _, ok := constant.SpecialDays[fmt.Sprintf("%v-%v", date.Day(), date.Month().String())]; ok {
		return true
	}
	return false
}
