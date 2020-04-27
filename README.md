# todo
- Go project to support a simple todo api. Technologies used: golang, postgresql, docker, terraform

- To build and run the project run docker-compose up command(almost works)

- To run database run `docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres`

- To build and run the todo app run `go build && DB_HOST=localhost ./todo`
    - Endpoints: 
        - GET /v1/todos params: optional -> status=[active, completed] to filter by status if needed or get all - List TODOs
        - POST /v1/todos - Create a new TODO
        - PATCH /v1/todos/:id/complete - Complete a TODO
        - PUT /v1/todos/:id - Update a TODO
        - DELETE /v1/todos/:id - Delete a TODO
    - Todo POJO: 
        - {
            name: string, <- required for POST
            id: int64,
            description: string, <- required for POST
            status: string,
            createDt: Time,
            updateDt: Time,
            isDeleted: bool
        }

- Future plans:
    - deploy and host in AWS using terraform(codeCommit, ECR, ECS)
    - make docker compose work, currently cannot connect to db for some reason, having issues authenticating
    - add swagger2 for self documentation
    - introduce pagination for GET method

