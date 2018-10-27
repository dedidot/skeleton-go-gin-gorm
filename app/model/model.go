package model

import (
	"../../config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAll(b *[]Scheme) (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
/*
func AddNew(b *Scheme) (err error) {
	if err = DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOne(b *Scheme, id string) (err error) {
	if err := DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOne(b *Scheme, id string) (err error) {
	fmt.Println(b)
	DB.Save(b)
	return nil
}

func Delete(b *Scheme, id string) (err error) {
	DB.Where("id = ?", id).Delete(b)
	return nil
}*/