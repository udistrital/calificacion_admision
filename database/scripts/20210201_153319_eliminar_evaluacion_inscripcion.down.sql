-- object: evaluacion_inscripcion.evaluacion_inscripcion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.evaluacion_inscripcion CASCADE;
CREATE TABLE evaluacion_inscripcion.evaluacion_inscripcion(
	id serial NOT NULL,
	inscripcion_id integer NOT NULL,
	nota_final numeric(5,2) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_evaluacion_inscripcion PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.evaluacion_inscripcion IS 'Tabla que almacena los puntajes de cada criterio de admisión de un proceso de admisión ';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.inscripcion_id IS 'Identificador del proceso de admisión al cual se hace referencia';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.nota_final IS 'Nota de cada criterio de evaluación';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.evaluacion_inscripcion.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

ALTER TABLE evaluacion_inscripcion.detalle_evaluacion RENAME COLUMN inscripcion_id TO evaluacion_inscripcion_id;

-- object: fk_detalle_evaluacion_evaluacion_inscripcion | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_evaluacion_inscripcion CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_evaluacion_inscripcion FOREIGN KEY (evaluacion_inscripcion_id)
REFERENCES evaluacion_inscripcion.evaluacion_inscripcion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --
