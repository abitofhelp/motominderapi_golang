// Package motorcycleRepositoryTests implements unit tests for the MotorcycleRepository.
package motorcycleRepositoryTests

import (
	"testing"

	"github.com/abitofhelp/motominderapi/clean/domain/entities"
	"github.com/abitofhelp/motominderapi/clean/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

// TestMotorcycleRepository_List verifies that an empty list of motorcycles is returned.
func TestMotorcycleRepository_ListEmpty(t *testing.T) {

	// ARRANGE

	// ACT
	repo, _ := repositories.NewMotorcycleRepository()

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 0)
}

// TestMotorcycleRepository_ListOfOne verifies that a list with one motorcycle is returned.
func TestMotorcycleRepository_ListOfOne(t *testing.T) {

	// ARRANGE

	// ACT
	repo, _ := repositories.NewMotorcycleRepository()

	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	repo.Insert(motorcycle)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 1)
}

// TestMotorcycleRepository_Insert verifies that an insert is successful.
func TestMotorcycleRepository_Insert(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)

	// ACT
	moto, _ := repo.Insert(motorcycle)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 1)
	assert.True(t, *moto == repo.Motorcycles[0])
}

// TestMotorcycleRepository_Insert verifies that an insert is successful.
func TestMotorcycleRepository_Insert_IDAlreadyExists(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)

	// ACT
	moto, err := repo.Insert(motorcycle)
	_, err = repo.Insert(moto)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleRepository_FindByID verifies that an insert is successful.
func TestMotorcycleRepository_FindByID(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	moto, _ := repo.Insert(motorcycle)

	// ACT
	foundMoto, _ := repo.FindByID(moto.ID)

	// ASSERT
	assert.True(t, moto.ID == foundMoto.ID)
}

// TestMotorcycleRepository_FindByID verifies that an insert is successful.
func TestMotorcycleRepository_FindByID_NotExist(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()

	// ACT
	foundMoto, _ := repo.FindByID(123)

	// ASSERT
	assert.Nil(t, foundMoto)
}

// TestMotorcycleRepository_Update verifies that an update is successful.
func TestMotorcycleRepository_Update(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	moto, _ := repo.Insert(motorcycle)
	moto.Make = "Harley Davidson"

	// ACT
	repo.Update(moto)

	// ASSERT
	assert.True(t, repo.Motorcycles[0].Make == "Harley Davidson")
}

// TestMotorcycleRepository_Update_NotExist verifies that an update
// fails if the entity does not exist.
func TestMotorcycleRepository_Update_NotExist(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	motorcycle.ID = 123

	// ACT
	foundMoto, _ := repo.Update(motorcycle)

	// ASSERT
	assert.Nil(t, foundMoto)
}

// TestMotorcycleRepository_Delete verifies that a delete is successful.
func TestMotorcycleRepository_Delete(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	moto, _ := repo.Insert(motorcycle)

	// ACT
	repo.Delete(moto)

	// ASSERT
	assert.True(t, len(repo.Motorcycles) == 0)
}

// TestMotorcycleRepository_Delete_NotExist verifies that a delete
// fails if the entity does not exist.
func TestMotorcycleRepository_Delete_NotExist(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()
	motorcycle, _ := entities.NewMotorcycle("Honda", "Shadow", 2006)
	motorcycle.ID = 123

	// ACT
	err := repo.Delete(motorcycle)

	// ASSERT
	assert.NotNil(t, err)
}

// TestMotorcycleRepository_Save verifies that a save of the unit of work/dbContext is successful.
func TestMotorcycleRepository_Save(t *testing.T) {

	// ARRANGE
	repo, _ := repositories.NewMotorcycleRepository()

	// ACT
	err := repo.Save()

	// ASSERT
	assert.Nil(t, err)
}
