package application

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed init.sql
var initSQL string

type App struct {
	Path     string
	Config   *Config
	Database *sql.DB
}

func NewApp() (*App, error) {

	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return nil, err
	}

	dir := filepath.Join(homeDir, ".config", "clai")

	// Create data directory if it doesn't exist
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	cfg, err := LoadConfig(dir, "config") // Load config from file
	if err != nil {
		fmt.Printf("Error loading config from file: %v\n", err)
		return nil, err
	}

	dbPath := filepath.Join(dir, "history.db")

	// TODO: Maybe replace with postgres in production
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Check if database is initialized by looking for a known table
	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' LIMIT 1").Scan(&tableName)
	if err == sql.ErrNoRows {
		// Database is not initialized, run init SQL
		_, err = db.Exec(initSQL)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to initialize database: %w", err)
		}
	} else if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to check database initialization: %w", err)
	}

	return &App{
		Path:     dir,
		Config:   cfg,
		Database: db,
	}, nil

}
