package migration

import (
	_ "embed"
	"github.com/TUM-Dev/Campus-Backend/server/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func (m TumDBMigrator) migrate20221119131300() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20221119131300",
		Migrate: func(tx *gorm.DB) error {

			if err := tx.AutoMigrate(
				&model.IOSDevice{},
				&model.Crontab{},
				&model.IOSDeviceRequestLog{},
				&model.EncryptedGrade{},
			); err != nil {
				return err
			}

			err := SafeEnumMigrate(tx, &model.Crontab{}, "type", "iosNotifications", "iosActivityReset")

			if err != nil {
				return err
			}

			return nil
		},

		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable(&model.IOSDevice{}); err != nil {
				return err
			}
			if err := tx.Migrator().DropTable(&model.IOSDeviceRequestLog{}); err != nil {
				return err
			}
			if err := tx.Migrator().DropTable(&model.EncryptedGrade{}); err != nil {
				return err
			}

			return SafeEnumRollback(tx, &model.Crontab{}, "type", "iosNotifications", "iosActivityReset")
		},
	}
}
