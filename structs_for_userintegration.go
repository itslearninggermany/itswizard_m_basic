package itswizard_m_basic

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DbPerson15 struct {
	ID                 string `gorm:"not null; type:VARCHAR(500); primary key" json:"id"` // Example: DbInstitution15ID++DbOrganisation15ID++SyncPersonKey"+PersonIDNumber
	SyncPersonKey      string `gorm:"not null; type:VARCHAR(45);" json:"sync_person_key"`
	FirstName          string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName           string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username           string `gorm:"not null; type:VARCHAR(45)" json:"username"`
	Profile            string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Password           string `gorm:"password; type:VARCHAR(45)" json:"password"`
	Email              string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone              string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile             string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1            string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2            string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode           string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City               string `gorm:"type:VARCHAR(45)" json:"city"`
	DbOrganisation15ID uint   `gorm:"TYPE:integer REFERENCES DbOrganisation15""`
	DbInstitution15ID  uint   `gorm:"TYPE:integer REFERENCES DbInstitution15" "`
}

type DbStudentParentRelationship15 struct {
	ID                   string `gorm:"primary key"` // Example: DbInstitution15ID++DbOrganisation15ID++StudentSyncPersonKey++ParentSyncPersonKey
	DbOrganisation15ID   uint   `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
	DbInstitution15ID    uint   `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	StudentSyncPersonKey string
	ParentSyncPersonKey  string
}

type DbMentorStudentRelationship15 struct {
	ID                   string `gorm:"primary key"` // Example: DbInstitution15ID++DbOrganisation15ID++MentorSyncPersonKey++StudentSyncPersonKey
	DbOrganisation15ID   uint   `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
	DbInstitution15ID    uint   `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	MentorSyncPersonKey  string
	StudentSyncPersonKey string
}

type IntegrationType struct {
	gorm.Model
	InstitutionID uint
	AAD           bool
	UCS           bool
	Berlin        bool
	BW            bool
}

type DbInstitution15 struct {
	gorm.Model
	Name         string
	CID          string `gorm:"unique"`
	ItslSiteName string `gorm:"unique"`
}

type DbGroup15 struct {
	ID                  string `gorm:"primary key"` //Example: DbInstitution15ID++DbOrganisation15ID++SyncID
	SyncID              string
	AzureGroupID        string `gorm:"type:VARCHAR(150)"`
	UniventionGroupName string `gorm:"type:VARCHAR(150)"`
	Name                string
	ParentGroupID       string // default: rootPointer
	Level               int    // default: 1
	IsCourse            bool
	DbInstitution15ID   uint `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	DbOrganisation15ID  uint `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
}

type DbGroupMembership15 struct {
	ID                 string `gorm:"primary key"` //Example: DbInstitution15ID++DbOrganisation15ID++GroupName++PersonSyncKey
	PersonSyncKey      string
	GroupName          string
	DbInstitution15ID  uint `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	DbOrganisation15ID uint `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
	Profile            string
}

type Membership struct {
	gorm.Model
	PersonSyncKey  string `gorm:"not null; type:VARCHAR(45);" json:"sync_person_key"`
	GroupSyncKey   string
	Organisation15 uint `gorm:"TYPE:integer"`
	Institution15  uint `gorm:"TYPE:integer"`
	Profile        string
}

// Neue durch Update 2.0 itswizard

type Person struct {
	gorm.Model
	PersonSyncKey  string `gorm:"not null; type:VARCHAR(45);" json:"sync_person_key"`
	FirstName      string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName       string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username       string `gorm:"not null; type:VARCHAR(45)" json:"username"`
	Profile        string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Password       string `gorm:"password; type:VARCHAR(45)" json:"password"`
	Email          string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone          string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile         string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1        string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2        string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode       string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City           string `gorm:"type:VARCHAR(45)" json:"city"`
	Organisation15 uint   `gorm:"TYPE:integer"`
	Institution15  uint   `gorm:"TYPE:integer"`
}

type Group struct {
	gorm.Model
	GroupSyncKey        string
	Name                string
	ParentGroupID       string // default: rootPointer
	IsCourse            bool
	AzureGroupID        string `gorm:"type:VARCHAR(150)"`
	UniventionGroupName string `gorm:"type:VARCHAR(150)"`
	Organisation15      uint   `gorm:"TYPE:integer"`
	Institution15       uint   `gorm:"TYPE:integer"`
	ToSync              bool
	ToDelete            bool
	AlreadyInTheSystem  bool
}

type HomeOrganisationDelete struct {
	gorm.Model
	Stammschule string
}

type StudentParentRelationship struct {
	gorm.Model
	DbOrganisation15ID   uint `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
	DbInstitution15ID    uint `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	StudentSyncPersonKey string
	ParentSyncPersonKey  string
}

type MentorStudentRelationship struct {
	gorm.Model
	DbOrganisation15ID   uint `gorm:"TYPE:integer REFERENCES DbOrganisation15"`
	DbInstitution15ID    uint `gorm:"TYPE:integer REFERENCES DbInstitution15"`
	MentorSyncPersonKey  string
	StudentSyncPersonKey string
}

type CsvSetup struct {
	gorm.Model
	OrganisationID uint `gorm:"unique"`
	InstitutionID  uint
	RootGroup      string
}

type ImsesSetup struct {
	gorm.Model
	OrganisationID uint `gorm:"unique"`
	InstitutionID  uint
	Username       string
	Password       string
	Endpoint       string
}

type SchildSetup struct {
	gorm.Model
	Prefix         string
	OrganisationID uint `gorm:"unique"`
	InstitutionID  uint
	RootGroup      string
	CourseSync     bool
}

type ChangeLog struct {
	gorm.Model
	UserOrGroup      string
	NewPerson        bool
	DeltePerson      bool
	UdpatePerson     bool
	CreateCourse     bool
	GroupImport      bool
	GroupDelete      bool
	MembershipImport bool
	MembershipDelete bool
	PSR              bool
	Error            string `gorm:"type:longtext"`
	OrganisationId   uint
	InstitutionID    uint
	InternalSyncID   uint
}

type CsvUpload struct {
	gorm.Model
	Filename       string
	Seperator      string
	Content        string `gorm:"type:LONGTEXT"`
	OrganisationID uint
	InstitutionID  uint
}

type UniventionService struct {
	gorm.Model
	OrganisationID            uint
	InsitutionID              uint
	RunPersonCrawler          bool
	RunGroupCrawler           bool
	Run                       bool
	SelectOrganisations       bool
	NewRun                    bool
	RunWithUsernameChange     bool
	UpdatingSystem            bool
	RunWithUpdate             bool
	RunWithUpdateGroupRenamed bool
}

type UniventionSelectOrganisation struct {
	gorm.Model
	OrganisationID uint
	InsitutionID   uint
	UniventionID   string
}

type UniventionSetup struct {
	gorm.Model
	OrganisationID                  uint
	InstitutionID                   uint
	AdminSpecification              bool
	MakeTeacherFirstnameToOneLetter bool
	MakeStudentFirstnameToOneLetter bool
	MakeStudentFirstnameToOneName   bool
	MakeTeacherFirstnameToOneName   bool
	EmailNotToSync                  bool
	SyncDisable                     bool
	DeleteEmptyGroups               bool
}

type UniventionAdminSpecifiaction struct {
	gorm.Model
	OrganisationID uint
	InstitutionID  uint
	AdminLastName  string
}

type Service struct {
	gorm.Model
	OrganisationID          uint `gorm:"unique"`
	InstitutionID           uint
	CSV                     bool
	Azure                   bool
	Schild                  bool
	DownloadExcel           bool
	Univention              bool
	Lusd                    bool
	Lusd_SSZB               bool // Berlin Lusdadmin
	SHAdmin                 bool //
	MVAdmin                 bool //
	BWAdmin                 bool //
	UniventionFullFirstname bool
	IServ                   bool
	BWSchoolAdmin           bool
	BWRest                  bool
}

/*
type Service struct {
	gorm.Model
	OrganisationID uint `gorm:"unique"`
	InstitutionID uint
	CSV  bool
	Azure bool
	Schild bool
	DownloadExcel bool
}
*/

type UniventionGroup struct {
	gorm.Model
	OrganisationID      uint
	GroupSyncKey        string `gorm:"type:MEDIUMTEXT"`
	Data                string `gorm:"type:LONGTEXT"`
	ToUpdate            bool
	ToDelete            bool
	Success             bool
	UniventionGroupname string `gorm:"type:MEDIUMTEXT"`
	Errorstring         string `gorm:"type:MEDIUMTEXT"`
	Error               bool
	UUID                string
}

/*
Special for SH. See the full firstname
*/
type UniventionPersonFullFirstName struct {
	gorm.Model
	PersonSyncKey string `gorm:"type:MEDIUMTEXT"`
}

type UniventionGroupAutoDelete struct {
	gorm.Model
	GroupID     string `gorm:"type:MEDIUMTEXT"`
	Error       bool
	Errorstring string `gorm:"type:MEDIUMTEXT"`
}

type UniventionPerson struct {
	gorm.Model
	PersonSyncKey                 string `gorm:"type:MEDIUMTEXT"`
	UUID                          string
	Data                          string `gorm:"type:LONGTEXT"`
	ToUpdate                      bool
	ToDelete                      bool
	Success                       bool
	Errorstring                   string `gorm:"type:MEDIUMTEXT"`
	Error                         bool
	FirstName                     string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName                      string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username                      string `gorm:"not null; type:VARCHAR(45)" json:"username"`
	Profile                       string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Email                         string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone                         string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile                        string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1                       string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2                       string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode                      string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City                          string `gorm:"type:VARCHAR(45)" json:"city"`
	OrganisationID                uint   `gorm:"TYPE:integer"`
	Stammschule                   string `gorm:"type:LONGTEXT"`
	GruppenMitgliedschaften       string `gorm:"type:LONGTEXT"`
	Schulmitgliedschaften         string `gorm:"type:LONGTEXT"`
	ForeignID                     string
	Disabled                      bool
	ToImport                      bool
	UpdateStammschule             bool
	UdpateFirstName               bool
	UdpateLastName                bool
	UdpateUsername                bool
	UdpateProfile                 bool
	UpdateGruppenMitgliedschaften bool
	UpdateSchulmitgliedschaften   bool
	UpdateEmail                   bool
	UpdateDisable                 bool
	Selected                      bool // Eine Schulmitgliedschaft kommt hinzu -> Alle Gruppenzugehörigkeiten löschen nciht vergessen
	Deselected                    bool // Eine Schulmitgliedschaft wird gelöscht
}

type UniventionUploadGroup struct {
	gorm.Model
	OrganisationID      uint
	GroupSyncKey        string `gorm:"type:MEDIUMTEXT"`
	Data                string `gorm:"type:LONGTEXT"`
	UniventionGroupName string `gorm:"type:MEDIUMTEXT"`
}

type UniventionUploadPerson struct {
	gorm.Model
	OrganisationID uint
	PersonSyncKey  string `gorm:"type:MEDIUMTEXT"`
	Username       string
	Data           string `gorm:"type:LONGTEXT"`
	KindOfUpdate   string
}

type UniventionInterpreterStatus struct {
	gorm.Model
	InstitutionID uint
	Person        bool
	Group         bool
	LastRun       string
}

type UniventionOrganisationSelect struct {
	gorm.Model
	InstitutionID  uint
	OrganisationID uint
	OUName         string
	Active         bool
}

type Siteadmin struct {
	gorm.Model
	UserID         uint
	UserName       string
	InstitutionID  uint
	OrganisationID uint
	Active         bool
}

type LusdSetup struct {
	gorm.Model
	ClientId           string
	ClientSecret       string
	ApiEndpoint        string
	InsecureSkipVerify bool
}

type LusdToItswizardError struct {
	gorm.Model
	Error string
}

type LusdCrawlerError struct {
	gorm.Model
	Error string
}

type LusdCrawlerLog struct {
	gorm.Model
	Synckey        string
	Name           string
	Emailsend      bool
	ImportOrUpdate bool
	Delete         bool
}

/*
type LusdPerson struct {
	gorm.Model
	PersonSyncKey   string `gorm:"unique; type:VARCHAR(100)"`
	Classes			string `gorm:"type:LONGTEXT"`    //[]string with ClassGroupIds
	Courses			string `gorm:"type:LONGTEXT"`    //[]string with ClassGroupIds
	Schools		    string `gorm:"type:LONGTEXT"`    //[]string with SchoolGroupIds
	ToUpdate 		bool
	ToDelete 		bool
	Success         bool
	Errorstring     string `gorm:"type:MEDIUMTEXT"`
	Error           bool
	Birthday 		string
	FirstName       string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName        string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username        string `gorm:"unique; not null; type:VARCHAR(45)" json:"username"`
	Profile         string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Email           string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone           string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile          string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1         string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2         string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode        string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City            string `gorm:"type:VARCHAR(45)" json:"city"`
	OriginalData    string `gorm:"type:LONGTEXT"`
	NoSync          bool
	Testuser        bool
	BlusDataOld     bool
	ParentToImport  bool
	ParentToUpdate  bool
	ParentToDelete  bool
	ParentSuccess   bool
	ParentError     bool
	ParentErrorString string `gorm:"type:MEDIUMTEXT"`
	SecondParentImport bool
	SecondParentToUpdate  bool
	SecondParentToDelete  bool
	SecondParentSuccess   bool
	SecondParentError     bool
	SecondParentErrorString string `gorm:"type:MEDIUMTEXT"`
	NoParentImport bool
}
*/

type LusdPerson struct {
	gorm.Model
	PersonSyncKey             string `gorm:"unique; type:VARCHAR(100)"`
	Classes                   string `gorm:"type:LONGTEXT"` //[]string with ClassGroupIds
	Courses                   string `gorm:"type:LONGTEXT"` //[]string with ClassGroupIds
	Schools                   string `gorm:"type:LONGTEXT"` //[]string with SchoolGroupIds
	ChildParentRelation       string `gorm:"type:LONGTEXT"` //[]string with SchoolGroupIds
	Errorstring               string `gorm:"type:MEDIUMTEXT"`
	Error                     bool
	Birthday                  string
	FirstName                 string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName                  string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username                  string `gorm:"unique; not null; type:VARCHAR(45)" json:"username"`
	Profile                   string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Email                     string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone                     string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile                    string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1                   string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2                   string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode                  string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City                      string `gorm:"type:VARCHAR(45)" json:"city"`
	OriginalData              string `gorm:"type:LONGTEXT"`
	NoSync                    bool
	Testuser                  bool
	BlusDataOld               bool   // Nicht mehr aktuell
	ParentToImport            bool   // Nicht mehr aktuell
	ParentToUpdate            bool   // Nicht mehr aktuell
	ParentToDelete            bool   // Nicht mehr aktuell
	ParentSuccess             bool   // Nicht mehr aktuell
	ParentError               bool   // Nicht mehr aktuell
	ParentErrorString         string `gorm:"type:MEDIUMTEXT"` // Nicht mehr aktuell
	SecondParentImport        bool   // Nicht mehr aktuell
	SecondParentToUpdate      bool   // Nicht mehr aktuell
	SecondParentToDelete      bool   // Nicht mehr aktuell
	SecondParentSuccess       bool   // Nicht mehr aktuell
	SecondParentError         bool   // Nicht mehr aktuell
	SecondParentErrorString   string `gorm:"type:MEDIUMTEXT"` // Nicht mehr aktuell
	NoParentImport            bool   // Nicht mehr aktuell
	ToInput                   bool
	ToUpdate                  bool
	ToDelete                  bool
	Success                   bool
	UpdateBirthday            bool
	UpdateProfile             bool
	UpdateClass               bool
	UpdateSchool              bool
	UpdateParentChildRelation bool
	UpdateNames               bool
}

type LusdClass struct {
	gorm.Model
	SchoolID               string
	ClassSyncKey           string `gorm:"unique; type:VARCHAR(100)"`
	ToUpdateOrImport       bool
	ToDelete               bool
	Success                bool
	LusdGroupname          string `gorm:"type:MEDIUMTEXT"`
	Errorstring            string `gorm:"type:MEDIUMTEXT"`
	Error                  bool
	NoSync                 bool
	ParentGroupToImport    bool
	ParentGroupToUpdate    bool
	ParentGroupToDelete    bool
	ParentGroupSuccess     bool
	ParentGroupError       bool
	ParentGroupErrorString string `gorm:"type:MEDIUMTEXT"`
}

type LusdExport struct {
	gorm.Model
	Type string
	Data string `gorm:"type:LONGTEXT"`
}

type LusdSchool struct {
	gorm.Model
	SchoolSyncKey    string `gorm:"unique; type:VARCHAR(100)"`
	SchoolNumber     string
	SchoolName       string
	ToUpdateOrImport bool
	ToDelete         bool
	Success          bool
	Errorstring      string `gorm:"type:MEDIUMTEXT"`
	Error            bool
	LastRun          time.Time
	NoSync           bool
	Parentportal     bool
}

type LusdCourse struct {
	gorm.Model
	KursBezeichnung   string
	KursDtGeaendertAm string
	KursFach          string
	KursJahrgang      string
	KursSchulform     string
	KursUID           string
	SchuleUID         string
	ToUpdateOrImport  bool
	ToDelete          bool
	Success           bool
	Errorstring       string `gorm:"type:MEDIUMTEXT"`
	Error             bool
	NoSync            bool
}

type LusdService struct {
	gorm.Model
	InsitutionID uint
	Run          bool
}

type AadSetup struct {
	gorm.Model
	InstitutionID     uint
	OrganisationID    uint
	RootGroupKey      string
	AadProfileSetting string
	TenantID          string
	ApplicationID     string
	ClientSecret      string
}

type SamlAuthGroup struct {
	gorm.Model
	InstitutionID  uint
	OrganisationID uint
	GroupSyncKey   string
	Role           string
}

type ProfilesToSync struct {
	gorm.Model
	InstitutionID  uint
	OrganisationID uint
	ProfileName    string
}

type SchoolAdmin struct {
	gorm.Model
	PersonSyncKey  string `gorm:"unique"`
	OrganisationID uint
	Firstname      string
	Lastname       string
	Active         bool
	Error          string
}

type SchoolAdminError struct {
	gorm.Model
	Error string
}

type SchoolAdminService struct {
	gorm.Model
	Institution uint
}

type AdminInstitutionMatch struct {
	gorm.Model
	UserID      uint
	Username    string
	Institution uint
}

type UpdateUcsDump struct {
	gorm.Model
	Content     string `gorm:"type:LONGTEXT"`
	Proceed     bool
	Cache       string `gorm:"type:LONGTEXT"`
	CacheToRun  bool
	CacheRunned bool
	ErrorString string
	Error       bool
}

type UpdateKelvin struct {
	gorm.Model
	Content     string `gorm:"type:LONGTEXT"`
	Proceed     bool
	Cache       string `gorm:"type:LONGTEXT"`
	CacheToRun  bool
	CacheRunned bool
	ErrorString string
	Error       bool
}

type SupportEmail struct {
	gorm.Model
	InsitutionID uint
	Email        string `gorm:"type:TEXT"`
	Service      string
}

/* TODO
type BWCluster struct {

}
*/

type BWSchoolAPI struct {
	//SchoolID string
	CreatedAt     string
	UpdatedAt     string
	SchoolNumber  string
	SchoolName    string
	SchoolType    string
	ClusterNumber int
}

type BWPersonAPI struct {
	PersonID  string `gorm:"unique; type:VARCHAR(100)"`
	CreatedAt string
	UpdatedAt string
	FirstName string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName  string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username  string `gorm:"unique; not null; type:VARCHAR(45)" json:"username"`
	Birthday  string
	Email     string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone     string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile    string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1   string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2   string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode  string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City      string `gorm:"type:VARCHAR(45)" json:"city"`
}

type BWClassesAPI struct {
	ClassID      string `gorm:"unique"`
	CreatedAt    string
	UpdatedAt    string
	Name         string
	SchoolNumber string
}

type BWPersonSchoolMembershipAPI struct {
	CreatedAt    string
	UpdatedAt    string
	Role         []string
	PersonID     string
	Schoolnumber string
}

type BWPersonClassMembershipAPI struct {
	CreatedAt string
	UpdatedAt string
	PersonID  string
	ClassID   string
}

type BWClusterAPI struct {
	SchoolNumber  string
	ClusterNumber uint
	ClusterSite   string
	SAMLMetadata  string
}

type BwCsvUpload struct {
	gorm.Model
	Delimeter      string
	UserID         uint   `gorm:"not null"`
	OrganisationID uint   `gorm:"not null"`
	Content        []byte `gorm:"not null;type:LONGTEXT"`
}

type BwSyncCache struct {
	gorm.Model
	UserID         uint   `gorm:"not null"`
	OrganisationID uint   `gorm:"not null"`
	Content        string `gorm:"not null;type:LONGTEXT"`
	Imported       bool
}

type ResponseAPI struct {
	Error  string
	Result interface{}
}

type CaritasDataInput struct {
	gorm.Model
	CaritasFileID string
	Timestamp     string
	Discription   string
	Data          string `gorm:"type:LONGTEXT"`
	Processed     bool
}

type IServPerson struct {
	gorm.Model
	OrganisationID uint
	PersonSyncKey  string `gorm:"unique; type:VARCHAR(100)"`
	ToUpdate       bool
	ToDelete       bool
	Success        bool
	Errorstring    string `gorm:"type:MEDIUMTEXT"`
	Error          bool
	Birthday       string
	FirstName      string `gorm:"not null; type:VARCHAR(45)" json:"first_name"`
	LastName       string `gorm:"not null; type:VARCHAR(45)" json:"last_name"`
	Username       string `gorm:"unique; not null; type:VARCHAR(45)" json:"username"`
	Profile        string `gorm:"not null; type:VARCHAR(45)" json:"profile"`
	Email          string `gorm:"type:VARCHAR(45); primary key" json:"email"`
	Phone          string `gorm:"type:VARCHAR(45)" json:"phone"`
	Mobile         string `gorm:"type:VARCHAR(45)" json:"mobile"`
	Street1        string `gorm:"type:VARCHAR(45)" json:"street_1"`
	Street2        string `gorm:"type:VARCHAR(45)" json:"street_2"`
	Postcode       string `gorm:"type:VARCHAR(45)" json:"postcode"`
	City           string `gorm:"type:VARCHAR(45)" json:"city"`
}

type IServGroup struct {
	gorm.Model
	OrganisationID   uint
	GroupSyncKey     string `gorm:"unique; type:VARCHAR(100)"`
	GroupName        string
	Members          string `gorm:"type:LONGTEXT"`
	ToUpdateOrImport bool
	ToDelete         bool
	Success          bool
	Errorstring      string `gorm:"type:MEDIUMTEXT"`
	Error            bool
	ToSync           bool
}

type IServSetup struct {
	gorm.Model
	OrganisationID uint
	Intervall      uint
	AuthToken      string
	Domain         string
}

type IServSync struct {
	gorm.Model
	OrganisationID uint
	ChangeLog      string `gorm:"type:LONGTEXT"`
}

type AADSyncData struct {
	gorm.Model
	InstitutionID        uint
	OrganisationID       uint
	ErrorCreatingCache   string
	ErrorProcessingCahce string
	Finished             bool
	SyncCache            string `gorm:"type:LONGTEXT"`
	UsersToImport        uint
	UsersToDelete        uint
	UsersInItslearning   uint
	EmailSend            bool
}

type AADFilter struct {
	gorm.Model
	InstitutionID  uint
	OrganisationID uint
	Filter         bool
	Groupchanges   uint
	Personchanges  uint
	ChangesinTotal uint
}

type MonitoringUCSFileToItslearning struct {
	gorm.Model
	Username                          string
	UserUUID                          string
	UUID                              string
	UserInUCSFile                     string `gorm:"type:LONGTEXT"`
	UserInItsl                        string `gorm:"type:LONGTEXT"`
	MembershipInUCSFile               string `gorm:"type:LONGTEXT"`
	MembershipInItslFile              string `gorm:"type:LONGTEXT"`
	LogicalErrorsInUCS                string `gorm:"type:LONGTEXT"`
	SyncErrorsMiddlewareToIrslearning string `gorm:"type:LONGTEXT"`
	WithItslearning                   bool
	Error                             string
}

type MonitoringUCSFileToItslearningProzess struct {
	gorm.Model
	UUID  string
	Ready bool
}

type MonitoringFileItslearning struct {
	gorm.Model
	InstitutionID uint
	UUID          string
	Data          string `gorm:"type:LONGTEXT"`
}

type UcsDatabaseToItslearning struct {
	gorm.Model
	InstitutionID uint
	Error         string
}

type UcsProtokoll struct {
	gorm.Model
	Username    string
	UUID        string
	Action      string
	Success     bool
	Errorstring string
}

type MonitorWizardItsl struct {
	gorm.Model
	MonitorData string `gorm:"type:LONGTEXT"`
}

type AuthGroupMembers struct {
	gorm.Model
	InstitutionID  uint
	OrganisationID uint
	IdInDb         uint
	Members        string `gorm:"type:LONGTEXT"`
}

type UcsTeacherGroupName struct {
	gorm.Model
	Name string `gorm:"type:VARCHAR(500)"`
}
