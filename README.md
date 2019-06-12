# Calificacion_admision_crud

API CRUD para la calificacion del proceso de admiciones y el manejo del entrevistador

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/calificacion_admision

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_CALIFICACION_ADMISION_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `CALIFICACION_ADMISION_CRUD__PGUSER`: Usuario de la base de datos
 - `CALIFICACION_ADMISION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `CALIFICACION_ADMISION_CRUD__PGURLS`: Host de conexión
 - `CALIFICACION_ADMISION_CRUD__PGDB`: Nombre de la base de datos
 - `CALIFICACION_ADMISION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_CALIFICACION_ADMISION_HTTP_PORT=8083 CALIFICACION_ADMISION_CRUD__PGUSER=userPG CALIFICACION_ADMISION_CRUD__PGPASS=passwordPG CALIFICACION_ADMISION_CRUD__PGURLS=localhost CALIFICACION_ADMISION_CRUD__PGDB=academica CALIFICACION_ADMISION_CRUD__SCHEMA=calificacion_admision bee run

## MODELO DE DATOS

Como modelos de datos del Api se utilizo el siguiente [Modelo](https://drive.google.com/drive/folders/1uIAI8LfF51TcbIommrKDmuFc0juQjobp)

![image](https://github.com/udistrital/calificacion_admision/blob/dev/academica.png) 

