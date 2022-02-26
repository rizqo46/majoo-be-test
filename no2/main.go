package main

import (
	"fmt"
	"no2-majoo-be-test/database"
	"no2-majoo-be-test/model"
	Repo "no2-majoo-be-test/repository"
)

func main() {
	db := database.Connect()
	repository := Repo.NewAreaRepo(db)
	ar := &model.Area{}
	if err := repository.InsertArea(10, 10, "persegi", ar); err != nil {
		fmt.Println(err.Error())
	}
}
