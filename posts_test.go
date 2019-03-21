package main

import "errors"

type postsMock struct {
	posts map[int]*post
	count int
}

func (postMock *postsMock) addPost(post *post) (*post, error) {
	post.PostID = postMock.count
	postMock.posts[postMock.count] = post
	postMock.count++
	return post, nil
}

func (postMock *postsMock) getPost(postID int) (*post, error) {
	post, ok := postMock.posts[postID]
	if !ok {
		return nil, errors.New("post with given ID does not exist")
	}
	return post, nil
}
