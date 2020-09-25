-- object: evaluacion_inscripcion | type: SCHEMA --
DROP SCHEMA IF EXISTS evaluacion_inscripcion CASCADE;
CREATE SCHEMA evaluacion_inscripcion;
-- ddl-end --
COMMENT ON SCHEMA evaluacion_inscripcion IS 'Esquema para almacenar cada uno de los criterios de admision';
-- ddl-end --

SET search_path TO pg_catalog,public,evaluacion_inscripcion;
-- ddl-end --

-- object: evaluacion_inscripcion.requisito | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.requisito CASCADE;
CREATE TABLE evaluacion_inscripcion.requisito(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_requisito PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.requisito IS 'Tabla que almacena los diferentes criterios de admisión que se pueden aplicar';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.nombre IS 'Nombre del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.descripcion IS 'Descripción del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.codigo_abreviacion IS 'Código de abreviación del criterio de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.activo IS 'Flag que indica si el criterio de admisiÛn esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.numero_orden IS 'N˙mero de orden en el que se deben mostrar los criterios de admisión';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.estado_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.estado_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.estado_entrevista(
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
COMMENT ON TABLE evaluacion_inscripcion.estado_entrevista IS 'Tabla que almacena los diferentes estados que puede tener una entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.nombre IS 'Nombre del estado de la entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.descripcion IS 'Descripción del estado de entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.codigo_abreviacion IS 'Código de abreviación del estado de la entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.activo IS 'Flag que indica si el estado esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.numero_orden IS 'Número de orden en el que se deben mostrar los estados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.estado_entrevista.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevista(
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
COMMENT ON TABLE evaluacion_inscripcion.entrevista IS 'Tabla que almacena la información de las entrevistas realizadas a un aspirante';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.inscripcion_id IS 'Proceso de admisión al cual se hace referencia';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_entrevista IS 'Día y hora en la que se agenda la entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevista.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --

-- object: fk_entrevista_estado_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevista DROP CONSTRAINT IF EXISTS fk_entrevista_estado_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevista ADD CONSTRAINT fk_entrevista_estado_entrevista FOREIGN KEY (estado_entrevista_id)
REFERENCES evaluacion_inscripcion.estado_entrevista (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.tipo_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.tipo_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.tipo_entrevista(
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
COMMENT ON TABLE evaluacion_inscripcion.tipo_entrevista IS 'Tabla que almacena los diferentes  tipos de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.nombre IS 'Nombre del tipo de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.descripcion IS 'DescripciÛn del tipo de entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.codigo_abreviacion IS 'Código de abreviación del tipo de entrevista.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.activo IS 'Flag que indica si el tipo esta activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.numero_orden IS 'Número de orden en el que se deben mostrar los tipos';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.tipo_entrevista.fecha_modificacion IS 'Fecha de la ˙ltima modificación del registro';
-- ddl-end --

-- object: fk_entrevista_tipo_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.entrevista DROP CONSTRAINT IF EXISTS fk_entrevista_tipo_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.entrevista ADD CONSTRAINT fk_entrevista_tipo_entrevista FOREIGN KEY (tipo_entrevista_id)
REFERENCES evaluacion_inscripcion.tipo_entrevista (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.entrevistador | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevistador CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevistador(
	id serial NOT NULL,
	persona_id serial NOT NULL,
	programa_academico_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_entrevistador PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.entrevistador IS 'Persona disgnada por el proyecto para realizar las entrevistas';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.persona_id IS 'Entrevistador designado, anteriormente ente, se cambia por futura migracion de ente a persona';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.programa_academico_id IS 'Programa académico que designa el entrevistador';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.activo IS 'Flag que permite saber si el entrevistador se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador.fecha_modificacion IS 'Fecha de la ˙ltima modificación del registro';
-- ddl-end --

-- object: evaluacion_inscripcion.requisito_programa_academico | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.requisito_programa_academico CASCADE;
CREATE TABLE evaluacion_inscripcion.requisito_programa_academico(
	id serial NOT NULL,
	programa_academico_id integer NOT NULL,
	periodo_id integer NOT NULL,
	requisito_id integer NOT NULL,
	porcentaje numeric(5,2) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_requisito_programa_academico PRIMARY KEY (id),
	CONSTRAINT uq_criterio_programa_periodo UNIQUE (programa_academico_id,periodo_id,requisito_id)

);
-- ddl-end --
COMMENT ON TABLE evaluacion_inscripcion.requisito_programa_academico IS 'Tabla que almacena los diferentes criterios para cada programa academico';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.programa_academico_id IS 'Programa academico con el que se relaciona el criterio';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.periodo_id IS 'Periodo al cual aplica el criterio de inscripción';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.porcentaje IS 'Porcentaje o peso del criterio de admisiÛn para el proyecto.';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.requisito_programa_academico.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT uq_criterio_programa_periodo ON evaluacion_inscripcion.requisito_programa_academico  IS 'Unique key que indica que un criterio solo puede aplicarse una vez en un programa durante un periodo.';
-- ddl-end --

-- object: fk_requisito_programa_academico_requisito | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito_programa_academico DROP CONSTRAINT IF EXISTS fk_requisito_programa_academico_requisito CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito_programa_academico ADD CONSTRAINT fk_requisito_programa_academico_requisito FOREIGN KEY (requisito_id)
REFERENCES evaluacion_inscripcion.requisito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.evaluacion_inscripcion | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.evaluacion_inscripcion CASCADE;
CREATE TABLE evaluacion_inscripcion.evaluacion_inscripcion(
	id serial NOT NULL,
	inscripcion_id integer NOT NULL,
	nota_final numeric(5,2) NOT NULL,
	requisito_programa_academico_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	entrevista_id integer,
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

-- object: fk_evaluacion_inscripcion_requisito_programa_academico | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion DROP CONSTRAINT IF EXISTS fk_evaluacion_inscripcion_requisito_programa_academico CASCADE;
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion ADD CONSTRAINT fk_evaluacion_inscripcion_requisito_programa_academico FOREIGN KEY (requisito_programa_academico_id)
REFERENCES evaluacion_inscripcion.requisito_programa_academico (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.entrevistador_entrevista | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.entrevistador_entrevista CASCADE;
CREATE TABLE evaluacion_inscripcion.entrevistador_entrevista(
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
COMMENT ON TABLE evaluacion_inscripcion.entrevistador_entrevista IS 'Tabla que almacena los entrevistadores de una entrevista';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.activo IS 'Indica el estado del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.entrevistador_entrevista.nota_parcial IS 'Nota asignada por el entrevistador';
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

-- object: fk_evaluacion_inscripcion_entrevista | type: CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion DROP CONSTRAINT IF EXISTS fk_evaluacion_inscripcion_entrevista CASCADE;
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion ADD CONSTRAINT fk_evaluacion_inscripcion_entrevista FOREIGN KEY (entrevista_id)
REFERENCES evaluacion_inscripcion.entrevista (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: evaluacion_inscripcion.cupos_por_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS evaluacion_inscripcion.cupos_por_dependencia CASCADE;
CREATE TABLE evaluacion_inscripcion.cupos_por_dependencia(
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
COMMENT ON TABLE evaluacion_inscripcion.cupos_por_dependencia IS 'Tabla para el manejo de los cupos en cada dependencia';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.dependencia_id IS 'Campo que referencia al esquema de OIKOS';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.cupos_habilitados IS 'campo para el registro de numero de cupos habilitados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.cupos_opcionados IS 'campo para el registro del numero de cupos opcionados';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.periodo_id IS 'Campo que referencia al esquema core';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.activo IS 'Campo para registrar si el registro se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.fecha_creacion IS 'Campo para el registro de la fecha de creaciÛn del registor';
-- ddl-end --
COMMENT ON COLUMN evaluacion_inscripcion.cupos_por_dependencia.fecha_modificacion IS 'Campo para el registro de la fecha de modificación del registro ';
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA evaluacion_inscripcion TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA evaluacion_inscripcion TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA evaluacion_inscripcion TO desarrollooas;
