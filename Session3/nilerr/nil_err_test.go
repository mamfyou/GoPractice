package nilerr

import (
	"errors"
	"testing"
)

/*
implement a function that checks if dynamic value of an error is nil:
- define the CustomError type; *CustomError must implement error.
- your function must return true if dynamic value of CustomError errors is nil.
*/

func TestIsErrNil(t *testing.T) {

	/* --------------- */

	var err error

	if !IsCustomErrNil(err) {
		t.Errorf("expected IsCustomErrNil to return true")
	}

	err = errors.New("an error")

	if IsCustomErrNil(err) {
		t.Errorf("expected IsCustomErrNil to return false")
	}

	/* --------------- */

	var cuerr *CustomError

	if !IsCustomErrNil(cuerr) {
		t.Errorf("expected IsCustomErrNil to return true")
	}

	err = cuerr

	if !IsCustomErrNil(err) {
		t.Errorf("expected IsCustomErrNil to return true")
	}

	/* --------------- */

	cuerr = &CustomError{
		Message: "error message",
	}

	if IsCustomErrNil(cuerr) {
		t.Errorf("expected IsCustomErrNil to return false")
	}

	err = cuerr

	if IsCustomErrNil(err) {
		t.Errorf("expected IsCustomErrNil to return false")
	}

	/* --------------- */

}
