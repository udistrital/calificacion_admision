package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearCampoCuposEspeciales_20200131_151227 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearCampoCuposEspeciales_20200131_151227{}
	m.Created = "20200131_151227"

	migration.Register("CrearCampoCuposEspeciales_20200131_151227", m)
}

// Run the migrations
func (m *CrearCampoCuposEspeciales_20200131_151227) Up() {
	m.SQL("ALTER TABLE evaluacion_inscripcion.cupos_por_dependencia ADD COLUMN cupos_especiales json;")

}

// Reverse the migrations
func (m *CrearCampoCuposEspeciales_20200131_151227) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
