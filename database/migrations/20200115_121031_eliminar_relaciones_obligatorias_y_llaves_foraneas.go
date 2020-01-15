package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarRelacionesObligatoriasYLlavesForaneas_20200115_121031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarRelacionesObligatoriasYLlavesForaneas_20200115_121031{}
	m.Created = "20200115_121031"

	migration.Register("EliminarRelacionesObligatoriasYLlavesForaneas_20200115_121031", m)
}

// Run the migrations
func (m *EliminarRelacionesObligatoriasYLlavesForaneas_20200115_121031) Up() {
	file, err := ioutil.ReadFile("../scripts/20200115_121031_eliminar_relaciones_obligatorias_y_llaves_foraneas.up.sql")

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
