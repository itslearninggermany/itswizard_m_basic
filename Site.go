package itswizard_m_basic

import (
	"github.com/itslearninggermany/itswizard_m_objects"
	"html/template"
)

func r() {

}

type Site struct {
	Navigation  template.HTML
	SessionUser itswizard_m_objects.SessionUser
	Sitename    string
	Special     interface{}
	Attributes  string
}
