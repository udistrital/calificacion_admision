package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarEvaluacionInscripcion_20210201_153319 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarEvaluacionInscripcion_20210201_153319{}
	m.Created = "20210201_153319"

	migration.Register("EliminarEvaluacionInscripcion_20210201_153319", m)
}

// Run the migrations
func (m *EliminarEvaluacionInscripcion_20210201_153319) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210201_153319_eliminar_evaluacion_inscripcion.up.sql")

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
func (m *EliminarEvaluacionInscripcion_20210201_153319) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210201_153319_eliminar_evaluacion_inscripcion.down.sql")

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
