CREATE TABLE evaluacion_inscripcion.tags_por_dependencia
(
    id serial NOT NULL,
    dependencia_id integer NOT NULL,
    periodo_id integer NOT NULL,
    lista_tags json NOT NULL,
    activo boolean NOT NULL,
    fecha_creacion timestamp NOT NULL DEFAULT NOW(),
    fecha_modificacion timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_tags_por_programa PRIMARY KEY (id)
);

COMMENT ON TABLE evaluacion_inscripcion.tags_por_dependencia
    IS 'Tabla que almacena en campo json los tags de inscripción requeridos para cada programa en particular y en un periodo específico.';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.id
    IS 'Identificador de la tabla';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.dependencia_id
    IS 'Id programa academico al cual se relacionan los tags';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.periodo_id
    IS 'Id del periodo que se relaciona al programa academico';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.lista_tags
    IS 'Campo JSON que lista los tags requeridos para determinado programa';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.activo
    IS 'Indica el estado del registro';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.fecha_creacion
    IS 'Fecha de creación del registro';

COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.fecha_modificacion
    IS 'Fecha de la última modificación del registro';