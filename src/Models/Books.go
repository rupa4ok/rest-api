package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	Config2 "rest/Config"
)

func GetAllBook(b *[]Book) (err error) {
	if err = Config2.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = Config2.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err := Config2.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	Config2.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string) (err error) {
	Config2.DB.Where("id = ?", id).Delete(b)
	return nil
}
