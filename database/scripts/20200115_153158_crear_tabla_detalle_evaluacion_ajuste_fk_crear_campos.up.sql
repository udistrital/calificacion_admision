-- Creacion tabla tabla y CONSTRAINT
-- object: evaluacion_inscripcion.detalle_evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.detalle_evaluacion CASCADE;
CREATE TABLE evaluacion_inscripcion.detalle_evaluacion (
	id serial NOT NULL,
	evaluacion_inscripcion_id integer,
	requisito_programa_academico_id integer NOT NULL,
	nota_requisito numeric(5,2) NOT NULL,
	activo timestamp NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	entrevista_id integer,
	detalle_calificacion json,
	CONSTRAINT pk_detalle_evaluacion PRIMARY KEY (id)

);
-- object: fk_detalle_evaluacion_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_entrevista FOREIGN KEY (entrevista_id)
REFERENCES evaluacion_inscripcion.entrevista (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_detalle_evaluacion_evaluacion_inscripcion | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_evaluacion_inscripcion CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_evaluacion_inscripcion FOREIGN KEY (evaluacion_inscripcion_id)
REFERENCES evaluacion_inscripcion.evaluacion_inscripcion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


-- object: fk_detalle_evaluacion_requisito_programa_academico | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_requisito_programa_academico CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_requisito_programa_academico FOREIGN KEY (requisito_programa_academico_id)
REFERENCES evaluacion_inscripcion.requisito_programa_academico (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- elminar uq por cambio nombre 
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico DROP CONSTRAINT uq_criterio_programa_periodo;

-- object: uq_programa_requisito_periodo | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito_programa_academico DROP CONSTRAINT IF EXISTS uq_programa_requisito_periodo CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico ADD CONSTRAINT uq_programa_requisito_periodo UNIQUE (programa_academico_id,periodo_id,requisito_id);
-- ddl-end --

-- Crear campo faltante en requisito
-- Add column formato en requisito
ALTER TABLE evaluacion_inscripcion.requisito ADD COLUMN formato json;

-- ADD columun en requisito_programa_academico y cambiar nombre de porcentaje 
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico ADD COLUMN porcentaje_especifico json;
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico RENAME COLUMN porcentaje TO porcentaje_general;