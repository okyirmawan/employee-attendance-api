package domain

import "time"

type Gender string

const (
	LakiLaki  Gender = "Laki-laki"
	Perempuan Gender = "Perempuan"
)

type Employee struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true"`
	Nip         string `gorm:"not null"`
	Name        string
	DateOfBirth time.Time `gorm:"type:date"`
	Gender      Gender    `gorm:"type:enum('Laki-laki','Perempuan')"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
