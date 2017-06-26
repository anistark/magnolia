# Magnolia

*Under Development.*


## Setup

```
cp config/config_sample.json config/config/json
```


## Run

```
go run main.go
```

## Run with Docker

### Building image

```
docker build -t magnolia .
```

### Run Docker image

```
docker run -p 8000:8080 --name magnolia_container magnolia
```

### Stop the container

```
docker stop [id of the container]
```

