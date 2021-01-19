ALTER TABLE evaluacion_inscripcion.requisito DROP CONSTRAINT IF EXISTS fk_requisito CASCADE;

ALTER TABLE evaluacion_inscripcion.requisito DROP COLUMN requisito_padre_id;

ALTER TABLE evaluacion_inscripcion.requisito ADD COLUMN formato JSON NULL;
