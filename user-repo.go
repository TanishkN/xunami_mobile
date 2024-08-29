// /package webserver

// import (
// 	"context"
// 	"log"
// 	"cloud.google.com/go/firestore"
// )

// type UserRepository interface {
// 	Save()
// }

// type test_repo struct struct{}

// //New Post Repository
// func NewUserRepository() UserRepository{
// 	return &test_repo{}
// }

// const (
// 	projectID string = "xunami-userbase"
// 	colletionName string = "user"
// )

// func Save(user *entity.User) (*entity.User,error){
// 	ctx:=context.Background()//gives an empty context
// 	client, err:= firestore.NewClient(ctx,projectID) //calling function and handling errors
// 	if err != nil{
// 		log.Fatalf("Failed to create Firestore Cllient: %v", err)
// 		return nil, err
// 	}
// 	client.Collection((colletionName)).Add(ctx, map[string]interface{}){
// 		"Email":user.Email,
// 		"LastName":user.LastName,
// 		"FirstName":user.FirstName,
// 		"Number":user.Number

// 	}
// 	if err !=nil {
// 		log.Fatalf("Failed adding a new User:"%v, err)
// 		return nil, err
// 	}
// }