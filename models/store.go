package models

import "time"

type Store struct {
    ID          int        `json:"id"`
    UUID        string     `json:"uuid"`
    UserID      int        `json:"user_id"`
    StoreName   string     `json:"store_name"`
    Address     string     `json:"address"`
    PhoneNumber string     `json:"phone_number"`
    Status      string     `json:"status"`
    DeletedAt   *time.Time `json:"deleted_at"`
    CreatedAt   time.Time  `json:"created_at"`
    ModifiedAt  time.Time  `json:"modified_at"`
}