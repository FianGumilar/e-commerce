package utils

import "gorm.io/gorm"

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	// return closure
	return func(db *gorm.DB) *gorm.DB {
		// Mengatur halaman
		page := offset

		if page <= 0 {
			page = 1
		}

		// Mengatur ukuran halaman
		pageSize := offset

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 10:
			pageSize = 10
		}

		// Menghitung offset & batasan
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)
	}
}
