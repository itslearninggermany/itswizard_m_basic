package itswizard_m_basic

import "github.com/jinzhu/gorm"

type MessageStruct struct {
	Headline     string
	Message      string
	TargetClose  string
	TargetSubmit string
	Buttontext   string
}

type DbSftpData15 struct {
	gorm.Model
	InstitutionId uint `gorm:"unique"`
	SftpUsername  string
	SftpPasswort  string
	SftpServer    string
}

type DbMakeCourse struct {
	gorm.Model
	OrganisationID uint `gorm:"unique"`
	Course         bool
}
