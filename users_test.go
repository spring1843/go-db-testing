package main

type usersMock struct {
	users map[int]*user
	count int
}

func (postMock *usersMock) addUser(user *user) (*user, error) {
	user.userID = postMock.count
	postMock.users[postMock.count] = user
	postMock.count++
	return user, nil
}
