package user

import (
	gq "github.com/graphql-go/graphql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var UserType = gq.NewObject(gq.ObjectConfig{
	Name: "User",
	Fields: gq.Fields{
		"id": &gq.Field{
			Type: gq.String,
		},
		"username": &gq.Field{
			Type: gq.String,
		},
	},
})
