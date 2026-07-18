package ex02_validation

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name      string
		user      User
		wantErr   bool
		wantField string
		wantMsg   string
	}{
		{"valid", User{Name: "ada", Age: 36}, false, "", ""},
		{"empty name", User{Name: "", Age: 20}, true, "name", "must not be empty"},
		{"negative age", User{Name: "ada", Age: -1}, true, "age", "must not be negative"},
		{"too old", User{Name: "ada", Age: 200}, true, "age", "is implausibly large"},
		// Name is checked before age, so an empty name wins even with a bad age.
		{"name beats age", User{Name: "", Age: -5}, true, "name", "must not be empty"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.user)
			if tc.wantErr == (err == nil) {
				t.Fatalf("Validate(%+v) err = %v, wantErr = %t", tc.user, err, tc.wantErr)
			}
			if !tc.wantErr {
				return
			}
			var fe *FieldError
			if !errors.As(err, &fe) {
				t.Fatalf("error is not a *FieldError: %v", err)
			}
			if fe.Field != tc.wantField || fe.Msg != tc.wantMsg {
				t.Errorf("got field=%q msg=%q, want field=%q msg=%q",
					fe.Field, fe.Msg, tc.wantField, tc.wantMsg)
			}
		})
	}
}

func TestFieldErrorFormat(t *testing.T) {
	fe := &FieldError{Field: "name", Msg: "must not be empty"}
	want := `field "name": must not be empty`
	if got := fe.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}
