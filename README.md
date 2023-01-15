# NBA Sports API Consumer

## Get Games by Date

### Request

`
GET /GamesByDate/{date}
`

### Example

```
curl -i localhost:8080/GamesByDate/2022-OCT-01
```


## Players by Team

### Request

`
GET /PlayersByTeam/{team}
`
### Example
```
curl -i localhost:8080/PlayersByTeam/LAC
```

## Player details by Name

### Request

`
GET /PlayerDetailsByName/{firstname}/{lastname}
`
### Example
```
curl -i localhost:8080/PlayerDetailsByName/Jalen/Green
```
