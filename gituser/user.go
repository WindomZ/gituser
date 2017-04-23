package gituser

type _GitUser struct {
	User  string `json:"user"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (g _GitUser) Valid() bool {
	return len(g.User) != 0 && len(g.Name) != 0 && len(g.Email) != 0
}

type _GitUsers struct {
	Users []*_GitUser `json:"users"`
}

func (g *_GitUsers) Add(u *_GitUser) *_GitUsers {
	if u != nil && u.Valid() {
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

func (g _GitUsers) Get(i int) *_GitUser {
	if i < 0 || len(g.Users) <= i {
		return nil
	}
	return g.Users[i]
}

func (g _GitUsers) Size() int {
	return len(g.Users)
}
