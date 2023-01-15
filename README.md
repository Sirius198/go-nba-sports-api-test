# NBA Sports API Consumer

## Get Games by Date

### Request

```
curl localhost:8080/GamesByDate/{date}
```

### Example

```
curl localhost:8080/GamesByDate/2022-OCT-01
```


## Players by Team

### Request

```
curl localhost:8080/PlayersByTeam/{team}
```
### Example
```
curl localhost:8080/PlayersByTeam/LAC
```

## Player details by Name

### Request

```
curl localhost:8080/PlayerDetailsByName/{firstname}/{lastname}
```
### Example
```
curl localhost:8080/PlayerDetailsByName/Jalen/Green
```