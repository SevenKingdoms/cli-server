package model

import (
)

type User struct {
	ID    int64  `json:"id"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	CreatedAt int64  `json:"createdAt"`
}
