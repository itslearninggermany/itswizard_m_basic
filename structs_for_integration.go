package itswizard_m_basic

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ItlPlugModul struct {
	Auth ItslAuth
	Sign string
}

type ItslAuth struct {
	PostTo           string `json:"PostTo"`
	CustomerID       string `json:"CustomerId"`
	PersonID         string `json:"PersonId"`
	Language         string `json:"Language"`
	Country          string `json:"Country"`
	EducationalLevel string `json:"EducationalLevel"`
	Role             string `json:"Role"`
	EditReference    string `json:"EditReference"`
	FirstName        string `json:"FirstName"`
	LastName         string `json:"LastName"`
	OAuthToken       string `json:"OAuthToken"`
	OAuthTokenSecret string `json:"OAuthTokenSecret"`
	TimeStamp        string `json:"TimeStamp"`
}

type MediathekenSite struct {
	Link  string
	Daten string
}
type MedaiathekenElementBeschreibung struct {
	Titel        string
	ItslAuth     string
	ItslSign     string
	LowestID     string
	SearchString string
	PageIDs      string
	Beschreibung string
	Dauer        string
	Datum        string
	Webseite     string
	Url          string
	Link         string
}

type MedaiathekenElement struct {
	Id           uint
	Sender       string
	Titel        string
	Beschreibung string
	Zeit         string
	Link         string
	ItslAuth     string
	ItslSign     string
	LowestID     string
	SearchString string
	Page_ids     string
}

type IntervallStruct struct {
	Counter       uint
	DatabaseID    uint
	LinkCharacter string
	ItslAuth      string
	ItslSign      string
	LowestID      string
	SearchString  string
	Page_ids      string
}

type MedaiathekenSeite struct {
	Medien           []MedaiathekenElement
	ItslAuth         string
	ItslSign         string
	LowestID         string
	SelectedID       string
	SearchString     string
	Page_ids         string
	IntervallStructs []IntervallStruct
}

type MediathekenData struct {
	gorm.Model
	Sender       string
	Thema        string
	Titel        string
	Zeit         string
	Dauer        string
	Beschreibung string
	Url          string `gorm:"unique"`
	Website      string
	Datum        time.Time
	Geo          string
	ToDelete     bool
	ToUpdate     bool
}

type MediathekenSetup struct {
	gorm.Model
	DownloadFile string
	SharedSecret string
}

type MediathekenError struct {
	gorm.Model
	Error string
}

type MediathekenSenderUrl struct {
	gorm.Model
	Sender string
	Url    string
}
