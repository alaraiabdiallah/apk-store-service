# APK Store Service

Small service to store your APKs

## System Requirement

- Docker engine v19.03.8

- Docker compose 1.25.4

## Environment Config

```
MONGO_HOST=
MONGO_PORT=
MONGO_USER=
MONGO_PASS=
MONGO_DBNAME=
HTTP_PROT_PORT=
API_KEY=
APP_BASE_URL=
```

| VARIABLE       | VALUE                   |
| -------------- | ----------------------- |
| MONGO_HOST     | default : localhost     |
| MONGO_PORT     | default : 27017         |
| MONGO_USER     | default : ""            |
| MONGO_PASS     | default : ""            |
| MONGO_DBNAME   | default : db-media      |
| HTTP_PROT_PORT | default : 2807          |
| API_KEY        | Mandatory Random string |
| APP_BASE_URL   | default: http://localhost:HTTP_PROT_PORT |

## How to run

```bash
docker-compose up -d --build
```
