package migrations

import (
	"github.com/google/uuid"
	"go-gofr/models"
	"gofr.dev/pkg/gofr/migration"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func createUsersTableAndSeeds() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			err := SeedUsers(d.SQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func SeedUsers(sql migration.SQL) error {
	err := createUsersTable(sql)
	if err != nil {
		return err
	}

	// note that this is an example.
	// normally DO NOT include passwords in repo
	seedUsers := [][]string{
		{"god", "almighty", "admin@johnscode.com", "god"},
		{"dev", "johnscode", "dev@johnscode.com", "dev"},
		{"john", "code", "john@johnscode.com", "j"},
	}
	for _, seedUser := range seedUsers {
		_, err2 := CreateUser(sql, seedUser[0], seedUser[1], seedUser[2], seedUser[3])
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func CreateUser(sql migration.SQL, first_name string, last_name string, email string, clear string) (*models.User, error) {
	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(clear), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	strHashPass := string(hashedPassword)
	apikey := uuid.New().String()
	tm := time.Now()
	query := "INSERT INTO users (id, first_name, last_name, email, hashedPass, apiKey, createdAt, updatedAt, lastLogin) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err = sql.Exec(query, id.String(), first_name, last_name, email, strHashPass, apikey, tm, tm, tm)
	if err != nil {
		return nil, err
	}
	u := models.User{
		ID:         id.String(),
		FirstName:  first_name,
		LastName:   last_name,
		Email:      email,
		HashedPass: strHashPass,
		ApiKey:     apikey,
		CreatedAt:  tm,
		UpdateAt:   tm,
		LastLogin:  tm,
	}
	return &u, nil
}

func createUsersTable(sql migration.SQL) error {
	_, err := sql.Exec(`
        CREATE TABLE IF NOT EXISTS users (
			id varchar(256) PRIMARY KEY,
			first_name varchar(256), 
			last_name varchar(256), 
			email varchar(256), 
			hashedPass varchar(60), 
			apiKey varchar(40),
			createdAt TIMESTAMP WITH TIME ZONE,
			updatedAt TIMESTAMP WITH TIME ZONE,
			lastLogin TIMESTAMP WITH TIME ZONE
		);
    `)
	return err
}
