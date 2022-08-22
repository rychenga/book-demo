// demo project main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//sqlite talbe
type BOOKS struct {
	gorm.Model
	NAME     string
	BORROWER string
}

func main() {

	fmt.Println("*****Web Service Start*****")
	fmt.Printf("\n")

	_gin := gin.Default() //宣告_gin元件

	_gin.GET("/ping", ping) //Http ping 指令

	_gin.GET("/GET", RGet) //RESTful GET 指令

	_gin.GET("/PUT", RPut) //RESTful PUT 指令

	_gin.GET("/POST", RPost) //RESTful POST 指令

	_gin.GET("/DELETE", RDEL) //RESTful DELETE 指令

	_gin.Run(":8080") //執行WebService part:8080

}

//HTTP ping 指令
func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})

}

//RESTful GET 指令
func RGet(c *gin.Context) {

	//查詢條件
	_name := c.DefaultQuery("name", "")

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("sqlite connect error")
	}
	//Migrate the schema
	db.AutoMigrate(&BOOKS{})

	//初始化book from BOOKS
	book := make([]BOOKS, 0)

	if _name != "" {
		db.Table("books").Where("NAME = ?", _name).Find(&book)
		c.String(http.StatusOK, "name: %s, msg: %s ", _name, book)

	} else {
		db.Table("books").Find(&book)
		c.String(http.StatusOK, "msg: %s ", book)
	}

}

//RESTful POST(Create) 指令
func RPost(c *gin.Context) {

	//更新項目
	_name := c.DefaultQuery("name", "")
	_brrower := c.DefaultQuery("borrower", "")

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("sqlite connect error")
	}
	//Migrate the schema
	db.AutoMigrate(&BOOKS{})

	if _name != "" {
		db.Create(&BOOKS{NAME: _name, BORROWER: _brrower})
		c.String(http.StatusOK, "create rows by name: %s ,brrower: %s", _name, _brrower)

	} else {
		c.String(http.StatusOK, "name can't empty!!")
	}

}

//RESTful PUT(UPDATE) 指令
func RPut(c *gin.Context) {

	//查詢條件
	_keyname := c.DefaultQuery("keyname", "")

	//更新項目
	_name := c.DefaultQuery("name", "")
	_brrower := c.DefaultQuery("borrower", "")

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("sqlite connect error")
	}
	//Migrate the schema
	db.AutoMigrate(&BOOKS{})

	//update 功能
	var book BOOKS

	if _keyname != "" {

		db.First(&book, "NAME = ?", _keyname)

		//name 欄位
		if _name != "" {
			db.Model(&book).Update("NAME", _name)
		}

		//_brrower 欄位
		if _brrower != "" {
			db.Model(&book).Update("BORROWER", _brrower)
		}

		c.String(http.StatusOK, "keyname: %s, name: %s ,brrower: %s", _keyname, _name, _brrower)

	} else {
		c.String(http.StatusOK, "keyname can't empty!!")
	}

}

//RESTful DELETE 指令
func RDEL(c *gin.Context) {

	//刪除項目
	_name := c.DefaultQuery("name", "")

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("sqlite connect error")
	}
	//Migrate the schema
	db.AutoMigrate(&BOOKS{})

	//update 功能
	var book BOOKS

	if _name != "" {
		db.First(&book, "NAME = ?", _name)
		db.Delete(&book)
		c.String(http.StatusOK, "delete rows by name: %s", _name)

	} else {
		c.String(http.StatusOK, "name can't empty!!")
	}

}
