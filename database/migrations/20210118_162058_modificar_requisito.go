package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarRequisito_20210118_162058 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarRequisito_20210118_162058{}
	m.Created = "20210118_162058"

	migration.Register("ModificarRequisito_20210118_162058", m)
}

// Run the migrations
func (m *ModificarRequisito_20210118_162058) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210118_162058_modificar_requisito.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *ModificarRequisito_20210118_162058) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210118_162058_modificar_requisito.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
