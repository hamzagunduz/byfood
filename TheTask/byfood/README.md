# Get Started
1) docker-compose up --build -d
2) migrate -path db/migration -database "postgresql://byfood_user:password@localhost:5432/byfood_db?sslmode=disable" up
3) Use .http file or Postman to test the API

# To Do
1) Write (getUserById), (updateUser), (deleteUser) functions
2) Separation of Concerns -- Organize the structure
3) Automize dummy data for postgres with Docker and migrate automatically
4) Complete React interface