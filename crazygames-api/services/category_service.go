package services

import (
	"crazygames.io/entities"
	"crazygames.io/handler/request"
	"crazygames.io/repositories"
	"io"
	"os"
)

type CategoryService struct {
	CategoryRepo repositories.CategoryRepositoryInterface
	MinioService MinIOServiceInterface
}

type CategoryServiceInterface interface {
	CreateCategory(request *request.CategoryRequestCreate) (*entities.Category, error)
	GetAll() ([]entities.Category, error)
	GetMenu() ([]entities.Category, error)
	GetByID(id uint) (*entities.Category, error)
	Update(request *request.CategoryRequestUpdate, id uint) (*entities.Category, error)
	Delete(id uint) error
}

func NewCategoryService(categoryRepo repositories.CategoryRepositoryInterface, minioService MinIOServiceInterface) *CategoryService {
	return &CategoryService{CategoryRepo: categoryRepo, MinioService: minioService}
}

func (ms *CategoryService) CreateCategory(request *request.CategoryRequestCreate) (*entities.Category, error) {
	// Upload icon to MinIO
	// Create temp file from multipart file
	tempFile, err := os.CreateTemp("", "category-icon-*")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	file, err := request.Icon.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, err
	}
	tempFile.Close()

	// Upload to MinIO
	iconURL, err := ms.MinioService.UploadFile(tempFile.Name(), "categories/"+request.Icon.Filename, "")
	if err != nil {
		return nil, err
	}
	// Create Category with icon URL
	isMenu := false
	if request.IsMenu == "1" {
		isMenu = true
	}
	category := &entities.Category{
		CategoryName: request.CategoryName,
		Description:  request.Description,
		Icon:         iconURL,
		Path:         request.Path,
		IsMenu:       isMenu,
	}
	err = ms.CategoryRepo.Create(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (ms *CategoryService) GetAll() ([]entities.Category, error) {
	category, err := ms.CategoryRepo.GetAll()
	if err != nil {
		return []entities.Category{}, err
	}
	return category, nil
}

func (ms *CategoryService) GetMenu() ([]entities.Category, error) {
	category, err := ms.CategoryRepo.GetMenu()
	if err != nil {
		return []entities.Category{}, err
	}
	return category, nil
}

func (ms *CategoryService) GetByID(id uint) (*entities.Category, error) {
	category, err := ms.CategoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (ms *CategoryService) Update(request *request.CategoryRequestUpdate, id uint) (*entities.Category, error) {
	category, err := ms.GetByID(id)
	if err != nil {
		return nil, err
	}
	iconURL := category.Icon
	if request.Icon != nil {
		// Create temp file from multipart file
		tempFile, err := os.CreateTemp("", "category-icon-*")
		if err != nil {
			return nil, err
		}
		defer os.Remove(tempFile.Name())

		file, err := request.Icon.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			return nil, err
		}
		tempFile.Close()

		// Upload to MinIO
		iconURL, err = ms.MinioService.UploadFile(tempFile.Name(), "categories/"+request.Icon.Filename, "")
		if err != nil {
			return nil, err
		}
	}
	isMenu := false
	if request.IsMenu == "1" {
		isMenu = true
	}
	categoryData := &entities.Category{
		ID:           id,
		CategoryName: request.CategoryName,
		Description:  request.Description,
		Icon:         iconURL,
		Path:         request.Path,
		IsMenu:       isMenu,
	}
	return ms.CategoryRepo.Update(categoryData)
}

func (ms *CategoryService) Delete(id uint) error {
	err := ms.CategoryRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
