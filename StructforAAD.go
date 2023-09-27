package itswizard_m_basic

import "github.com/jinzhu/gorm"

type DbAAdLog struct {
	gorm.Model
	OrganisationId uint
	InstitutionId  uint
	Information    string
}
