# Books REST

A simple repository to store a multitude of rest APIs written in many languages and frameworks using the same ideia and database schema 

## Ideia

The REST API is simple, just a book recording system with the following operatios:
- Authors CRUD
- Books CRUD
- Users CRUD

## Schema

- **Authors**
  - id (pk)
  - first_name
  - last_name

- **Books**
  - id (pk)
  - isbn (unique)
  - name
  - author_id (fk -> authors)
  - creator_id (fk -> users)

- **Users**
  - id (pk)
  - first_name
  - last_name
  - username (unique)
  - password
  
