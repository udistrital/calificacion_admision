# Evaluacion_inscripcion_crud

API CRUD para la evaluación del proceso de inscripciónes y el manejo del entrevistador

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/evaluacion_inscripcion_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_PORT`: Puerto asignado para la ejecución del API
 - `EVALUACION_INSCRIPCION_CRUD__PGUSER`: Usuario de la base de datos
 - `EVALUACION_INSCRIPCION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `EVALUACION_INSCRIPCION_CRUD__PGURLS`: Host de conexión
 - `EVALUACION_INSCRIPCION_CRUD__PGDB`: Nombre de la base de datos
 - `EVALUACION_INSCRIPCION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_PORT=8083 EVALUACION_INSCRIPCION_CRUD__PGUSER=userPG EVALUACION_INSCRIPCION_CRUD__PGPASS=passwordPG EVALUACION_INSCRIPCION_CRUD__PGURLS=localhost EVALUACION_INSCRIPCION_CRUD__PGDB=academica EVALUACION_INSCRIPCION_CRUD__SCHEMA=evaluacion_inscripcion bee run

## MODELO DE DATOS

Como modelos de datos del Api se utilizo el siguiente [Modelo](https://drive.google.com/drive/folders/1fiz177bvO0WaxdV8QZz4iFnh94jofGvX?usp=sharing)


![image](./evaluacion_admision.png)