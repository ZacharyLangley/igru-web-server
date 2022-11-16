package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	authenticationv1 "github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1"
	"github.com/ZacharyLangley/igru-web-server/pkg/proto/authentication/v1/authenticationv1connect"
	"github.com/bufbuild/connect-go"
)

const maxUsers = 5
const port = 80

func main() {
	userClient := authenticationv1connect.NewUserServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://localhost:%d/", port),
	)
	groupClient := authenticationv1connect.NewGroupServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://localhost:%d/", port),
	)
	// Create users
	userIDs := make([]string, maxUsers)
	for i := 0; i < maxUsers; i++ {
		req := connect.NewRequest(&authenticationv1.CreateUserRequest{
			FullName: fmt.Sprintf("Eric-%d", i),
			Email:    fmt.Sprintf("bobcob-%d@hotmail.com", i),
			Password: fmt.Sprintf("password%d%d%d", i, i, i),
		})
		ctx, done := context.WithTimeout(context.Background(), time.Second*5)
		defer done()
		res, err := userClient.CreateUser(ctx, req)
		if err != nil {
			log.Fatalln(err)
		}
		userIDs[i] = res.Msg.User.Id
	}
	// Create group
	req := connect.NewRequest(&authenticationv1.CreateGroupRequest{
		Name: "Eric",
	})
	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	defer done()
	res, err := groupClient.CreateGroup(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	groupID := res.Msg.Group.Id
	// Add group members
	for _, userID := range userIDs {
		req := connect.NewRequest(&authenticationv1.AddGroupMemberRequest{
			GroupId: groupID,
			UserId:  userID,
		})
		ctx, done := context.WithTimeout(context.Background(), time.Second*5)
		defer done()
		_, err := groupClient.AddGroupMember(ctx, req)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
