package initializer

import models "github.com/AlmedinNasufi/go-book-app/Models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})
}
