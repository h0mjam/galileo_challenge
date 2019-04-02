# Metrics service with in-memory data storage

## Build

```
$> go build .
```

## Run 
```
$> ./galileo
```

### Endpoints

 * User registration `http://localhost:3003/user`
 * Device registration `http://localhost:3003/device`
 * Append device metric `http://localhost:3003/metrics`
 * Devices list `http://localhost:3003/user/devices`
 * Device stat `http://localhost:3003/device/stat`


## Test
```
$> go test galileo/types galileo/store galileo/handler
```