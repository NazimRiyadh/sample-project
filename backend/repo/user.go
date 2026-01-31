package repo

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsOwner   bool   `json:"is_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (r *userRepo) Create(u User) (*User, error) {
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return &u, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}
	return nil, nil
}
