-- Eliminar constraint fk_evaluacion_inscripcion_requisito_programa_academico
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion DROP CONSTRAINT fk_evaluacion_inscripcion_requisito_programa_academico;

-- Quitar NN del campo requisito_programa_academico_id para eleminar
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion  ALTER COLUMN requisito_programa_academico_id DROP NOT NULL;

-- Eliminar column entrevista_id y requisito_programa_academico_id
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion DROP COLUMN requisito_programa_academico_id;
ALTER TABLE evaluacion_inscripcion.evaluacion_inscripcion DROP COLUMN entrevista_id