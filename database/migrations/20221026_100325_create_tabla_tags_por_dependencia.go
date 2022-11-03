package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTablaTagsPorDependencia_20221026_100325 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTablaTagsPorDependencia_20221026_100325{}
	m.Created = "20221026_100325"

	migration.Register("CreateTablaTagsPorDependencia_20221026_100325", m)
}

// Run the migrations
func (m *CreateTablaTagsPorDependencia_20221026_100325) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20221026_100325_create_tabla_tags_por_dependencia.up.sql")

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
func (m *CreateTablaTagsPorDependencia_20221026_100325) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20221026_100325_create_tabla_tags_por_dependencia.down.sql")

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
