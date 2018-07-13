package model

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nu7hatch/gouuid"
	"logger"
	"dao"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `sql:"DEFAULT:now();not null" json:"created_at"`
	UpdatedAt time.Time `sql:"DEFAULT:now();not null" json:"updated_at"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`
}

type Product struct {
	BaseModel
	Code  string `gorm:"unique;not null;size:255"`
	Price uint
}

type User struct {
	BaseModel
	Name   string `gorm:"type:varchar(100);unique_index;not null" json:"name" valid:"required~Name is blank"`
	Team   string `gorm:"type:varchar(15)" json:"team" valid:"required~Team is blank"`
	Member string `gorm:"size:255"`
}

// Set field `AnimalID` as primary field
type Animal struct {
	AnimalID  string    `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(100);unique;not null"`
	Age       int64     `gorm:"default:5"`
	CreatedAt time.Time `sql:"DEFAULT:current_timestamp" gorm:"not null"`
	UpdatedAt time.Time `sql:"DEFAULT:now()" gorm:"not null" json:"updated_at"`
}

func (animal *Animal) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		logger.Log.Error("Error creating uuid: ", err)
		return err
	}
	scope.SetColumn("AnimalID", uuid)
	return nil
}

func FetchById(v *User, id int) (*gorm.DB) {
	db := dao.GetDb()
	return db.Where("id = ?", id).First(&v)
}

func CreateEntry(e interface{}) (interface{}) {
	db := dao.GetDb()
	c := db.Create(e)
	logger.Log.Infof("Rows Affected %s", c.RowsAffected)
	return c.Value
}
