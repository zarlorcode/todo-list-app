# To-Do List Web Application

Este proyecto es una aplicación web de To-Do List (lista de tareas) construida con Go y Gin Gonic, diseñada para ofrecer una API REST que permite realizar operaciones CRUD sobre tareas. El proyecto está containerizado con Docker y utiliza Docker Compose para orquestar un entorno multi-contenedor con un servidor de aplicación y una base de datos.

## Tabla de Contenidos
- [Descripción](#descripción)
- [Requisitos Previos](#requisitos-previos)
- [Uso](#uso)
- [Descripción de componentes](#descripción de componentes)

## Descripción
Este proyecto implementa una API REST que permite:
- Crear una tarea
- Leer (listar) todas las tareas
- Actualizar una tarea existente
- Eliminar una tarea

La aplicación es útil para explorar conceptos básicos de desarrollo de APIs, containerización y orquestación de servicios con Docker.

## Requisitos Previos
- **Docker**: Asegúrate de tener Docker instalado. Puedes instalarlo siguiendo [estas instrucciones](https://docs.docker.com/get-docker/).
- **Docker Compose**: Asegúrate de tener Docker Compose instalado, que normalmente viene junto con Docker Desktop.

## Uso
1. Lanzar a ejecución:
    ```bash
    docker-compose up --build
    
A partir de ahora podemos realizar todas las operaciones, crear una nueva tarea, actualizar una tarea, ver la lista de tareas y borrar una tarea. Todo ello se puede hacer desde otra terminal(diferente a la del docker-compose up --build). Prueba a ejecutar los siguientes comandos:
    
2. Crear una nueva tarea
Método: POST
URL: /todos
Ejemplo de uso: 
    ```bash
    curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"id": "1", "title": "First Task", "status": "pending"}'

3. Obtener todas las tareas
Método: GET
URL: /todos
Ejemplo de uso: 
    ```bash
    curl -X GET http://localhost:8080/todos

4. Actualizar una tarea existente
Método: PUT
URL: /todos/:id (reemplaza :id con el ID de la tarea que deseas actualizar)
Ejemplo de uso: 
    ```bash
    curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title": "First Task actualizada", "status": "completed"}'

5. Eliminar una tarea
Método: DELETE
URL: /todos/:id (reemplaza :id con el ID de la tarea que deseas eliminar)
Ejemplo de uso: 
    ```bash
    curl -X DELETE http://localhost:8080/todos/1
    
6. Parar la ejecución:
    ```bash
    docker-compose down
    
## Descripción de componentes

**Dockerfile:**

El Dockerfile es un archivo de configuración que define cómo se debe construir la imagen del contenedor para la aplicación Go. Utiliza una construcción de múltiples etapas para optimizar el tamaño final de la imagen, comenzando con una imagen base ligera de golang:alpine para compilar el código fuente, seguido de una segunda etapa que utiliza alpine:latest para crear una imagen más pequeña que contiene solo el binario compilado y los archivos necesarios para ejecutar la aplicación.

**docker-compose.yml:**

Este archivo de configuración es utilizado por Docker Compose para definir y orquestar múltiples contenedores que componen la aplicación. En este caso, define un servicio para la aplicación Go y otro para una base de datos PostgreSQL, especificando las dependencias, puertos expuestos y variables de entorno necesarias para el funcionamiento de ambos contenedores, lo que simplifica su despliegue y gestión.

**main.go:** 
Este archivo contiene el código fuente principal de la aplicación, escrito en Go. Utiliza el framework Gin para crear un servidor web que maneja las solicitudes HTTP y proporciona una API RESTful para realizar operaciones CRUD en una lista de tareas. También incluye la lógica para conectar y migrar la base de datos PostgreSQL, gestionando las interacciones entre la API y la base de datos.

**.env:**
El archivo .env es utilizado para almacenar variables de entorno que configuran la aplicación y el contenedor de la base de datos. Contiene información sensible, como el nombre de usuario, la contraseña y el nombre de la base de datos, que son utilizados por la aplicación Go para establecer la conexión con PostgreSQL, permitiendo que la configuración sea fácilmente ajustable sin modificar el código fuente.

**go.mod:** 
Este archivo es utilizado por el sistema de gestión de módulos de Go para definir las dependencias del proyecto. Especifica el módulo del proyecto y las versiones de las bibliotecas requeridas, asegurando que el entorno de desarrollo tenga acceso a las versiones correctas de las dependencias utilizadas en el código.

**go.sum:** 
El archivo go.sum complementa al go.mod al registrar las versiones exactas de las dependencias y sus hash de verificación. Esto garantiza la integridad y consistencia del entorno de desarrollo y producción, permitiendo que otros desarrolladores y sistemas recuperen las mismas versiones de las bibliotecas, evitando problemas de compatibilidad.


   

