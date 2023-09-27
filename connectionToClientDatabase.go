package itswizard_m_basic

import (
	"encoding/json"
	"github.com/itslearninggermany/itswizard_m_awsbrooker"
	"github.com/jinzhu/gorm"
)

type databases map[string]*gorm.DB

type DatabaseConfig struct {
	NameOrCID string `json:"name"`
	Dialect   string `json:"dialect"`
	Password  string `json:"pasword"`
	Username  string `json:"username"`
	Host      string `json:"host"`
}

/*
path has to end with an "/" dsf
*/
func NewDatabases(filename string, bucketname string) (out databases, err error) {
	var databaseConfig []DatabaseConfig
	b, _ := itswizard_m_awsbrooker.DownloadFileFromBucket(bucketname, filename)
	err = json.Unmarshal(b, &databaseConfig)
	if err != nil {
		return nil, err
	}

	a := make(map[string]*gorm.DB)
	for i := 0; i < len(databaseConfig); i++ {
		database, err := gorm.Open(databaseConfig[i].Dialect, databaseConfig[i].Username+":"+databaseConfig[i].Password+"@tcp("+databaseConfig[i].Host+")/"+databaseConfig[i].NameOrCID+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			return nil, err
		}

		a[databaseConfig[i].NameOrCID] = database
	}
	return a, nil
}

func (p databases) GetDatabase(input string) *gorm.DB {
	return p[input]
}
