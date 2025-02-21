package user

import (
	"errors"
	"testing"
)

var validationError *ValidationError

func TestInsertCorrectUser(t *testing.T) {
	var user_id interface{}

	user := &User{
		FullName: "Ali Mohammadi2",
		UserName: "ali_new",
		Role:     Manager,
		Address:  "test address",
		IsActive: true,
	}

	user_id, err := Insert(user)
	if err != nil {
		t.Errorf("expected 'id' field but unexpected error happened: %s", err.Error())
	}

	if _, ok := user_id.(int64); !ok {
		t.Error("expected id to be a number")
	}
}


func TestInsertIncorrectUser(t *testing.T) {
	empty_name_user := &User{
		FullName: "",
		UserName: "hosna",
		Role:     Manager,
		Address:  "test address", 
		IsActive: true,
	}

	_, err := Insert(empty_name_user)
	if err == nil {
		t.Error("expected error for empty full_name got no error")
	} else if !errors.As(err, &validationError) {
		t.Errorf("expected error type 'ValidationError' got: %s", err.Error())
	} else if err.Error() != "user field 'fullname' can not be empty string" {
		t.Errorf("expected: %s, got: %s", "user field 'fullname' can not be empty string", err.Error())
	}



	empty_username_user := &User{
		FullName: "test",
		UserName: "",
		Role:     Standard,
		Address:  "test address", 
		IsActive: true,
	}

	_, err = Insert(empty_username_user)
	if err == nil {
		t.Error("expected error for empty fullName got no error")
	} else if !errors.As(err, &validationError) {
		t.Errorf("expected error type 'ValidationError' got: %s", err.Error())
	} else if err.Error() != "user field 'username' can not be empty string" {
		t.Errorf("expected: %s, got: %s", "user field 'username' can not be empty string", err.Error())
	}

	empty_address_user := &User{
		FullName: "test",
		UserName: "noAddressUser",
		Role:     Standard,
		Address:  "", 
		IsActive: true,
	}

	_, err = Insert(empty_address_user)
	if err == nil {
		t.Error("expected error for empty fullName got no error")
	} else if !errors.As(err, &validationError) {
		t.Errorf("expected error type 'ValidationError' got: %s", err.Error())
	} else if err.Error() != "user field 'address' can not be empty string" {
		t.Errorf("expected: %s, got: %s", "user field 'address' can not be empty string", err.Error())
	}


	wrong_role_user := &User{
		FullName: "test",
		UserName: "test",
		Role:     4,
		Address:  "test address", 
		IsActive: true,
	}

	_, err = Insert(wrong_role_user)
	if err == nil {
		t.Error("expected error for empty fullName got no error")
	} else if !errors.As(err, &validationError) {
		t.Errorf("expected error type 'ValidationError' got: %s", err.Error())
	} else if err.Error() != "invalid value for field 'role'" {
		t.Errorf("expected: %s, got: %s", "invalid value for field 'role'", err.Error())
	}


	repeated_username_user := &User{
		FullName: "repeated",
		UserName: "repeated",
		Role:     Standard,
		Address:  "test address", 
		IsActive: true,
	}

	Insert(repeated_username_user)
	_, err = Insert(repeated_username_user)
	if err == nil {
		t.Error("expected error for empty fullName got no error")
	} else if !errors.As(err, &validationError) {
		t.Errorf("expected error type 'ValidationError' got: %s", err.Error())
	} else if err.Error() != "user with this username already exists" {
		t.Errorf("expected: %s, got: %s", "user with this username already exists", err.Error())
	}
}


func TestGetUser(t *testing.T) {
	user := &User{
		FullName: "correct",
		UserName: "correct",
		Role:     Manager,
		Address:  "test address", 
		IsActive: true,
	}

	user_id, err := Insert(user)
	if err != nil {
		t.Errorf("got error while creating user for test: %s", err.Error())
	}

	_, err = Get(-1)
	if err == nil {
		t.Error("expected error for invalid user id got no error")
	} else if !errors.As(err, &validationError) {
		t.Error("expected error type 'ValidationError' got unexpected error type")
	} else if err.Error() != "user with provided id does not exists" {
		t.Errorf("expected: %s, got: %s", "user with provided id does not exists", err.Error())
	}

	_, err = Get(user_id)
	if err != nil {
		t.Errorf("expected user detail got error: %s", err.Error())
	}

}

func TestUpdate(t *testing.T) {
	user := &User{
		FullName: "testUpdateUser2",
		UserName: "testUpdateUser2",
		Role:     Manager,
		Address:  "test address", 
		IsActive: true,
	}

	
	user_id, err := Insert(user)
	updated_user := &User{
		ID: user_id,
		FullName: "testUpdate2",
		UserName: "testUpdate2",
		Role:     Manager,
		Address:  "test address", 
		IsActive: true,
	}

	if err != nil {
		t.Errorf("got error while creating user for test: %s", err.Error())
	}

	err = Update(updated_user)
	if err != nil {
		t.Errorf("expected no error got: %s", err.Error())
	} else if user, _ = Get(user_id);(user.FullName != updated_user.FullName || user.UserName != updated_user.UserName) {
		t.Error("user is not updated")
	}

}

func TestDelete(t *testing.T) {
	user := &User{
		FullName: "testDeleteUser",
		UserName: "testDeleteUser",
		Role:     Manager,
		Address:  "test address", 
		IsActive: true,
	}

	
	user_id, err := Insert(user)

	err = Delete(user_id)
	if err != nil {
		t.Errorf("expected no error got: %s", err.Error())
	}

	err = Delete(user_id)
	if err == nil {
		t.Error("expected error for invalid id got no error")
	} else if !errors.As(err, &validationError) {
		t.Error("expected error type 'ValidationError' got unexpected error")
	}
}