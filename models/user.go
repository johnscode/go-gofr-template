package models

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	HashedPass string    `json:"hashedPass"`
	ApiKey     string    `json:"apiKey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
	LastLogin  time.Time `json:"lastLogin"`
}

func CreateUser(ctx *gofr.Context, first_name string, last_name string, email string, clear string) (*User, error) {
	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clear), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	strHashPass := string(hashedPassword)
	apikey := uuid.New().String()
	tm := time.Now()
	query := "INSERT INTO users (id, first_name, last_name, email, hashedPass, apiKey) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = ctx.SQL.Exec(query, id.String(), first_name, last_name, email, strHashPass, apikey, tm, tm)
	if err != nil {
		return nil, err
	}
	u := User{
		ID:         id.String(),
		FirstName:  first_name,
		LastName:   last_name,
		Email:      email,
		HashedPass: strHashPass,
		ApiKey:     apikey,
		CreatedAt:  tm,
		UpdateAt:   tm,
	}
	return &u, nil
}
