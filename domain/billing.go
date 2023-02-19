package domain

import "time"

type Billing struct {
	Id          uint
	CreatedDate time.Time
	Total       float64
	Duration    int
}
