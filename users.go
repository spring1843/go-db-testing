package main

import "database/sql"

type (
	user struct {
		userID    int
		firstName string
		lastName  string
		email     string
	}

	usersDB struct {
		addStmt *sql.Stmt
	}
)

func newUserDB() *usersDB {
	return &usersDB{
		addStmt: stmts[addUser],
	}
}

func (postDB *usersDB) addUser(user *user) (*user, error) {
	res, err := postDB.addStmt.Exec(user.firstName, user.lastName, user.email)
	if err != nil {
		return nil, err
	}
	userID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.userID = int(userID)
	return user, nil
}
