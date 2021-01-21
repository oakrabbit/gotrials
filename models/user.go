package models

type User struct {
	Id       string
	Name     string
	Fullname string
}

type Address struct {
	Street   string
	Door     string
	Location string
	ZipCode  string
}

type UserBuilder interface {
	SetName(string) UserBuilder
	SetFullname(string) UserBuilder
	Build() User
}

func (c *User) SetName(name string) UserBuilder {
	c.Name = name
	return c
}
func (c *User) SetFullname(fname string) UserBuilder {
	c.Fullname = fname
	return c
}

func (c *User) Build() User {
	return User{
		Name:     c.Name,
		Fullname: c.Fullname,
	}
}
