// models.go - Định nghĩa models cho CMS Redis

package models

import "time"

type Location struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Game struct {
	ID          uint      `json:"id"`
	LocationID  uint      `json:"location_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type News struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Setting struct {
	ID         uint      `json:"id"`
	LocationID uint      `json:"location_id"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Banner struct {
	ID         uint      `json:"id"`
	LocationID uint      `json:"location_id"`
	Image      string    `json:"image"`
	Link       string    `json:"link"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}