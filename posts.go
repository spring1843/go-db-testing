package main

import "database/sql"

type (
	post struct {
		PostID int
		UserID int
		Title  string
		Body   string
	}
	postsDB struct {
		addStmt *sql.Stmt
		getStmt *sql.Stmt
	}
)

func newPostsDB() *postsDB {
	return &postsDB{
		addStmt: stmts[addPost],
		getStmt: stmts[getPost],
	}
}

func (postDB *postsDB) addPost(post *post) (*post, error) {
	res, err := postDB.addStmt.Exec(post.UserID, post.Title, post.Body)
	if err != nil {
		return nil, err
	}
	postID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	post.PostID = int(postID)
	return post, nil
}

func (postDB *postsDB) getPost(postID int) (*post, error) {
	res := postDB.getStmt.QueryRow(postID)
	post := new(post)
	if err := res.Scan(&post.PostID, &post.UserID, &post.Title, &post.Body); err != nil {
		return nil, err
	}
	return post, nil
}
