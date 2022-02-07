package task

import (
	"part3/configs"
	"part3/models/task"
	"part3/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&task.Task{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		mockTask := task.Task{Name_Task: "anonim123", Priority: 1}
		res, err := repo.Create(mockTask)
		log.Info(res, err)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run Create", func(t *testing.T) {
		mockTask := task.Task{Name_Task: "anonim123", Priority: 1}
		_, err := repo.Create(mockTask)
		log.Info(err)
		assert.NotNil(t, err)
	})
}

// func TestGetById(t *testing.T)  {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("success run GetById", func(t *testing.T) {
// 		res, err := repo.GetById(1)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, int(res.ID))
// 	})

// 	t.Run("fail run GetById", func(t *testing.T) {
// 		res, err := repo.GetById(2)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, 1, int(res.ID))
// 	})
// }

// func TestUpdateById(t *testing.T)  {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	repo := New(db)

// 	t.Run("success run UpdateById", func(t *testing.T) {
// 		mockBook := request.BookRequest{Name: "anonim321", Publisher: "anonim321", Author: "anonim321"}
// 		res, err := repo.UpdateById(1, mockBook)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "anonim321", res.Name)
// 	})

// 	t.Run("fail run UpdateById", func(t *testing.T) {
// 		mockBook := request.BookRequest{Name: "anonim456", Publisher: "anonim456", Author: "anonim456"}
// 		res, err := repo.UpdateById(2, mockBook)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, 1, int(res.ID))
// 	})
// }

// func TestDeleteById(t *testing.T)  {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	lib := New(db)

// 	t.Run("success run DeleteById", func(t *testing.T) {
// 		res, err := lib.DeleteById(1)
// 		assert.Nil(t, err)
// 		assert.Equal(t, true, res.Valid)
// 	})

// 	t.Run("fail run DeleteById", func(t *testing.T) {
// 		res, err := lib.DeleteById(1)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, false, res.Valid)
// 	})
// }

// func TestGetAll(t *testing.T)  {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)
// 	lib := New(db)

// 	t.Run("success run GetAll", func(t *testing.T) {
// 		res, err := lib.GetAll()
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 	})

// 	t.Run("fail run GetAll", func(t *testing.T) {
// 		db.Migrator().DropTable(&book.Book{})
// 		_, err := lib.GetAll()
// 		assert.NotNil(t, err)
// 	})
// }
