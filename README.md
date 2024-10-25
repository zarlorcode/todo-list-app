# To-Do List Web Application

Este proyecto es una aplicación web de To-Do List (lista de tareas) construida con Go y Gin Gonic, diseñada para ofrecer una API REST que permite realizar operaciones CRUD sobre tareas. El proyecto está containerizado con Docker y utiliza Docker Compose para orquestar un entorno multi-contenedor con un servidor de aplicación y una base de datos.

## Tabla de Contenidos
- [Descripción](#descripción)
- [Requisitos Previos](#requisitos-previos)
- [Instalación](#instalación)
- [Uso](#uso)
- [Endpoints de la API](#endpoints-de-la-api)
- [Tecnologías Utilizadas](#tecnologías-utilizadas)
- [Licencia](#licencia)

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
1. Crear una nueva tarea
Método: POST
URL: /todos
Ejemplo de uso: curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"id": "2", "title": "Nueva tarea", "status": "pending"}'

2. Obtener todas las tareas
Método: GET
URL: /todos
Ejemplo de uso: curl -X GET http://localhost:8080/todos

3. Actualizar una tarea existente
Método: PUT
URL: /todos/:id (reemplaza :id con el ID de la tarea que deseas actualizar)
Ejemplo de uso: curl -X PUT http://localhost:8080/todos/2 -H "Content-Type: application/json" -d '{"title": "Tarea actualizada", "status": "completed"}'

4. Eliminar una tarea
Método: DELETE
URL: /todos/:id (reemplaza :id con el ID de la tarea que deseas eliminar)
Ejemplo de uso: curl -X DELETE http://localhost:8080/todos/2


## Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/tu_usuario/tu_repositorio.git
   cd tu_repositorio
   

