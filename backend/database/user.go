package database

var users []Users

type Users struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsOwner   bool   `json:"is_owner"`
}

func StoreUser(u Users) Users {
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func GetUser(id int) *Users {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
