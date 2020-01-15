package main

import (
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
 m.SQL("ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion
 update CONSTRAINT fk_requisito_programa_academico_requisito ON UPDATE CASCADE ON DELETE SET NULL;") 

}

// Reverse the migrations
func (m *EliminarRelacionesObligatoriasYLlavesForaneas_20200115_121031) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
