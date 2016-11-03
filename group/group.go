package group

import (
	gq "github.com/graphql-go/graphql"

	"github.com/Jimskapt/test-graphql/user"
)

type Group struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Users []user.User `json:"users"`
}

var GroupType = gq.NewObject(gq.ObjectConfig{
	Name: "Group",
	Fields: gq.Fields{
		"id": &gq.Field{
			Type: gq.String,
		},
		"name": &gq.Field{
			Type: gq.String,
		},
		"users": &gq.Field{
			Type: gq.NewList(user.UserType),
		},
	},
})
