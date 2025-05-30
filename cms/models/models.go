package models

import (
	"time"

	"gorm.io/gorm"
)

// Location model
type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Address   string    `json:"address" gorm:"type:varchar(255);not null"`
	Status    int       `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook for Location
func (l *Location) BeforeCreate(tx *gorm.DB) error {
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for Location
func (l *Location) BeforeUpdate(tx *gorm.DB) error {
	l.UpdatedAt = time.Now()
	return nil
}

// Game model
type Game struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	LocationID  uint      `json:"location_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Status      int       `json:"status" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Location    Location  `json:"location" gorm:"foreignKey:LocationID"`
}

// BeforeCreate hook for Game
func (g *Game) BeforeCreate(tx *gorm.DB) error {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for Game
func (g *Game) BeforeUpdate(tx *gorm.DB) error {
	g.UpdatedAt = time.Now()
	return nil
}

// News model
type News struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Status    int       `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook for News
func (n *News) BeforeCreate(tx *gorm.DB) error {
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for News
func (n *News) BeforeUpdate(tx *gorm.DB) error {
	n.UpdatedAt = time.Now()
	return nil
}

// Setting model
type Setting struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	LocationID uint      `json:"location_id" gorm:"not null"`
	Key        string    `json:"key" gorm:"type:varchar(255);not null;column:key"`
	Value      string    `json:"value" gorm:"type:varchar(5000);not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Location   Location  `json:"location" gorm:"foreignKey:LocationID"`
}

// BeforeCreate hook for Setting
func (s *Setting) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for Setting
func (s *Setting) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

// Banner model
type Banner struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	LocationID uint      `json:"location_id" gorm:"not null"`
	Image      string    `json:"image" gorm:"type:varchar(255);not null"`
	Link       string    `json:"link" gorm:"type:varchar(255);not null"`
	Status     int       `json:"status" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Location   Location  `json:"location" gorm:"foreignKey:LocationID"`
}

// BeforeCreate hook for Banner
func (b *Banner) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook for Banner
func (b *Banner) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

// Initialize database tables
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Location{},
		&Game{},
		&News{},
		&Setting{},
		&Banner{},
	)
}
