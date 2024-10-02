# Shop-Dev-Ecommerce

## Init project
```go
go mod init
```

## Project Structure

Shop-Dev-Ecommerce/  
├── cmd/  
│   ├── cli/  
│   ├── cronjob/  
│   └── server/  
├── configs/  
├── docs/  
├── global/  
├── internal/  
│   ├── controller/  
│   ├── initialize/  
│   ├── middleware/  
│   ├── models/  
│   ├── repo/  
│   ├── router/  
│   └── service/  
├── migrations/  
│   └── 1_create_table.up.sql  
├── pkg/  
│   └── response/  
│       └── httpStatusCode.go  
├── scripts/  
├── tests/  
├── third_party/  
├── go.mod  
├── Makefile  
└── README.md



## Explanation of Folders

### **cmd**
This directory contains the entry points for your application. It includes the server, CLI commands, and cron jobs.

- **cli**: Handles command-line interface (CLI) commands.
- **cronjob**: Manages scheduled background jobs.
- **server**: The main server executable for running the application.

### **configs**
Contains configuration files and common constants used throughout the application. It centralizes settings like environment variables, API keys, and database configurations.

### **docs**
Contains documentation related to the project. This can include architecture documentation, API references, or other important guides.

### **global**
This folder contains global variables or settings that can be accessed throughout the project. It is useful for shared state or common configurations.

### **internal**
The `internal` folder contains core internal logic. The packages inside `internal` cannot be imported from outside the module, ensuring that your internal logic stays encapsulated.

- **controller**: Handles business logic for HTTP requests and routes.
- **initialize**: Initializes components like database connections, configurations, etc.
- **middleware**: Contains custom middleware functions (e.g., logging, authentication).
- **models**: Defines the application's data structures and database models.
- **repo**: Manages data access and interacts with the database.
- **router**: Handles the routing of incoming HTTP requests to the correct controller.
- **service**: Contains business logic and higher-level functions that may involve multiple controllers or models.

### **migrations**
Stores SQL migration files used to update the database schema. For example, `1_create_table.up.sql` likely contains a script to create initial database tables.

### **pkg**
The `pkg` folder contains reusable code that can be shared across different projects.

- **response**: A package for handling common HTTP responses, including status codes and error messages.

### **scripts**
This directory contains scripts that help automate tasks related to the project, such as deployment scripts or build commands.

### **tests**
Contains unit and integration tests for the project to ensure the correctness of the application code.

### **third_party**
Holds third-party libraries or dependencies that are not managed by the Go module system.

### **go.mod**
The Go modules file declares the module name and lists dependencies for the project.

### **Makefile**
The `Makefile` is used for automating common tasks, such as building the project, running tests, and deploying. It contains commands that can be run using the `make` command.

### **README.md**
This file provides documentation and instructions for the project, including how to set it up, run it, and contribute to it.


