package mysql

import (
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQL it is a MySQL module context structure
type MySQL struct {
	client *gorm.DB
}

// Settings contains settings for MySQL
type Settings struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Connect connects to MySQL
func Connect(s Settings) (MySQL, error) {

	client, err := gorm.Open(gmysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			s.User,
			s.Password,
			s.Host,
			s.Port,
			s.Database)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		return MySQL{}, err
	}

	return MySQL{
		client: client,
	}, nil
}

// Close closes MySQL connection
func (m *MySQL) Close() error {
	db, _ := m.client.DB()
	return db.Close()
}
