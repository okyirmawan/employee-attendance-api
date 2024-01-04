package domain

import "time"

type Gender string

const (
	LakiLaki  Gender = "Laki-laki"
	Perempuan Gender = "Perempuan"
)

type Employee struct {
	ID          uint64       `gorm:"primaryKey;autoIncrement:true"`
	Nip         string       `gorm:"not null;unique"`
	Name        string       `gorm:"not null"`
	DateOfBirth time.Time    `gorm:"type:date"`
	Gender      Gender       `gorm:"type:enum('Laki-laki','Perempuan')"`
	Username    string       `gorm:"not null;unique" json:"username"`
	Email       string       `gorm:"not null;unique" json:"email"`
	Password    string       `gorm:"not null;" json:"password"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	Attendance  []Attendance `gorm:"foreignKey:EmployeeID"`
}
