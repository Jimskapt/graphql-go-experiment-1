package main

import (
	"encoding/json"
	"fmt"

	gq "github.com/graphql-go/graphql"

	"github.com/Jimskapt/test-graphql/group"
	"github.com/Jimskapt/test-graphql/root"
	"github.com/Jimskapt/test-graphql/user"
)

var userSchema gq.Schema

func main() {
	root.UserList = []user.User{}

	rootuser := user.User{ID: "1", Username: "root"}
	jimskapt := user.User{ID: "2", Username: "Jimskapt"}
	gopher := user.User{ID: "3", Username: "Gopher"}
	nobody := user.User{ID: "4", Username: "nobody"}

	root.UserList = append(root.UserList, rootuser, jimskapt, gopher, nobody)

	// *******************************************************

	root.GroupList = []group.Group{}

	powerusers := group.Group{ID: "5", Name: "PowerUsers", Users: []user.User{rootuser, jimskapt}}
	normalusers := group.Group{ID: "6", Name: "NomalUsers", Users: []user.User{gopher, nobody}}

	root.GroupList = append(root.GroupList, powerusers, normalusers)

	// *******************************************************

	root.IdCount = 6

	// *******************************************************

	schema, err := gq.NewSchema(gq.SchemaConfig{
		Query:    root.RootQuery,
		Mutation: root.RootMutation,
	})

	if err != nil {
		fmt.Println(err)
	} else {

		userSchema = schema

		exec(`mutation {
				newUser(username:"newbie") {
					id
				}
			}`)

		exec(`mutation {
				newUser(username:"Rami Malek") {
					id
				}
			}`)

		exec(`{
				users {
					username
				}
			}`)

		exec(`{
				group(id:"5") {
					name
				}
			}`)

		exec(`{
				groups {
					name
				}
			}`)

		exec(`{
				groups {
					id,
					name,
					users {
						id,
						username
					}
				}
			}`)

	}

}

func exec(query string) {

	fmt.Println("*****************************************************************")

	fmt.Println("NEW QUERY : ")

	fmt.Println(query)

	fmt.Println()

	fmt.Println("RESULT : ")

	result := gq.Do(gq.Params{
		Schema:        userSchema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Errors : %v", result.Errors)
	} else {
		data, _ := json.MarshalIndent(result, "", "\t")
		fmt.Println(string(data))
	}

	fmt.Println()
}
