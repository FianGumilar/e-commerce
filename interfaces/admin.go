package interfaces

import (
	"context"

	"github.com/FianGumilar/e-commerce/models/dto"
	"github.com/FianGumilar/e-commerce/models/entity"
)

type AdminRepository interface {
	FindAllAdmin(ctx context.Context, offset int, limit int) ([]entity.Admin, error)
	FindOneAdminByID(ctx context.Context, id int) (*entity.Admin, error)
	FindOneAdminByEmail(ctx context.Context, email string) (*entity.Admin, error)
	InsertAdmin(ctx context.Context, admin entity.Admin) (*entity.Admin, error)
	UpdateAdmin(ctx context.Context, admin entity.Admin) (*entity.Admin, error)
	DeleteAdmin(ctx context.Context, admin entity.Admin) error
	CountAdmin(ctx context.Context) (int64, error)
}

type AdminService interface {
	FindAllAdmin(ctx context.Context, offset int, limit int) ([]entity.Admin, error)
	FindOneAdminByID(ctx context.Context, id int) (*entity.Admin, error)
	FindOneAdminByEmail(ctx context.Context, email string) (*entity.Admin, error)
	InsertAdmin(ctx context.Context, admin dto.AdminRequest) (*entity.Admin, error)
	UpdateAdmin(ctx context.Context, id int, admin dto.AdminRequest) (*entity.Admin, error)
	DeleteAdmin(ctx context.Context, id int) error
	CountAdmin(ctx context.Context) (int64, error)
}
