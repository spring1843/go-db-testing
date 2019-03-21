package main

type (
	postsDS interface {
		addPost(post *post) (*post, error)
		getPost(postID int) (*post, error)
	}
	usersDS interface {
		addUser(user *user) (*user, error)
	}
)
