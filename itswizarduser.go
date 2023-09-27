package itswizard_m_basic

import "github.com/jinzhu/gorm"

/*
Ein nutzer auf itswizard
*/
type DbItswizardUser15 struct {
	gorm.Model
	Tel              string `gorm:"not null; type:VARCHAR(45)" json:"tel"`
	Firstname        string `gorm:"not null; type:VARCHAR(45)"  json:"firstname"`
	Lastname         string `gorm:"not null; type:VARCHAR(45)"  json:"lastname"`
	Hash             string `gorm:"not null; type:VARCHAR(500)" json:"hash"`
	Username         string `gorm:"not null; unique; type:VARCHAR(45)" json:"username"`
	Email            string `gorm:"not null; type:VARCHAR(45)" json:"email"`
	TwoFac           bool   `gorm:"not null" json:"twofac"`
	Admin            bool   `gorm:"not null" json:"admin"`
	Superadmin       bool   `gorm:"not null" json:"superadmin"`
	Supportagent     bool   `gorm:"not null" json:"supportagent"`
	Integrationagent bool   `gorm:"not null" json:"integrationagent"`
	OrganisationID   uint   `gorm:"TYPE:integer REFERENCES DbOrganisation15" json:"organisationID"`
	Language         string `gorm:"not null; type:VARCHAR(45)" json:"language"`
}

type DbOrganisation15 struct {
	gorm.Model
	Name          string
	InstitutionID uint `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	UniventionID  string
	LusdId        string
}
