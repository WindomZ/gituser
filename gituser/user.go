package gituser

import (
	"errors"
	"regexp"
)

type _GitUser struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (g _GitUser) Valid() error {
	if len(g.User) == 0 {
		return errors.New("Invalid user: " + g.User)
	}
	if len(g.Name) == 0 {
		return errors.New("Invalid name: " + g.Name)
	}
	if !regexp.MustCompile(`^[a-z_0-9.-]{1,64}@([a-z0-9-]{1,200}.){1,5}[a-z]{1,6}$`).MatchString(g.Email) {
		return errors.New("Invalid email: " + g.Email)
	}
	return nil

}

type _GitUsers struct {
	Users []*_GitUser `json:"users"`
}

func (g *_GitUsers) Add(u *_GitUser) *_GitUsers {
	if u != nil && u.Valid() == nil {
		for i, v := range g.Users {
			if v.User == u.User {
				g.Users[i] = u
				return g
			}
		}
		g.Users = append(g.Users, u)
	}
	return g
}

func (g *_GitUsers) Remove(u string) bool {
	if len(u) != 0 {
		for i, v := range g.Users {
			if v.User == u {
				if i+1 >= len(g.Users) {
					g.Users = g.Users[:i]
				} else {
					g.Users = append(g.Users[:i], g.Users[i+1:]...)
				}
				return true
			}
		}
	}
	return false
}

func (g _GitUsers) Get(i int) *_GitUser {
	if i < 0 || len(g.Users) <= i {
		return nil
	}
	return g.Users[i]
}

func (g _GitUsers) Size() int {
	return len(g.Users)
}

func (g _GitUsers) Find(user string) (*_GitUser, bool) {
	if len(user) != 0 {
		for _, v := range g.Users {
			if v.User == user {
				return v, true
			}
		}
	}
	return nil, false
}
