package postgres

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"login-user/config"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgresConnection(c *config.Config) (*Postgres, error) {
	db, err := gorm.Open(sqlite.Open(c.DB.Name), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
