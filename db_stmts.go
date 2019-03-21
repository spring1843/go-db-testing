package main

import "database/sql"

const (
	addPost = iota
	getPost
	addUser
)

var (
	queries = map[int]string{
		addPost: "INSERT INTO posts (user_id, title, body) VALUES (?, ?, ?)",
		getPost: "SELECT post_id, user_id, title, body FROM posts WHERE post_id = ?",
		addUser: "INSERT INTO users (first_name, last_name, email) VALUES (?,?,?)",
	}
	stmts map[int]*sql.Stmt
)

func prepareStmts(db *sql.DB) error {
	stmts = make(map[int]*sql.Stmt, len(queries))
	for i, query := range queries {
		stmt, err := db.Prepare(query)
		if err != nil {
			return err
		}
		stmts[i] = stmt
	}
	return nil
}

func closeStmts() error {
	for _, stmt := range stmts {
		if err := stmt.Close(); err != nil {
			return err
		}
	}
	return nil
}
