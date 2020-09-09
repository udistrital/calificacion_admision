package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambioTipoCampoActivoDetalleEvaluacion_20200127_212423 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambioTipoCampoActivoDetalleEvaluacion_20200127_212423{}
	m.Created = "20200127_212423"

	migration.Register("CambioTipoCampoActivoDetalleEvaluacion_20200127_212423", m)
}

// Run the migrations
func (m *CambioTipoCampoActivoDetalleEvaluacion_20200127_212423) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP COLUMN activo")
	m.SQL("ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD COLUMN activo boolean NOT NULL")

}

// Reverse the migrations
func (m *CambioTipoCampoActivoDetalleEvaluacion_20200127_212423) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
