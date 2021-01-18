
-- ALTER TABLE evaluacion_inscripcion.requisito DROP COLUMN requisito_id;
ALTER TABLE evaluacion_inscripcion.requisito ADD COLUMN requisito_id INTEGER NULL;

-- object: fk_requisito | type CONSTRAINT --
-- ALTER TABLE evaluacion_inscripcion.requisito DROP CONSTRAINT IF EXISTS fk_requisito CASCADE;
ALTER TABLE evaluacion_inscripcion.requisito ADD CONSTRAINT fk_requisito FOREIGN KEY (requisito_id)
REFERENCES evaluacion_inscripcion.requisito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE evaluacion_inscripcion.requisito DROP COLUMN formato;
