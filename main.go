package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Command to create a new table
//_ = db.Exec("CREATE TABLE notes ( id int(11) NOT NULL AUTO_INCREMENT, title varchar(100) NOT NULL, content text, image json DEFAULT NULL, has_finised tinyint(1) DEFAULT '0', status int(11) DEFAULT '1', created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (id))\n")

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {
	dsn := os.Getenv("DBConnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to run ")
	}
	// Insert new note
	//newNote := Note{Title: "Demo note", Content: "This is content of demo note"}
	//if err := db.Create(&newNote); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(newNote)

	// Select note
	var notes []Note
	db.Where("status = ?", 1).Find(&notes)

	var note Note
	if err := db.Where("id = ?", 4).First(&note); err != nil {
		log.Println(err)
	}

	fmt.Println(notes)

	// Delete note
	db.Table(Note{}.TableName()).Where("id = 4").Delete(nil)

	// Update note
	db.Table(Note{}.TableName()).Where("id = 3").Updates(map[string]interface{}{
		"title": "Demo 2",
	})
}
