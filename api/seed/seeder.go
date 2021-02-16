package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/bimosyah/go-test/api/models"
)

var users = []models.User{
	models.User{
		Username: "Steven victor",
		Password: "password",
	},
	models.User{
		Username: "Martin Luther",
		Password: "password1",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
