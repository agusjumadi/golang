package entities

import "time"

type Product struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
