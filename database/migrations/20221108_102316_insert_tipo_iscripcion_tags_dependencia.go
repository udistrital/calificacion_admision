package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTipoIscripcionTagsDependencia_20221108_102316 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoIscripcionTagsDependencia_20221108_102316{}
	m.Created = "20221108_102316"

	migration.Register("InsertTipoIscripcionTagsDependencia_20221108_102316", m)
}

// Run the migrations
func (m *InsertTipoIscripcionTagsDependencia_20221108_102316) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.tags_por_dependencia ADD COLUMN tipo_inscripcion_id integer NOT NULL;")
	m.SQL("COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.tipo_inscripcion_id IS 'Id tipo inscripcion a la cual se relacionan los tags';")
}

// Reverse the migrations
func (m *InsertTipoIscripcionTagsDependencia_20221108_102316) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE IF EXISTS evaluacion_inscripcion.tags_por_dependencia DROP COLUMN IF EXISTS tipo_inscripcion_id;")
}
