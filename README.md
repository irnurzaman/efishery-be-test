# eFishery BackEnd Skill Test
This repository contains 2 services. Auth services for users registration and authentication. Fetch services for fetching and aggregating data from 3rd API.

## Deployed Host
- Auth Service address : http://116.193.190.7:5000/
- Fetch Service address : http://116.193.190.7:8000/

## Usage
1. Go to root directory
2. In the root directory, run services with:
```sh
docker compose up #for Docker compose v2
```
or
```sh
docker-compose up #for Docker compose v1
```
3. Auth service will be listening on port 5000 and fetch service will be listening on port 8000

You can change the configuration by changing the `.env` file in the root directory

## Getting Started
Auth service consist of 3 API:

1. `/auth/register` for users registration
2. `/auth/login` for users login
3. `/auth/verify` for JWT verification

Detail auth service API documentation, go to `{{auth service addr}}/docs/index.html`. You will see API documentation as shown below:

![auth service API](https://github.com/irnurzaman/efishery-be-test/blob/documentation/assets/auth-service-documentation.png?raw=true)

Fetch service consist of 3 API:
1. `/fetch` for requesting data from eFishery API
2. `/aggregate` for requesting aggregated data from eFishery API
3. `/verify` for JWT verification

Detail fetch service API documentation, go to `{{auth service addr}}/docs`. You will see API documentation as shown below:

![fetch service API](https://github.com/irnurzaman/efishery-be-test/blob/documentation/assets/fetch-service-documentation.png?raw=true)

## Directory Tree
Root directory structure

```
.
├── app --> Services folder
│   ├── auth 
│   └── fetch
├── assets --> README.md assets folder
└── pkg --> Utility packages
    ├── logging
    └── security
```

Auth services directory structure
```
app
├── api --> REST API module
│   └── rest.go
├── docs --> Swagger documentation module
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── entities --> Service entities modules
│   └── entities.go
├── interfaces --> Service and repository interfaces modules
│   └── interfaces.go
├── models --> Request and response modules
│   └── models.go
├── repository --> Repository modules to database
│   └── repository.go
├── service --> Service's main logic modules
│   └── service.go
├── Dockerfile
├── auth
├── build.sh --> Shell script for building binary
├── config.env
└── main.go --> Main packages
```

Fetch services directory structure
```
app
├── Dockerfile
├── main.py --> REST API and entry point
├── models.py --> Request and response library
├── requirements.txt --> Library requirements list
└── service.py --> Service's main logic library
```

## Manual Usage
### Development environment
- Ubuntu 20.04 (WSL2)
- go version go1.15.4 linux/amd64
- python 3.8.5 (venv activated)
- sqlite3

### Steps
1. Clone git repository
```sh
git clone https://github.com/irnurzaman/efishery-be-test.git
```

2. Go to inside repository
```sh
cd efishery-be-test
```

3. Go to auth service folder and download go modules
```sh
cd app/auth
go mod download
```

4. Go to fetch service folder and download python libraries (from root repository)
```sh
cd app/fetch
pip install -r requirements.txt
```

5. For building auth service binary. Go to auth services folder. Then:
```sh
./build.sh
```
The output will be binary file named `auth`. Run `./auth --secret <secret> --services auth` for running the services. The services will be running on `http://localhost:5000`

6. For running fetch service. Go to fetch services folder. Then:
```sh
uvicorn main:app
```
 The services will be running on `http://localhost:8000`

 ## C4 Model
 ### System Context Diagram
 ![system context diagram](https://github.com/irnurzaman/efishery-be-test/blob/documentation/assets/system-context-diagram.png?raw=true)
 ### Deployment Diagram
 ![deployment diagram](https://github.com/irnurzaman/efishery-be-test/blob/documentation/assets/deployment-diagram.png?raw=true)