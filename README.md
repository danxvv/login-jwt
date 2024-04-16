# login-jwt

## Respuesta a preguntas
El siguiente archivo contiene las respuestas a las preguntas planteadas en el archivo `RESPUESTA.md`.
[RESPUESTA.md](RESPUESTA.md)

## Description
Este proyecto es un servicio de registro y autenticación de usuarios utilizando JWT (JSON Web Tokens) en Go.

## Tecnologías
- Go
- JWT
- GORM
- Gin
- Docker

## Arquitectura
Este proyecto esta ligeramente basado en la arquitectura de **Clean Architecture**. La estructura de carpetas es la siguiente:

### /cmd
Este directorio contiene el punto de entrada de la aplicación. En este caso, el archivo `main.go` que inicializa la aplicación.

### /config
Este directorio contiene la configuración de la aplicación. En este caso, el archivo `config.go` que contiene las configuraciones de la base de datos y del servidor.

### /internal/app
Este directorio contiene la inicialización de la aplicación

### /internal/controllers
Este directorio contiene los controladores de la aplicación que manejan las peticiones HTTP asi como los middlewares.

### /internal/entity
Este directorio contiene las entidades de la aplicación. En este caso, la entidad `User`.

### /internal/usecase
Este directorio contiene los casos de uso de la aplicación. En este caso, los casos de uso para el registro y login de usuarios.

### /internal/usercase/repository
Este directorio contiene las capas de acceso a datos de la aplicación. En este caso, la capa de acceso a datos para la entidad `User`.

### /pkg
Este directorio contiene paquetes que pueden ser utilizados por toda la aplicación. 


## Endpoints
- POST /api/v1/register
  - Este endpoint permite registrar un usuario.
    - Request Body:
      ```json
        {
        "username":"user",
        "password":"ValidPass.1",
        "email":"valid@email.com",
        "phone":"1234567890"
        }
      ```
    - Response:
      ```json
        {
        "message": "User registered successfully"
        }
      ```
- POST /api/v1/login
    - Este endpoint permite autenticar un usuario.
      - Request Body:
        ```json
          {
          "username":"userOrEmail",
          "password":"ValidPass.1"
          }
        ```
      - Response:
        ```json
          {
        "data": {
          "username": "user",
          "email": "valid@email.com",
          "phone": "1234567890",
          "id": "f27f71dc-0c94-4ecb-a34a-f5b0a8d21029",
          "created_at": "2024-04-16 21:30:32.48785329 +0000 UTC",
          "updated_at": "2024-04-16 21:30:32.487853347 +0000 UTC"
         },
        "message": "User logged in successfully",
        "token": "jwt_token"
         }
        ```
- GET /api/v1/user
    - Este endpoint permite obtener la información de un usuario autenticado.
      - Request Header:
        ```
        Authorization : Bearer jwt_token   
        ```
      - Response:
        ```json
        {
        "data": {
        "username": "user",
        "email": "valid@email.com",
        "phone": "1234567890",
        "id": "f27f71dc-0c94-4ecb-a34a-f5b0a8d21029",
        "created_at": "2024-04-16 21:30:32.48785329 +0000 UTC",
        "updated_at": "2024-04-16 21:30:32.487853347 +0000 UTC"
        },
        "message": "User data"
        }
        

## Instalación
1. Clonar el repositorio
    ```bash
    git clone
    ```
2. Crear un archivo `.env` en la raíz del proyecto con las siguientes variables de entorno:
    ```bash
   cp .env.example .env
    ```
3. Ejecutar el siguiente comando para construir la imagen de Docker:
    ```bash
    docker build -t login-jwt .
    ```

4. Ejecutar el siguiente comando para levantar el contenedor de Docker:
    ```bash
   docker run --rm -it --env-file=.env -v $(pwd):/usr/src/app -p 8080:8080 --name login-jwt-server login-jwt
    ```
5. La aplicación estará disponible en `http://localhost:8080`


## Live Demo
Puedes usar el siguiente enlace para probar el API:
- https://login-service.deploys.danielmojica.me