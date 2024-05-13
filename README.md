# Coffee Xpress App Backend
The Coffee Xpress App Backend is the server-side implementation for the Coffee Xpress application. This project provides the backend infrastructure necessary to support the functionality of the Coffee Xpress app, which allows users to browse and order coffee and other beverages from a menu, track their orders, and manage their accounts.

## Overview
The Coffee Xpress app backend is built using Golang (Go) programming language and MongoDB database. It serves as the middleware between the frontend client and the database, handling user authentication, order management, menu management, and other essential operations.

### Project Structure
The initial part implemented focuses on designing the project structure and setting up the basic routing for the Coffee Xpress app backend.

For more detailed documentation, see the [Project Structure](./Readme/ProjectStructure.md).

### Main.go
The main function serves as the core orchestrator of the Coffee Xpress backend application, responsible for configuring the server, defining routes, and handling incoming HTTP requests. It reads the port number from the environment variable or defaults to port 8000, initializes a Gin router with logging middleware, and sets up routes for user management, product handling, menu management, favorites, order history, cart items, and payment processing. Authentication middleware ensures secure access to protected routes. Finally, the server starts listening on the specified port, ready to handle requests from the Coffee Xpress frontend client.

### Database
The database package establishes a connection to the MongoDB database for the Coffee Xpress backend application. 

For more detailed documentation on the logic, see the [Database](./Readme/Database.md).

### Routes
The routes folder contains the router configuration for handling incoming HTTP requests. This is where the routes for various endpoints are defined and mapped to their respective controller functions.

### Models
This repository contains the MongoDB models for the Coffee Xpress backend. These models represent the various entities and relationships within the Coffee Xpress application.

For more detailed documentation on the logic, see the [Models](./Readme/Models.md).