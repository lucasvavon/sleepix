package adapters

import (
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db,
	}
}

// NewPostgresRepository creates a new instance of PostgresRepository
/*func NewPostgresRepository(connectionString string) (*PostgresRepository, error) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
	}

	return &PostgresRepository{
		db: db,
	}, nil
}*/

func (repo *PostgresRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
