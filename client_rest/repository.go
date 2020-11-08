package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type (
	Client struct {
		gorm.Model
		Code      string `gorm: "column:'rct_code';type:text;default:'';unique_index"`
		Name      string `gorm: "column:'rct_name';default:'';type:text"`
		Dba       string `gorm: "column:'rct_dba';default:'';type:text"`
		Addr1     string `gorm: "column:'rct_addr1';default:'';type:text"`
		Addr2     string `gorm: "column:'rct_addr2';default:'';type:text"`
		City      string `gorm: "column:'rct_city';default:'';type:text"`
		State     string `gorm: "column:'rct_state';default:'';type:text"`
		Zip       string `gorm: "column:'rct_zip';default:'';type:text"`
		Country   string `gorm: "column:'rct_country';default:'';type:text"`
		Landline  string `gorm: "column:'rct_landline';default:'';type:text"`
		Cellphone string `gorm: "column:'rct_cellphone';default:'';type:text"`
		Fax       string `gorm: "column:'rct_fax';default:'';type:text"`
		Email     string `gorm: "column:'rct_email';default:'';type:text"`
	}

	ClientRepository struct {
		Repository
		Db     *gorm.DB
		Client *Client
	}
)

const (
	linesperpage = 50
)

func (Client) TableName() string {
	return "rct_clients"
}

func NewClinetRepository(db *gorm.DB, client *Client) *ClientRepository {
	db.AutoMigrate(&Client{})
	return &ClientRepository{
		Db:     db,
		Client: client,
	}
}

func Get(id uint) []Client {
	outdata := make([]Client, 0)
	tbl := db.Table("rct_clients")
	tbl = tbl.Select("rct_clients.id,rct_clients.rct_name")
	tbl = tbl.Where("rct_clients.id = ?", id)
	tbl.Find(&outdata)
	return outdata
}

func List(page uint) []Client {
	outdata := make([]Client, 0)
	tbl := db.Table("rct_clients")
	tbl = tbl.Select("rct_clients.*")
	//	tbl = tbl.Where("rct_clients.id = ?", filterStr)
	offset := linesperpage * page
	db.Limit(linesperpage).Offset(offset)
	tbl.Find(&outdata)
	return outdata
}

func Search(fltstr string) []Client {
	outdata := make([]Client, 0)
	tbl := db.Table("rct_clients")
	tbl = tbl.Select("rct_clients.*")
	searchFields := []string{"rct_code", "rct_name", "rct_dba"}
	filterList := strings.Split(fltStr, " ")
	for _, a := range filterList {
		if len(a) == 0 {
			continue
		}
		tbl = tbl.Or(fmt.Sprintf("rct_clients.rct_code like ?", fmt.Sprintf("%%%s%%", a)))
	}
	offset := linesperpage * page
	db.Limit(linesperpage).Offset(offset)
	tbl.Find(&outdata)

	return outdata
}

func Save(client *Client) bool {
	outdata := make([]Client, 0)
	tbl := db.Table("rct_clients")
	tbl = tbl.Select("rct_clients.Id")
	tbl = tbl.Where("rct_clients.id = ?", (*client).ID)
	tbl.Find(&outdata)
	if len(outdata) == 0 {
		_ = db.Create(client)
	} else if len(outdata) == 1 {
		_ = db.Save(client)
	} else {
		return false
	}

	return true
}
func Delete(id uint) bool {
	client := Client{}
	db.Where("id = ?", id).Delete(&client)
	return true
}
