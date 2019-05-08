package users

type UserManager struct {

}

func Instance() *UserManager {
	m := &UserManager {
	}

	return m
}
