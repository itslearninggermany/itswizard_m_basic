package itswizard_m_basic

import (
	"fmt"
	"io/ioutil"
)

/*
Inhalte der Webseiten
*/
type DbHtmlContent struct {
	Id       string `gorm:"unique"` //language_Name
	Name     string `gorm:"not null"`
	Language string `gorm:"not null"`
	Field0   string `gorm:"type:MEDIUMTEXT"`
	Field1   string `gorm:"type:MEDIUMTEXT"`
	Field2   string `gorm:"type:MEDIUMTEXT"`
	Field3   string `gorm:"type:MEDIUMTEXT"`
	Field4   string `gorm:"type:MEDIUMTEXT"`
	Field5   string `gorm:"type:MEDIUMTEXT"`
	Field6   string `gorm:"type:MEDIUMTEXT"`
	Field7   string `gorm:"type:MEDIUMTEXT"`
	Field8   string `gorm:"type:MEDIUMTEXT"`
	Field9   string `gorm:"type:MEDIUMTEXT"`
	Field10  string `gorm:"type:MEDIUMTEXT"`
	Field11  string `gorm:"type:MEDIUMTEXT"`
	Field12  string `gorm:"type:MEDIUMTEXT"`
	Field13  string `gorm:"type:MEDIUMTEXT"`
	Field14  string `gorm:"type:MEDIUMTEXT"`
	Field15  string `gorm:"type:MEDIUMTEXT"`
	Field16  string `gorm:"type:MEDIUMTEXT"`
	Field17  string `gorm:"type:MEDIUMTEXT"`
	Field18  string `gorm:"type:MEDIUMTEXT"`
	Field19  string `gorm:"type:MEDIUMTEXT"`
	Field20  string `gorm:"type:MEDIUMTEXT"`
	Field21  string `gorm:"type:MEDIUMTEXT"`
	Field22  string `gorm:"type:MEDIUMTEXT"`
	Field23  string `gorm:"type:MEDIUMTEXT"`
	Field24  string `gorm:"type:MEDIUMTEXT"`
	Field25  string `gorm:"type:MEDIUMTEXT"`
	Field26  string `gorm:"type:MEDIUMTEXT"`
	Field27  string `gorm:"type:MEDIUMTEXT"`
}

func SaveHtmlStringsToFile(input map[string]string, path string, name string) error {
	for key, value := range input {

		err := ioutil.WriteFile(fmt.Sprint(path, "/", key, "_", name, ".", "html"), []byte(value), 0644)
		if err != nil {
			return err
			fmt.Println(err)
		}
	}
	return nil
}
