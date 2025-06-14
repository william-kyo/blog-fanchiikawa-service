package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"blog-fanchiikawa-service/db"
	"blog-fanchiikawa-service/graph/model"
	"blog-fanchiikawa-service/greetings"
	"blog-fanchiikawa-service/sdk"
	"context"
	"database/sql"
	"log"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.User, error) {
	// If email is in the database, return user info
	row := db.MySQL.QueryRow("SELECT * FROM user WHERE email = ?", input.Email)
	var user model.User
	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == nil {
		// User exists, return the user
		return &user, nil
	} else if err != sql.ErrNoRows {
		// If error is not "no rows", return the error
		return nil, err
	}

	// User doesn't exist, create new user
	var newUser model.User
	newUser.Nickname = input.Nickname
	newUser.Email = input.Email

	// Start transaction
	tx, err := db.MySQL.Begin()
	if err != nil {
		return nil, err
	}
	// Defer rollback in case of error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert user
	result, err := tx.Exec("INSERT INTO user (nickname, email) VALUES (?, ?)", newUser.Nickname, newUser.Email)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Insert device
	_, err = tx.Exec("INSERT INTO user_device (user_id, device_id) VALUES (?, ?)", lastInsertId, input.DeviceID)
	if err != nil {
		return nil, err
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	message, _ := greetings.Hello(newUser.Nickname)
	log.Println(message)

	// Get the newly created user
	row = db.MySQL.QueryRow("SELECT * FROM user WHERE id = ?", lastInsertId)
	err = row.Scan(&user.ID, &user.Nickname, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DetectLanguage is the resolver for the detectLanguage field.
func (r *mutationResolver) DetectLanguage(ctx context.Context, input string) (string, error) {
	language, err := sdk.DetectLanguage(input)
	if err != nil {
		return "", err
	}
	return language, nil
}

// DetectSentiment is the resolver for the detectSentiment field.
func (r *mutationResolver) DetectSentiment(ctx context.Context, input string) (string, error) {
	sentiment, err := sdk.DetectSentiment(input)
	if err != nil {
		return "", err
	}
	return sentiment, nil
}

// TranslateText is the resolver for the translateText field.
func (r *mutationResolver) TranslateText(ctx context.Context, input *model.TranslateText) (string, error) {
	translatedText, err := sdk.TranslateText(input.Text, input.SourceLanguage, input.TargetLanguage)
	if err != nil {
		return "", err
	}
	return translatedText, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	rows, err := db.MySQL.Query("SELECT * FROM user limit 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

// FetchLastData is the resolver for the fetchLastData field.
func (r *queryResolver) FetchLastData(ctx context.Context) (string, error) {
	return sdk.GetLastData()
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
