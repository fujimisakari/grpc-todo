package domain

type User UserDto

func (u *User) ToDto() *UserDto {
	return (*UserDto)(u)
}

func (u *UserDto) ToEntity() (*User, error) {
	return (*User)(u), nil
}
