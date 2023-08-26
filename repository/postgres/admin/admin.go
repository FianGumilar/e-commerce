package admin

import (
	"context"

	"github.com/FianGumilar/e-commerce/interfaces"
	"github.com/FianGumilar/e-commerce/models/entity"
	"github.com/FianGumilar/e-commerce/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewAdminRepository(con *gorm.DB) interfaces.AdminRepository {
	return &repository{db: con}
}

// FindAllAdmin implements interfaces.AdminRepository.
func (r repository) FindAllAdmin(ctx context.Context, offset int, limit int) (admins []entity.Admin, err error) {
	err = r.db.Debug().Scopes(utils.Paginate(offset, limit)).Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

// FindOneAdminByID implements interfaces.AdminRepository.
func (r repository) FindOneAdminByID(ctx context.Context, id int) (admin *entity.Admin, err error) {
	err = r.db.Debug().Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// FindOneAdminByEmail implements interfaces.AdminRepository.
func (r repository) FindOneAdminByEmail(ctx context.Context, email string) (admin *entity.Admin, err error) {
	err = r.db.Debug().Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil

}

// InsertAdmin implements interfaces.AdminRepository.
func (r repository) InsertAdmin(ctx context.Context, admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Debug().Create(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// UpdateAdmin implements interfaces.AdminRepository.
func (r repository) UpdateAdmin(ctx context.Context, admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Debug().Save(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// DeleteAdmin implements interfaces.AdminRepository.
func (r repository) DeleteAdmin(ctx context.Context, admin entity.Admin) error {
	err := r.db.Debug().Delete(&admin).Error
	if err != nil {
		return err
	}
	return nil
}

// CountAdmin implements interfaces.AdminRepository.
func (r repository) CountAdmin(ctx context.Context) (count int64, err error) {
	err = r.db.Debug().Model(&entity.Admin{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
