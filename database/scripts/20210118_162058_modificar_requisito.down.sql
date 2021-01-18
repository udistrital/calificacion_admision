ALTER TABLE evaluacion_inscripcion.requisito DROP CONSTRAINT IF EXISTS fk_requisito CASCADE;

ALTER TABLE evaluacion_inscripcion.requisito DROP COLUMN requisito_id;

ALTER TABLE evaluacion_inscripcion.requisito_id ADD COLUMN formato JSON NULL;
