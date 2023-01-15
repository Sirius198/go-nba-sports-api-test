# NBA Sports API Consumer

## Setup

### Setup the database

Install [Postgres](https://www.postgresql.org/)

#### login to the database
```
$ sudo su - postgres
$ psql
```

#### create database and user

```
CREATE DATABASE nba;
CREATE USER '<your_user>' WITH ENCRYPTED PASSWORD '<your_password>';
GRANT ALL PRIVILEGES ON DATABASE nba TO '<your_user>';
```

#### run sql queries that you can find insde the `database/schema` folder

```
$ cd /path/to/go-nba-sports-api-test/database/schema
$ psql -U '<your_user>' -d nba -f team.sql
$ psql -U '<your_user>' -d nba -f player.sql
```

### Change `env` file

## REST APIs

### Get Games by Date

#### Request

`
GET /GamesByDate/{date}
`

#### Example

```
curl -i localhost:8080/GamesByDate/2022-OCT-01
```


### Players by Team

#### Request

`
GET /PlayersByTeam/{team}
`
#### Example
```
curl -i localhost:8080/PlayersByTeam/LAC
```

### Player details by Name

#### Request

`
GET /PlayerDetailsByName/{firstname}/{lastname}
`
#### Example
```
curl -i localhost:8080/PlayerDetailsByName/Jalen/Green
```
