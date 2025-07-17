# ⚡ Workflow Editor

A modern workflow editor app for designing and executing custom automation workflows (e.g., weather notifications). Users can visually build workflows, configure parameters, and view real-time execution results.

## 🛠️ Tech Stack

- **Frontend:** React + TypeScript, @xyflow/react (drag-and-drop), Radix UI, Tailwind CSS, Vite
- **Backend:** Go API, PostgreSQL database
- **DevOps:** Docker Compose for orchestration, hot reloading for rapid development

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose (recommended for development)
- Node.js v18+ (for local frontend development)
- Go v1.23+ (for local backend development)

> **Tip:** Node.js and Go are only required if you want to run frontend or backend outside Docker.

### 1. Start All Services

```bash
docker-compose up --build
```

- This launches frontend, backend, and database with hot reloading enabled for code changes.
- To stop and clean up:
  ```bash
  docker-compose down
  ```

### 2. Access Applications

- **Frontend (Workflow Editor):** [http://localhost:3003](http://localhost:3003)
- **Backend API:** [http://localhost:8086](http://localhost:8086)
- **Database:** PostgreSQL on `localhost:5876`

### 3. Verify Setup

1. Open [http://localhost:3003](http://localhost:3003) in your browser.
2. You should see the workflow editor with sample nodes.

## 🏗️ Project Architecture

```text
workflow-code-test/
├── api/                    # Go Backend (Port 8086)
│   ├── main.go
│   ├── services/
│   ├── pkg/
│   ├── go.mod
│   └── Dockerfile
├── web/                    # React Frontend (Port 3003)
│   ├── src/
│   ├── public/
│   ├── package.json
│   ├── vite.config.ts
│   ├── tsconfig.json
│   ├── nginx.conf
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🔧 Development Workflow

### 🌐 Frontend

- Edit files in `web/src/` and see changes instantly at [http://localhost:3003](http://localhost:3003) (hot reloading via Vite).

### 🖥️ Backend

- Edit files in `api/` and changes are reflected automatically (hot reloading in Docker).
- If you add new dependencies or make significant changes, rebuild the API container:
  ```bash
  docker-compose up --build api
  ```

### 🗄️ Database

- Schema/configuration details: see [API README](api/README.md#database)
- After schema changes or migrations, restart the database:
  ```bash
  docker-compose restart postgres
  ```
- To apply schema changes to the API after updating the database:
  ```bash
  docker-compose restart api
  ```
  
## Trade-offs and Analysis
### Trade-offs
- Selection of JSONB as the storage data-type in the PostgreSQL database for future scalability.
  - JSONB has a larger memory foot-print but enables for quicker node/edge queries and updates given the requirement for a single table.
- Handlers separated into 'get_workflow_handler.go' and 'execute_workflow_handler.go' to reduce code complexity.
- Enable 'execute_node.go' in the workflow engine to handle base metadata to allow easier extension of future nodes.
  - Additional node-specific is computed in each separate node and added to the returned ExecutionStep.
  - Output variables of each node are stored since some nodes utilise variables from more than just the previous node.
- Separation of validator to enable scaling and increase in complexity - this could include detection of cyclic workflow graphs etc.
  - Only explicitly required validation has been added thus far which includes:
    - Start and End nodes must respectively not be the target and source node of an edge.
    - Node ids must be unique since they are being used by edges to identify source and target.
    - Edge ids must be unique as a method for identifying a given edge.
- Separation of store (repository-layer) from handler to reduce function complexity and ensure code modularity.
  - Additionally enables easier testing for integration tests with DockerTest (discussed further below).
- Decision to use Protobuf/Proto-json as opposed to GoLang struct for marshalling/unmarshalling based on performance.
  - Protobuf ensures solid schema and type-checking on the backend (to match frontend type definitions). Ideally they would be defined in the same location.

### Packages and Libraries
- 'GoLang Migrate' - Utilised for up/down database migrations and is mainly used for the schema (DDL) definition.
- 'DockerTest' - Enables the spinning up of short-lived docker containers for integration tests on the store (repository-layer). These containers can therefore also be used for testing within CI/CD pipelines.
- 'Protobuf/Proto-json' - Define the backend types and perform high-level validation on input when unmarshalling from request. Additionally, protojson has improved marshal and unmarshal functions as compared to the native GoLang libraries.
- 'Testify' - Standard GoLang package used for handling testing - specific usage is with DockerTest whereby Testify's TestMain function is able to spin up the docker container, and ensure that the database has been setup and migrated successfully.
- 'UUID' (Google) - Given that the initial example provides a UUIDv4 workflow ID, and PostgreSQL natively handles UUID data-type, Google's UUID enable for quick generation for new workflows and can also be used for generating execution ids in workflows.

### Assumptions
- A node with type 'email' will follow a 'condition' node.
- Nodes can have at most two edges, and cannot follow both edges simultaneously (e.g. lead to the execution of two nodes in parallel).
- The validation of a 'form' node requires the options from an 'integration' node in order to be valid.
- Backend JSON stub returned "success" and frontend type definition specified 'success' or 'error. However, utilising either of these string results in an incorrect UI display - therefore the backend definition has been changed from 'success' -> 'completed' and 'error' -> 'failed'.
- Expected duration and total duration should be milliseconds.
- Temperature and threshold should be accurate to a maximum of 2 decimal places.
- As per specification, the executed workflow must first persist the workflow configuration and separately retrieve it.

### Addition of New Node Types
New node types can be introduced by:
- Adding an expected result to the Execution Step output.
- Adding the implementation to 'execute_node.go' and creating a new file for node-specific execution.
- Adding required validations to the 'validator' package.

### Execute Workflow Sequence Diagram
![execute_workflow_sd.png](https://private-user-images.githubusercontent.com/221332256/467558785-7a095b50-e970-44bf-b5fd-54efdeabcfa4.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NTI3NTgwNjcsIm5iZiI6MTc1Mjc1Nzc2NywicGF0aCI6Ii8yMjEzMzIyNTYvNDY3NTU4Nzg1LTdhMDk1YjUwLWU5NzAtNDRiZi1iNWZkLTU0ZWZkZWFiY2ZhNC5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjUwNzE3JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI1MDcxN1QxMzA5MjdaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT00ODdiMjg2NTUyNmZlZGY3MWUwYjQ4NWVmNjY0NzdiZTExMDU4YWE2YWExY2NjM2FmY2JlOGI1Y2E3NDUxNjQ5JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.dl-_O_xpSg7zRIOR6tqlcFxtklPu4k-dwmgImOY-AT8)

### Get Workflow Sequence Diagram
![execute_workflow_sd.png](https://private-user-images.githubusercontent.com/221332256/467560357-4f25e199-8e92-421b-b2ea-7c73d553502c.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NTI3NTgyOTAsIm5iZiI6MTc1Mjc1Nzk5MCwicGF0aCI6Ii8yMjEzMzIyNTYvNDY3NTYwMzU3LTRmMjVlMTk5LThlOTItNDIxYi1iMmVhLTdjNzNkNTUzNTAyYy5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjUwNzE3JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI1MDcxN1QxMzEzMTBaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT1kYzRkM2Y4NjI5NWEyN2Q4NGY4MTQwNTI5ZDFmNGY2M2QyYzc3ZWE2NTVjYjQyMDY5YjBmNzIyZmI4MDYzN2UzJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.iwV3ja1pgKIe3IwGm6WpxggmF-KECQmm6n1b4_ZY1Uw)