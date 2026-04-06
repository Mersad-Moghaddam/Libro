package auth_test

import (
	"libro/apiSchema/authSchema"
	"testing"
)

func TestRegisterRequestShape(t *testing.T) {
	req := authSchema.RegisterRequest{Name: "A", Email: "a@b.com", Password: "secret"}
	if req.Email == "" || req.Password == "" || req.Name == "" {
		t.Fatal("invalid register request")
	}
}
