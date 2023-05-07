
# descuento_academica_crud
El API provee la gestion de las diferentes descuentos que puede tener un tercero en el Sistema de Gestión Académica


Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `descuento_academico_crud master/develop`


### Variables de Entorno
```shell
DESCUENTO_ACADEMICO_CRUD_PGDB=[nombre de la base de datos]
DESCUENTO_ACADEMICO_CRUD_PGPASS=[password del usuario]
DESCUENTO_ACADEMICO_CRUD_PGHOST=[direccion de la base de datos]
DESCUENTO_ACADEMICO_CRUD_PGPORT=[puerto de conexión con la base de datos]
DESCUENTO_ACADEMICO_CRUD_PGUSER=[usuario con acceso a la base de datos]
DESCUENTO_ACADEMICO_CRUD_SCHEMA=[esquema donde se ubican las tablas]
DESCUENTO_ACADEMICO_CRUD_HTTP_PORT=[puerto de ejecucion] bee run
```

## Requerimientos
Go version >= 1.8.


## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/descuento_academico_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `DESCUENTO_ACADEMICO_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `DESCUENTO_ACADEMICO_CRUD__PGUSER`: Usuario de la base de datos
 - `DESCUENTO_ACADEMICO_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `DESCUENTO_ACADEMICO_CRUD__PGURLS`: Host de conexión
 - `DESCUENTO_ACADEMICO_CRUD__PGDB`: Nombre de la base de datos
 - `DESCUENTO_ACADEMICO__SCHEMA`: Esquema a utilizar en la base de datos


# 4. alimentar todas las variables de entorno que utiliza el proyecto.
DESCUENTO_ACADEMICO_CRUD_HTTP_PORT=8080 DESCUENTO_ACADEMICO_CRUD_PGHOST=127.0.0.1:27017 DESCUENTO_ACADEMICO_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=descuento_academica_crud . --no-cache
# docker run -p 80:80 descuento_academica_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/descuento_academica_crud

#2. Moverse a la carpeta del repositorio
cd descuento_academica_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master | Main |
| -- | -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/descuento_academico_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/descuento_academico_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/descuento_academico_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/descuento_academico_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/descuento_academico_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/descuento_academico_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/descuento_academico_crud/status.svg?ref=refs/heads/main)](https://hubci.portaloas.udistrital.edu.co/udistrital/descuento_academico_crud) |


## Modelo de Datos
[Modelo de Datos API CRUD Descuento Académico](https://github.com/planesticud/descuento_academico_crud/blob/develop/modelo_descuento_academico_crud.png)


## Licencia

This file is part of descuento_academica_crud.

descuento_academica_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

descuento_academica_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.

