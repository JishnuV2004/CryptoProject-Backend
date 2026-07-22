package webconfiguration

import "time"

type FeatureFlag struct {
	ID          uint      `gorm:"primaryKey"`
	FeatureName string    `gorm:"unique;not null"`
	Enabled     bool      `gorm:"default:true"`
	UpdatedAt   time.Time
}