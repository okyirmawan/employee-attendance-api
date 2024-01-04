package domain

import "time"

type Attendance struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement:true"`
	EmployeeID  uint64    `gorm:"not null"`
	CheckIn     time.Time `gorm:"not null"`
	CheckOut    time.Time
	TotalHours  uint8
	LocationIn  string
	LocationOut string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
