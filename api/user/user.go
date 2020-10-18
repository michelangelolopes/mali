package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/michelangelolopes/mali/api/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

//Create Struct
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

func GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var collection = helper.ConnectDB("users")
	var users []User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		defer cancel()
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("user: %v", user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		defer cancel()
		return
	}
	defer cancel()
	json.NewEncoder(response).Encode(users)
}
