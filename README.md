# To-Do List Web Application

Este proyecto es una aplicación web de To-Do List (lista de tareas) construida con Go y Gin Gonic, diseñada para ofrecer una API REST que permite realizar operaciones CRUD sobre tareas. El proyecto está containerizado con Docker y utiliza Docker Compose para orquestar un entorno multi-contenedor con un servidor de aplicación y una base de datos.

## Tabla de Contenidos
- [Descripción](#descripción)
- [Requisitos Previos](#requisitos-previos)
- [Uso](#uso)

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
    curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"id": "3", "title": "Third Task", "status": "pending"}'

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
    curl -X PUT http://localhost:8080/todos/3 -H "Content-Type: application/json" -d '{"title": "Tarea 3 actualizada", "status": "completed"}'

5. Eliminar una tarea
Método: DELETE
URL: /todos/:id (reemplaza :id con el ID de la tarea que deseas eliminar)
Ejemplo de uso: 
    ```bash
    curl -X DELETE http://localhost:8080/todos/3
    
6. Parar la ejecución:
    ```bash
    docker-compose down


   

