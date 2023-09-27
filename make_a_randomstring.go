package itswizard_m_basic

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func MakeARandomString() string {
	var err error
	tmpSlice := uuid.Must(uuid.NewV4(), err).String()
	tmp := strings.Split(tmpSlice, "-")
	output := ""
	for i := 0; i < len(tmp); i++ {
		output = output + tmp[i]
	}
	return output
}

func Test2() {
	fmt.Println("Test")
}
