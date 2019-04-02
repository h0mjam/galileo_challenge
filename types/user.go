package types

type User struct {
	ID      uint64      `json:"id"`
	Token   string      `json:"token"`
	Email   string      `json:"email"`
	Devices DeviceStore `json:"-"`
}

func NewUser() *User {
	return &User{
		Token: NewToken(),
	}
}

func (u *User) AppendDevice(d *Device) {
	u.Devices.Add(d)
}

func (u *User) FindDevice(k uint64) (*Device, error) {
	return u.Devices.FindByID(k)
}

type UserDevices struct {
	Devices []*Device `json:"devices"`
}

type UserCreateRequest struct {
	Email string `json:"email"`
}

type UserStore interface {
	Add(user *User)
	FindByToken(token string) (*User, error)
	FindByEmail(email string) (*User, error)
}
