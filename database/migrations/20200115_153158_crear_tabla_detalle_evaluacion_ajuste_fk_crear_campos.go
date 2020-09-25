package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDetalleEvaluacionAjusteFkCrearCampos_20200115_153158 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDetalleEvaluacionAjusteFkCrearCampos_20200115_153158{}
	m.Created = "20200115_153158"

	migration.Register("CrearTablaDetalleEvaluacionAjusteFkCrearCampos_20200115_153158", m)
}

// Run the migrations
func (m *CrearTablaDetalleEvaluacionAjusteFkCrearCampos_20200115_153158) Up() {
	file, err := ioutil.ReadFile("../scripts/20200115_153158_crear_tabla_detalle_evaluacion_ajuste_fk_crear_campos.up.sql")

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
func (m *CrearTablaDetalleEvaluacionAjusteFkCrearCampos_20200115_153158) Down() {
	file, err := ioutil.ReadFile("../scripts/20200115_153158_crear_tabla_detalle_evaluacion_ajuste_fk_crear_campos.down.sql")

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
