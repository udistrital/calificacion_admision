package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarOpcion23_20240731_173704 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarOpcion23_20240731_173704{}
	m.Created = "20240731_173704"

	migration.Register("EliminarOpcion23_20240731_173704", m)
}

// Run the migrations
func (m *EliminarOpcion23_20240731_173704) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito ADD COLUMN examen_estado boolean;")
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico ADD COLUMN opcion json;")
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico ADD COLUMN puntaje_minimo_examen integer;")

}

// Reverse the migrations
func (m *EliminarOpcion23_20240731_173704) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito DROP COLUMN examen_estado;")
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico DROP COLUMN opcion;")
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico DROP COLUMN puntaje_minimo_examen;")

}
