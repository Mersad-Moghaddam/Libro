package user_test

import (
	"libro/apiSchema/userSchema"
	"testing"
)

func TestChangePasswordRequest(t *testing.T) {
	req := userSchema.ChangePasswordRequest{CurrentPassword: "old", NewPassword: "newpassword"}
	if len(req.NewPassword) < 6 {
		t.Fatal("password too short")
	}
}
