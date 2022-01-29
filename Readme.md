# CRUD GraphQl ent.io

Api CRUD products and categories.

Api user access:

| username | password |
| :------: | :------: |
| user1    | $user1*  |

## Requirement


## API
Consume Api use Graphql.

**Description**

- Port: 8182
- [Playground: /api/gql/playground](http://127.0.0.1:8182/api/gql/playground)
- [Api: /api/gql](http://127.0.0.1:8182/api/gql)

## Database
Database squeme mapping using [ent.io](https://entgo.io/)

**Tables**

- Users  

| id | username | name | password |
| -- | -------- | ---- | -------- |

- Categories

| id | name | description |
| -- | ---- | -------- |

- Products

| id | id_category | name | price | description | stock |
| -- | ----------- | ---- | ----- | ----------- | ----- |


## Deployment

## Develop