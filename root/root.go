package root

import (
	"errors"
	"strconv"

	gq "github.com/graphql-go/graphql"

	"github.com/Jimskapt/test-graphql/group"
	"github.com/Jimskapt/test-graphql/user"
)

var GroupList []group.Group
var UserList []user.User
var IdCount int

var RootQuery = gq.NewObject(gq.ObjectConfig{
	Name: "RootQuery",
	Fields: gq.Fields{
		"user": &gq.Field{
			Type: user.UserType,
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type: gq.String,
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				requestedID, castingOK := params.Args["id"].(string)
				if !castingOK {
					return user.User{}, errors.New("user.User : error while casting ID.")
				}

				for _, user := range UserList {
					if user.ID == requestedID {
						return user, nil
					}
				}

				return user.User{}, errors.New("user.User : ID not found")
			},
		},
		"users": &gq.Field{
			Type: gq.NewList(user.UserType),
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return UserList, nil
			},
		},
		"group": &gq.Field{
			Type: group.GroupType,
			Args: gq.FieldConfigArgument{
				"id": &gq.ArgumentConfig{
					Type: gq.String,
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				requestedID, castingOK := params.Args["id"].(string)
				if !castingOK {
					return group.Group{}, errors.New("group.Group : error while casting ID.")
				}

				for _, group := range GroupList {
					if group.ID == requestedID {
						return group, nil
					}
				}

				return group.Group{}, errors.New("group.Group : ID not found")
			},
		},
		"groups": &gq.Field{
			Type: gq.NewList(group.GroupType),
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				return GroupList, nil
			},
		},
	},
})

var RootMutation = gq.NewObject(gq.ObjectConfig{
	Name: "RootMutation",
	Fields: gq.Fields{
		"newUser": &gq.Field{
			Type: user.UserType,
			Args: gq.FieldConfigArgument{
				"username": &gq.ArgumentConfig{
					Type: gq.NewNonNull(gq.String),
				},
			},
			Resolve: func(params gq.ResolveParams) (interface{}, error) {
				name, castingOK := params.Args["username"].(string)
				if !castingOK {
					return user.User{}, errors.New("Error while casting username.")
				}

				IdCount++
				user := user.User{ID: strconv.Itoa(IdCount), Username: name}

				UserList = append(UserList, user)

				return user, nil
			},
		},
	},
})
