package models

import "time"

type User struct {
	ID           int        `json:"id"`
	UUID         string     `json:"uuid"`
	Email        string     `json:"email"`
	Password     string     `json:"-"`
	NamaDepan    string     `json:"nama_depan"`
	NamaBelakang string     `json:"nama_belakang"`
	ImageURL     *string    `json:"image_url"`
	Role         string     `json:"role"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	ModifiedAt   time.Time  `json:"modified_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
