package admin

import (
	"context"

	"github.com/FianGumilar/e-commerce/interfaces"
	"github.com/FianGumilar/e-commerce/models/dto"
	"github.com/FianGumilar/e-commerce/models/entity"
	"github.com/FianGumilar/e-commerce/utils"
)

type service struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminService(adminRepoitory interfaces.AdminRepository) interfaces.AdminService {
	return &service{
		adminRepository: adminRepoitory,
	}
}

// FindAllAdmin implements interfaces.AdminService.
func (s service) FindAllAdmin(ctx context.Context, offset int, limit int) ([]entity.Admin, error) {
	res, err := s.adminRepository.FindAllAdmin(ctx, offset, limit)
	if err != nil {
		return nil, utils.HandleError("Failed to get admin", 500)
	}
	return res, nil
}

// FindOneAdminByID implements interfaces.AdminService.
func (s service) FindOneAdminByID(ctx context.Context, id int) (*entity.Admin, error) {
	res, err := s.adminRepository.FindOneAdminByID(ctx, id)
	if err != nil {
		return nil, utils.HandleError("admin ID not found", 404)
	}
	return res, nil
}

// FindOneAdminByEmail implements interfaces.AdminService.
func (s service) FindOneAdminByEmail(ctx context.Context, email string) (*entity.Admin, error) {
	res, err := s.adminRepository.FindOneAdminByEmail(ctx, email)
	if err != nil {
		return nil, utils.HandleError("admin email not found", 404)
	}
	return res, nil
}

// InsertAdmin implements interfaces.AdminService.
func (s service) InsertAdmin(ctx context.Context, admin dto.AdminRequest) (*entity.Admin, error) {
	hassPassword := utils.HashPassword(*admin.Password)
	data := entity.Admin{
		Email:    admin.Email,
		Name:     admin.Name,
		Password: hassPassword,
	}

	if admin.CreatedBy != nil {
		data.CreatedByID = admin.CreatedBy
	}

	res, err := s.adminRepository.InsertAdmin(ctx, data)
	if err != nil {
		return nil, utils.HandleError("failed to insert admin", 500)
	}

	return res, nil
}

// UpdateAdmin implements interfaces.AdminService.
func (s service) UpdateAdmin(ctx context.Context, id int, admin dto.AdminRequest) (*entity.Admin, error) {
	panic("unimplemented")
}

// DeleteAdmin implements interfaces.AdminService.
func (s service) DeleteAdmin(ctx context.Context, id int) error {
	panic("unimplemented")
}

// CountAdmin implements interfaces.AdminService.
func (s service) CountAdmin(ctx context.Context) (int64, error) {
	panic("unimplemented")
}
