-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 0.9.4
-- PostgreSQL version: 13.0
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: evaluacion_admision | type: DATABASE --
-- DROP DATABASE IF EXISTS evaluacion_admision;
CREATE DATABASE evaluacion_admision;
-- ddl-end --


-- object: evaluacion_inscripcion | type: SCHEMA --
-- DROP SCHEMA IF EXISTS evaluacion_inscripcion CASCADE;
CREATE SCHEMA evaluacion_inscripcion;
-- ddl-end --
COMMENT ON SCHEMA evaluacion_inscripcion IS E'Esquema para almacenar cada uno de los criterios de admision';
-- ddl-end --

SET search_path TO pg_catalog,public,evaluacion_inscripcion;
-- ddl-end --

-- object: evaluacion_inscripcion.requisito | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.requisito CASCADE;
CREATE TABLE evaluacion_inscripcion.requisito (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	asistencia boolean,
	examen_estado boolean,
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	requisito_padre_id integer,
	CONSTRAINT pk_requisito PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.requisito IS E'Tabla que almacena los diferentes criterios de admisión que se pueden aplicar';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.nombre IS E'Nombre del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.descripcion IS E'Descripción del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.asistencia IS E'Indica si el aspirante debe cumplir el requisito de manera presencial';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.examen_estado IS E'Indica si el requisito es un examen de estado';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.codigo_abreviacion IS E'Código de abreviación del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.activo IS E'Flag que indica si el criterio de admisión esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.numero_orden IS E'Número de orden en el que se deben mostrar los criterios de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.estado_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.estado_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.estado_entrevista (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_estado_entrevista PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.estado_entrevista IS E'Tabla que almacena los diferentes  estados que puede tener una entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.nombre IS E'Nombre del estado de la entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.descripcion IS E'Descripción del estado de entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.codigo_abreviacion IS E'Código de abreviación del estado de la entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.activo IS E'Flag que indica si el estado esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.numero_orden IS E'Número de orden en el que se deben mostrar los estados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevista (
	id serial NOT NULL,
	inscripcion_id integer NOT NULL,
	fecha_entrevista timestamp NOT NULL,
	estado_entrevista_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	tipo_entrevista_id integer NOT NULL,
	CONSTRAINT pk_entrevista PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.entrevista IS E'Tabla que almacena la información de las entrevistas realizadas a un aspirante';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.inscripcion_id IS E'Proceso de admisión al cual se hace referencia';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_entrevista IS E'Día y hora en la que se agenda la entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.activo IS E'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: fk_entrevista_estado_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevista DROP CONSTRAINT IF EXISTS fk_entrevista_estado_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevista ADD CONSTRAINT fk_entrevista_estado_entrevista FOREIGN KEY (estado_entrevista_id)
REFERENCES evaluacion_inscripcion.estado_entrevista (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.tipo_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.tipo_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.tipo_entrevista (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_tipo_entrevista PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.tipo_entrevista IS E'Tabla que almacena los diferentes  tipos de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.nombre IS E'Nombre del tipo de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.descripcion IS E'Descripción del tipo de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.codigo_abreviacion IS E'Código de abreviación del tipo de entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.activo IS E'Flag que indica si el tipo esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.numero_orden IS E'Número de orden en el que se deben mostrar los tipos';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: fk_entrevista_tipo_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevista DROP CONSTRAINT IF EXISTS fk_entrevista_tipo_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevista ADD CONSTRAINT fk_entrevista_tipo_entrevista FOREIGN KEY (tipo_entrevista_id)
REFERENCES evaluacion_inscripcion.tipo_entrevista (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.entrevistador | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevistador CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevistador (
	id serial NOT NULL,
	persona_id serial NOT NULL,
	programa_academico_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_entrevistador PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.entrevistador IS E'Persona disgnada por el proyecto para realizar las entrevistas';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.persona_id IS E'Entrevistador designado, anteriormente ente, se cambia por futura migracion de ente a persona';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.programa_academico_id IS E'Programa académico que designa el entrevistador';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.activo IS E'Flag que permite saber si el entrevistador se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.requisito_programa_academico | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico CASCADE;
CREATE TABLE evaluacion_inscripcion.requisito_programa_academico (
	id serial NOT NULL,
	programa_academico_id integer NOT NULL,
	periodo_id integer NOT NULL,
	porcentaje_general numeric(5,2) NOT NULL,
	porcentaje_especifico json,
	opcion json,
	puntaje_minimo_examen_estado integer,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	requisito_id integer NOT NULL,
	CONSTRAINT pk_requisito_programa_academico PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.requisito_programa_academico IS E'Tabla que almacena los diferentes criterios para cada programa academico';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.programa_academico_id IS E'Programa academico con el que se relaciona el criterio';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.periodo_id IS E'Periodo al cual aplica el criterio de inscripción';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.porcentaje_general IS E'Porcentaje o peso del criterio de admisión para el proyecto.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.activo IS E'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.entrevistador_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevistador_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevistador_entrevista (
	id serial NOT NULL,
	entrevistador_id integer NOT NULL,
	activo timestamp NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	nota_parcial numeric(5,2) NOT NULL,
	entrevista_id integer NOT NULL,
	CONSTRAINT pk_entrevistador_entrevista PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.entrevistador_entrevista IS E'Tabla que almacena los entrevistadores de una entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.activo IS E'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.nota_parcial IS E'Nota asignada por el entrevistador';
-- ddl-end --

-- object: fk_entrevistador_entrevista_entrevistador | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevistador_entrevista DROP CONSTRAINT IF EXISTS fk_entrevistador_entrevista_entrevistador CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevistador_entrevista ADD CONSTRAINT fk_entrevistador_entrevista_entrevistador FOREIGN KEY (entrevistador_id)
REFERENCES evaluacion_inscripcion.entrevistador (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_entrevistador_entrevista_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevistador_entrevista DROP CONSTRAINT IF EXISTS fk_entrevistador_entrevista_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevistador_entrevista ADD CONSTRAINT fk_entrevistador_entrevista_entrevista FOREIGN KEY (entrevista_id)
REFERENCES evaluacion_inscripcion.entrevista (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.cupos_por_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.cupos_por_dependencia CASCADE;
CREATE TABLE evaluacion_inscripcion.cupos_por_dependencia (
	id serial NOT NULL,
	dependencia_id integer NOT NULL,
	cupos_habilitados integer NOT NULL,
	cupos_opcionados integer NOT NULL,
	periodo_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_cupos_por_dependencia PRIMARY KEY (id),
	CONSTRAINT uq_dependencia_periodo UNIQUE (dependencia_id,periodo_id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.cupos_por_dependencia IS E'Tabla para el manejo de los cupos en cada dependencia';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.dependencia_id IS E'Campo que referencia al esquema de OIKOS';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.cupos_habilitados IS E'campo para el registro de numero de cupos habilitados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.cupos_opcionados IS E'campo para el registro del numero de cupos opcionados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.periodo_id IS E'Campo que referencia al esquema core';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.activo IS E'Campo para registrar si el registro se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.fecha_creacion IS E'Campo para el registro de la fecha de creación del registor';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.fecha_modificacion IS E'Campo para el registro de la fecha de modificación del registro';
-- ddl-end --
ALTER TABLE evaluacion_inscripcion.cupos_por_dependencia OWNER TO postgres;
-- ddl-end --

-- object: evaluacion_inscripcion.detalle_evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.detalle_evaluacion CASCADE;
CREATE TABLE evaluacion_inscripcion.detalle_evaluacion (
	id serial NOT NULL,
	nota_requisito numeric(5,2) NOT NULL,
	inscripcion_id integer NOT NULL,
	activo timestamp NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	detalle_calificacion json,
	entrevista_id integer,
	requisito_programa_academico_id integer NOT NULL,
	CONSTRAINT pk_detalle_evaluacion PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.detalle_evaluacion IS E'Tabla que almacena los entrevistadores de una entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.nota_requisito IS E'Nota parcial  del criterio sobre la evaluacion total';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.inscripcion_id IS E'Identificador de la inscripción a la que corresponde el detalle de la evaluación';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.activo IS E'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.detalle_evaluacion.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --

-- object: fk_detalle_evaluacion_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_entrevista FOREIGN KEY (entrevista_id)
REFERENCES evaluacion_inscripcion.entrevista (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_requisito_programa_academico_requisito | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito_programa_academico DROP CONSTRAINT IF EXISTS fk_requisito_programa_academico_requisito CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico ADD CONSTRAINT fk_requisito_programa_academico_requisito FOREIGN KEY (requisito_id)
REFERENCES evaluacion_inscripcion.requisito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_detalle_evaluacion_requisito_programa_academico | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.detalle_evaluacion DROP CONSTRAINT IF EXISTS fk_detalle_evaluacion_requisito_programa_academico CASCADE;
ALTER TABLE evaluacion_inscripcion.detalle_evaluacion ADD CONSTRAINT fk_detalle_evaluacion_requisito_programa_academico FOREIGN KEY (requisito_programa_academico_id)
REFERENCES evaluacion_inscripcion.requisito_programa_academico (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_programa_requisito_periodo | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito_programa_academico DROP CONSTRAINT IF EXISTS uq_programa_requisito_periodo CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico ADD CONSTRAINT uq_programa_requisito_periodo UNIQUE (programa_academico_id,periodo_id,requisito_id);
-- ddl-end --

-- object: fk_requisito | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito DROP CONSTRAINT IF EXISTS fk_requisito CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito ADD CONSTRAINT fk_requisito FOREIGN KEY (requisito_padre_id)
REFERENCES evaluacion_inscripcion.requisito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.tags_por_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.tags_por_dependencia CASCADE;
CREATE TABLE evaluacion_inscripcion.tags_por_dependencia (
	id serial NOT NULL,
	dependencia_id integer NOT NULL,
	periodo_id integer NOT NULL,
	lista_tags json NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_inscripcion_id integer NOT NULL,
	CONSTRAINT pk_tags_por_programa PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.tags_por_dependencia IS E'Tabla que almacena en campo json los tags de inscripción requeridos para cada programa en particular y en un periodo específico.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.dependencia_id IS E'Id programa academico al cual se relacionan los tags';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.periodo_id IS E'Id del periodo que se relaciona al programa academico';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.lista_tags IS E'Campo JSON que lista los tags requeridos para determinado programa';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.activo IS E'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tags_por_dependencia.tipo_inscripcion_id IS E'Id tipo inscripcion a la cual se relacionan los tags';
-- ddl-end --
ALTER TABLE evaluacion_inscripcion.tags_por_dependencia OWNER TO postgres;
-- ddl-end --


