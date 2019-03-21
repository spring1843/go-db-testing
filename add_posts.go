package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const maxPosts = 10

var (
	postDB     postsDS
	addedPosts int
	userID     int
)

func addPosts() {
	userID = insertUser(newUserDB())

	postDB = newPostsDB()
	doEvery(2*time.Second, insertPost)
}

func insertUser(userDS usersDS) int {
	user, err := userDS.addUser(&user{
		firstName: "foo",
		lastName:  "bar",
		email:     fmt.Sprintf("foo@a%d.com", rand.New(rand.NewSource(time.Now().UnixNano())).Int()),
	})
	if err != nil {
		log.Fatalf("Failed adding new user. Error :%s", err)
	}
	return user.userID
}

func doEvery(d time.Duration, f func(postsDS)) {
	for _ = range time.Tick(d) {
		if addedPosts != maxPosts {
			f(postDB)
		} else {
			os.Exit(0)
		}
	}
}

func insertPost(postDB postsDS) {
	addedPosts++
	_, err := postDB.addPost(&post{
		UserID: userID,
		Title:  fmt.Sprintf("post #%d", addedPosts),
		Body:   fmt.Sprintf("Body of post #%d", addedPosts),
	})
	if err != nil {
		log.Fatalf("Failed adding post %d. Error :%s", addedPosts, err)
	}
	fmt.Printf("Added post #%d\n", addedPosts)
}
