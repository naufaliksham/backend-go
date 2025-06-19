package models

import (
    "gorm.io/gorm"
    "time"
)

type Schedule struct {
    gorm.Model
    MovieTitle string
    Theater    string
    ShowTime   time.Time
}
