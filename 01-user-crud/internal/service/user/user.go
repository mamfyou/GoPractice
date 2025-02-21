package user

import (
	"errors"
	"mamfyou/internal/database"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) error {
	return &ValidationError{Message: message}
}

var ErrorUnexpectedDB error = errors.New("unexpected database errors")

func Insert(user *User) (id int64, err error) {
	err = validate(user, "POST")
	if err != nil {
		return 0, err
	}

	db := database.GetDB()
	err = db.Create(user).Error

	if err != nil {
		return -1, ErrorUnexpectedDB
	}

	return user.ID, nil
}

func Get(id int64) (user *User, err error) {
	user = &User{ID: id}

	db := database.GetDB()
	var count int64
	if err := db.Model(&User{}).Where("id = ?", user.ID).Count(&count).Error; err == nil && !(count > 0) {
		return user, NewValidationError("user with provided id does not exists")
	}

	if err := db.Model(&User{}).Where("id = ?", user.ID).First(user).Error; err != nil {
		return nil, err
	}
	
	return user, nil
}

func Update(user *User) (err error) {
	err = validate(user, "PUT")
	if err != nil {
		return err
	}

	user_old, err := Get(user.ID)
	if err != nil {
		return err
	}

	db := database.GetDB()

	err = db.Model(&user_old).Updates(User{FullName: user.FullName, UserName: user.UserName, IsActive: user.IsActive, Address: user.Address, Role: user.Role}).Error
	if err != nil {
		return ErrorUnexpectedDB
	}

	return nil
}

func Delete(id int64) (err error) {
	validate(&User{ID: id}, "DELETE")

	user, err := Get(id)

	if err != nil {
		return err
	}

	db := database.GetDB()

	err = db.Delete(&user).Error
	if err != nil {
		return ErrorUnexpectedDB
	}

	return nil
}

// ----- validators -----

func validate(user *User, method string) (err error) {
	if method != "POST" {
		if user == nil {
			return NewValidationError("user can not be nil")
		}

		if user.ID < 1 && user.ID != 0 {
			return NewValidationError("user field 'id' can not be less than 1")
		}

		if _, err := Get(user.ID); err != nil {
			return NewValidationError("user not exists")
		}
	}

	if method != "GET" && method != "DELETE" {
		db := database.GetDB()
		var count int64
		if err := db.Model(&User{}).Where("user_name = ?", user.UserName).Count(&count).Error; err == nil && count > 0 {
			return NewValidationError("user with this username already exists")
		}
	}

	if user.FullName == "" {
		return NewValidationError("user field 'fullname' can not be empty string")
	}

	if user.UserName == "" {
		return NewValidationError("user field 'username' can not be empty string")
	}

	if user.Address == "" {
		return NewValidationError("user field 'address' can not be empty string")
	}

	if user.Role <= startRoleType || user.Role >= endRoleType {
		return NewValidationError("invalid value for field 'role'")
	}

	return nil
}
