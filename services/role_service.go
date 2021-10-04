package services

import (
	"go-fx-test/lib"
	"go-fx-test/models"
	"go-fx-test/repository"
)

type RoleService struct {
	repository repository.RoleRepository
}

func NewRoleService(
	repository repository.RoleRepository,
) RoleService {
	return RoleService{
		repository: repository,
	}
}

func (s RoleService) GetAllRoles() (roles []models.Role, err error) {
	return roles, s.repository.Find(&roles).Error
}

func (s RoleService) Create(role models.Role) error {
	return s.repository.Create(&role).Error
}

func (s RoleService) GetOneRole(roleID lib.BinaryUUID) (role models.Role, err error) {
	return role, s.repository.First(&role, "id = ?", roleID).Error
}

func (s RoleService) UpdateRole(role models.Role) error {
	if err := s.repository.Model(&models.Role{}).Where("id = ?", role.ID).Updates(map[string]interface{}{
		"role":       role.Role,
		"department": role.Department,
	}).Error; err != nil {
		return err
	}
	return nil
}
