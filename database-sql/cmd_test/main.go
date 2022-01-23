package main

import (
	"geeorm"
	"geeorm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee s")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS user;").Exec()
	_, _ = s.Raw("CREATE TABLE user(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE user(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO user(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	log.Infof("Exec success, %d affected\n", count)
}
