package models

import "github.com/jinzhu/gorm"

// User encapsulates information about a user account.
type User struct {
	AccountName   string
	DisplayName   string
	Email         string `gorm:"unique"`
	Administrator bool

	// Password is the bcrypt hashed password for this user.
	Password []byte

	// The videos that this user has uploaded
	Uploads []Video

	// ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
}

type Video struct {
	Title       string
	Uuid        string
	ContentType string
	Views       uint
	UserID      uint
	gorm.Model
}

// Subscription represents the subscription of one account
// to the contents of the other
type Subscription struct {
	// Subscriber is the user that IS subscribed
	Subscriber uint
	// Creator is the user that is susbcribed TO
	Creator uint

	gorm.Model
}
