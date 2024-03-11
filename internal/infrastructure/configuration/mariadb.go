package configuration

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/nosilex/crebito/pkg/helper"
)

func newMariaDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, errors.New("missing database host")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, errors.New("missing database user")
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		return nil, errors.New("missing database password")

	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return nil, errors.New("missing database name")

	}

	port := helper.Coalesce(os.Getenv("DB_PORT"), "3306")

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}
