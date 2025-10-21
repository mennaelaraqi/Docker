package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	ID      int
	Name    string // must be Name not name :D, won't work
	Age     int
	Address Address
}

var (
	nextID = countExistingFiles()
)

// create a file for each added user
// when a user is requested by id you should look for it's file to return it

func countExistingFiles() int {
	files, err := ioutil.ReadDir("users_saved")
	if err != nil {
		return 1
	}
	return len(files) + 1
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	/*
		TODO
		1) create a file named as the "u.ID".txt and save into users_saved directory
		2) marshal the user's json into it
	*/
	fileName := fmt.Sprintf("users_saved/%d.txt", u.ID)

	userJSON, _ := json.Marshal(u)

	os.WriteFile(fileName, userJSON, 0644)

	return u, nil
}

func GetUserByID(id int) (User, error) {
	/*
		TODO
		1) look for the file named "u.ID".txt in users_saved directory
		2) unmarshal the read json into a user and return that user with nil error
	*/

	fileName := fmt.Sprintf("users_saved/%d.txt", id)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return User{}, fmt.Errorf("User with ID '%v' not found", id)
	}

	var u User
	json.Unmarshal(data, &u)
	return u, nil
}
