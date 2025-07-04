package config

import (
	"os"
	"testing"
)
var DSN string = os.Getenv("DATABASE_URL")

func TestConnectDB_MissingEnv(t *testing.T) {
	os.Unsetenv("DATABASE_URL")

	db, err := ConnectDB()
	if err == nil {
		t.Fatal("expected error when DATABASE_URL is not set, got nil")
	}
	if db != nil {
		t.Fatal("expected db to be nil on error")
	}
}

func TestConnectDB_InvalidDSN(t *testing.T) {
    os.Setenv("DATABASE_URL", "invalid_dsn")

    db, err := ConnectDB()
    if err == nil {
        t.Fatal("expected error when DSN is invalid, got nil")
    }
    if db != nil {
        t.Fatal("expected db to be nil on error")
    }
}

func TestConnectDB_Success(t *testing.T) {
	os.Setenv("DATABASE_URL", DSN)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("Skipping integration test because DATABASE_URL is not set")
	}

	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if db == nil {
		t.Fatal("expected non-nil db")
	}

	db.Close()
}