package repository

import (
	"testing"
	"time"
	"zupload/config"
	"zupload/model"
)

func TestCreateOne(t *testing.T) {
	err := config.InitForTest()
	if err != nil {
		t.Fatal(err)
	}

	fileRepo := NewFileRpo()

	id, err := fileRepo.CreateOne(&model.File{
		FileName:  "zhougang6",
		CreatedAt: time.Now().Unix(),
	})

	if err != nil {
		t.Log(err)
		return
	}

	t.Log(id)
}

func TestList(t *testing.T) {
	err := config.InitForTest()
	if err != nil {
		t.Fatal(err)
	}

	fileRepo := NewFileRpo()

	list, err := fileRepo.List(10)

	if err != nil {
		t.Log(err)
		return
	}

	t.Log(list)
}
