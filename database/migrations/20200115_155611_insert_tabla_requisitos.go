package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTablaRequisitos_20200115_155611 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTablaRequisitos_20200115_155611{}
	m.Created = "20200115_155611"

	migration.Register("InsertTablaRequisitos_20200115_155611", m)
}

// Run the migrations
func (m *InsertTablaRequisitos_20200115_155611) Up() {
	file, err := ioutil.ReadFile("../scripts/20200115_155611_insert_tabla_requisitos.up.sql")

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
func (m *InsertTablaRequisitos_20200115_155611) Down() {
	m.SQL("DELETE FROM evaluacion_inscripcion.requisito WHERE codigo_abreviacion = 'ICFES';")
	m.SQL("DELETE FROM evaluacion_inscripcion.requisito WHERE codigo_abreviacion = 'ENTREVISTA';")
	m.SQL("DELETE FROM evaluacion_inscripcion.requisito WHERE codigo_abreviacion = 'PRUEBA';")

}
