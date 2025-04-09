package main

import (
	"testing"
	"time"
)

func TestSessionManager_StartSession(t *testing.T) {
	sessionManager := NewSessionManager()
	userID := "testuser"

	sessionID := sessionManager.StartSession(userID)

	// checking if session was created
	if sessionID == "" {
		t.Error("StartSession did not return a valid session ID.")
	}

	// checking if session exists in the manager
	session, exists := sessionManager.sessions[sessionID]
	if !exists {
		t.Error("Session not found in the session manager.")
	}

	// checking if session properties are correct
	if session.UserID != userID {
		t.Errorf("Expected UserID: %s, Got: %s", userID, session.UserID)
	}

	// checking if session expiration time is within the expected range
	expectedExpiration := time.Now().Add(120 * time.Second)
	if session.ExpiresAt.Before(expectedExpiration.Add(-time.Second)) || session.ExpiresAt.After(expectedExpiration.Add(time.Second)) {
		t.Errorf("Expected expiration time: %v, Got: %v", expectedExpiration, session.ExpiresAt)
	}
}

func TestSessionManager_GetSession(t *testing.T) {

}
