# evaluacion_inscripcion_crud
API para la gestión de las calificaciones del proceso de inscripción y admisión

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `evaluacion_inscripcion_crud master/develop`

## Requirements
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/evaluacion_inscripcion_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `EVALUACION_INSCRIPCION_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `EVALUACION_INSCRIPCION_CRUD__PGUSER`: Usuario de la base de datos
 - `EVALUACION_INSCRIPCION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `EVALUACION_INSCRIPCION_CRUD__PGURLS`: Host de conexión
 - `EVALUACION_INSCRIPCION_CRUD__PGDB`: Nombre de la base de datos
 - `EVALUACION_INSCRIPCION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_PORT=8083 EVALUACION_INSCRIPCION_CRUD__PGUSER=userPG EVALUACION_INSCRIPCION_CRUD__PGPASS=passwordPG EVALUACION_INSCRIPCION_CRUD__PGURLS=localhost EVALUACION_INSCRIPCION_CRUD__PGDB=academica EVALUACION_INSCRIPCION_CRUD__SCHEMA=evaluacion_inscripcion bee run

## Ejemplo
EVALUACION_INSCRIPCION_HTTP_PORT=9016 EVALUACION_INSCRIPCION_CRUD__PGUSER=user EVALUACION_INSCRIPCION_CRUD__PGPASS=password EVALUACION_INSCRIPCION_CRUD__PGURLS=localhost EVALUACION_INSCRIPCION_CRUD__PGDB=bd EVALUACION_INSCRIPCION_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![modelo_evaluacion_inscripcion_crud](https://user-images.githubusercontent.com/14035745/66088107-8b0b8880-e53f-11e9-99e3-a972a986c817.png)
