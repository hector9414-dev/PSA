package main

const token = "tokenmuyseguro123456"

var users []*User

type User struct {
	Nick     string `json: "nick`
	Password string `json: "password`
}

func addUser(u *User) {
	users = append(users, u)
}

func login(u *User) bool {
	for _, v := range users {
		if v.Nick == u.Nick && v.Password == u.Password {
			return true
		}
	}
	return false
}

type MessageResponse struct {
	Type string `json:"type"`
	From string `json:"from"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type WebSocketResponse struct {
	Type string `json:"type"`
	From string `json:"from"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}