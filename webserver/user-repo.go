package webserver

import (
"context"
	"../src"
	"cloud.google.com/go/firestore"
)

type UserRepository interface {
	Save()
}

type test_repo struct struct{}

//New Post Repository
func NewUserRepository() UserRepository{
	return &repo{}
}

const (
	projectID string = "xunami-userbase"
	colletionName string = "user"
)

func Save(post *entity.Post) (*entity.Post,error){
	ctx:=context.Background()//gives an empty context
	client, err:= firestore.NewClient(ctx,projectID) //calling function and handling errors  
	if err != nil{
		log.Fatalf("Failed to create Firestore Cllient: %v", err)
		return nil, err
	}
	client.Collection((colletionName)).Add(ctx, map[string]interface{}){
		"Email":post.Email,
		"LastName":post.LastName,
		"FirstName": post.FirstName,
		"Number": post.Number

	}
}


