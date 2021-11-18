package _facade

// User 实体
type User struct {}

// UserService 实现
type UserService struct{}

// IUserService 接口
type IUserService interface {
	Exists(name string) (bool, error)
	Login(phone int, nickname string) (*User, error)
	Register(phone int, nickname string) (*User, error)
}

// IUserFacade 门面
type IUserFacade interface {
	LoginOrRegister(phone int, nickname string) (*User, error)
}

func (u *UserService) Exists(nickname string) (bool, error) {
	// ...
	return false, nil
}

func (u *UserService) Login(phone int, nickname string) (*User, error) {
	// 基础校验...
	// 数据库校验...
	return &User{}, nil
}

func (u *UserService) Register(phone int, nickname string) (*User, error) {
	// 基础校验...
	// 数据库校验...
	// 创建用户...
	return &User{}, nil
}

func (u *UserService) LoginOrRegister(phone int, nickname string) (*User, error) {
	ok, err := u.Exists(nickname)
	if err != nil {
		return nil, err
	}
	if ok {
		return u.Register(phone, nickname)
	} else {
		return u.Login(phone, nickname)
	}
}