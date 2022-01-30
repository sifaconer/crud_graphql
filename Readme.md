# CRUD GraphQl

Api CRUD products y categories.

Usuario por defecto:

| username | password |
| :------: | :------: |
| user1    | $user1*  |

## API

Consume del api usando Graphql Playground.

**Description**

- Port: 8182

- [Playground: /gql/playground](http://127.0.0.1:8182/gql/playground)

- [Api: /gql/api](http://127.0.0.1:8182/gql/api)

## Database

- Database [Postgres v13](https://www.postgresql.org/)

- ORM [ent.io](https://entgo.io/)

**Tablas**

- Users  

| Col. Name | Type     |
| --------- | -------- |
| id        | int      |
| username  | string   |
| name      | string   |
| password  | string   |


- Categories

| Col. Name | Type     |
| --------- | -------- |
| id        | int      |
| name      | string   |
| description  | string   |

- Products

| Col. Name | Type     |
| --------- | -------- |
| id | int |
| id_category | int |
| name | string |
| price | float |
| description | string |
| stock | int |


## Despliegue

Iniciar el servidor.

**Variable de entorno**

POSTGRES_URI="host=127.0.0.1 port=5432 user=root dbname=postgres password=root sslmode=disable"

**Docker Compose**

Ajuste los parámetros de postgres en la variable de entorno URI_DATABASE dentro del archivo docker-compose.yml

Build and Run
: docker-compose up -d

**Docker**

Build
: docker build -t api:v1 .

Run
: docker run -e URI_DATABASE='{POSTGRES_URI}' -p 8182:8182 api:v1

**Desarrollo**

Install Task: [Task](https://taskfile.dev/#/)

Run
: task run

Compile
: task compile