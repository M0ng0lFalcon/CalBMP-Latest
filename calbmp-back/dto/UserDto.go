package dto

import "calbmp-back/model"

type UserDto struct {
	Username string `json:"name"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
	}
}
