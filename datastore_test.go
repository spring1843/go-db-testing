package main

var _ postsDS = new(postsDB)
var _ postsDS = new(postsMock)
var _ usersDS = new(usersDB)
var _ usersDS = new(usersMock)
