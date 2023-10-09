package migration

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	migrator.AddMigration(&Migration{
		ID:        0,
		Name:      "20231009162119_users",
		Batch:     0,
		CreatedAt: time.Time{},
		Up:        mig_20231009162119_users_up,
		Down:      mig_20231009162119_users_down,
	})
}

func mig_20231009162119_users_up(tx *gorm.DB) error {
	err := tx.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL NOT NULL,
		name VARCHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		role_id INT NOT NULL,
		password TEXT NOT NULL,
		sallary numeric(20, 2) NULL,
		created_at timestamptz NOT NULL,
		created_by VARCHAR(255) NOT NULL,
		updated_at timestamptz NOT NULL,
		updated_by VARCHAR(255) NOT NULL,
		deleted_at timestamptz NULL,
		deleted_by VARCHAR(255) NULL,
		CONSTRAINT users_pkey PRIMARY KEY (id),
		CONSTRAINT fk_role_id FOREIGN KEY(role_id) REFERENCES roles(id)
	)`).Error

	return err
}

func mig_20231009162119_users_down(tx *gorm.DB) error {
	err := tx.Exec(`DROP TABLE IF EXISTS users`).Error
	return err
}
