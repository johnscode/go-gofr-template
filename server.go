package main

import (
	"go-gofr/migrations"
	"go-gofr/models"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

func main() {

	app := gofr.New()

	app.Migrate(migrations.All())

	app.GET("/health", func(c *gofr.Context) (interface{}, error) {
		status := struct {
			Status string `json:"status"`
		}{Status: "OK"}
		return response.Raw{Data: status}, nil
	})

	app.GET("/users", func(ctx *gofr.Context) (interface{}, error) {
		var users []models.User
		rows, err := ctx.SQL.Query("SELECT * FROM users")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.FirstName,
				&user.LastName, &user.Email, &user.HashedPass, &user.ApiKey, &user.CreatedAt, &user.UpdateAt, &user.LastLogin); err != nil {
				return nil, err
			}
			users = append(users, user)
		}
		return response.Raw{Data: users}, nil
	})

	app.Run()
}
