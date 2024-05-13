# Project Structure
The project structure follows a standard layout, with different components organized into separate folders:

## routes
The routes folder contains the router configuration for handling incoming HTTP requests. This is where the routes for various endpoints are defined and mapped to their respective controller functions.

## controller
The controller folder contains the implementation of the controller functions corresponding to the routes defined in the router. These functions handle the business logic for processing incoming requests, interacting with the database, and returning responses.

## models
The models folder contains the data models or structures used to represent entities in the application, such as users, orders, menu items, etc. These models are used to interact with the database and marshal/unmarshal data.

## middleware
The middleware folder contains any middleware functions used in the application, such as authentication middleware.

## database
The database folder contains logic related to MongoDB integration and database operations. This includes functions for establishing connections to the MongoDB database, performing CRUD operations on collections, and handling data persistence.

## helpers
The helpers folder contains utility functions and modules related to JWT (JSON Web Token) authentication and token generation. These functions are responsible for generating JWT tokens, validating tokens, and handling user authentication and authorization.