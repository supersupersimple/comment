package api

import "testing"

// unittest
func TestCheckPasswordHash(t *testing.T) {
	hash, err := HashPassword("mypassword")
	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
		return
	}
	if !CheckPasswordHash("mypassword", hash) {
		t.Errorf("CheckPasswordHash failed")
	}
	if CheckPasswordHash("wrongpassword", hash) {
		t.Errorf("CheckPasswordHash should return false for wrong password")
	}
}
