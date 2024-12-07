# Comandos de ejemplos para el despliegue 

## Investigacion y certificado

[Ver documento completo](Investigacion.pdf)
[Ver documento completo](certificado.pdf)


## Crear y corre un contenedor docker de las apliciones

### Estos comandos deben ejecutarse en la ruta de cada uno
```console
docker build -t tasks-management-app .

docker run -p 3000:3000 tasks-management-app
```

## Crear y correr un contenedor docker de posgret

```console
docker run -d --name postgres-tasks-management-app -p 5432:5432 -e POSTGRES_PASSWORD=12345678 -e POSTGRES_DB=taskMicro postgres:latest
```


## Exporter para La base de datos de users

```console
sudo docker run -d \
  --name postgres_exporter_users \
  -p 9187:9187 \
  -e DATA_SOURCE_NAME="host=192.168.0.168 port=5435 user=postgres password=12345678 dbname=usuarios sslmode=disable" \
  wrouesnel/postgres_exporter
```



## Exporter para la base de datos tasks

```console
sudo docker run -d \
  --name postgres_exporter_tasks \
  -p 9188:9187 \
  -e DATA_SOURCE_NAME="host=192.168.0.168 port=5432 user=postgres password=12345678 dbname=taskMicro sslmode=disable" \
  wrouesnel/postgres_exporter
```

## Cargar y lanzar prometheus

### Recordar usar una configuracion adecuada para su entorno
```console
docker run -d --name mi-prometheus-container -p 9090:9090 -v "$PWD/prometheus-config:/etc/prometheus" prom/prometheus --config.file=/etc/prometheus/prometheus.yml
```

## Cargar y lanzar grafana

```console
docker run -d --name mi-grafana-container -p 3000:3000 grafana/grafana
```


## Comandos de ayuda

Para saber  todo los contenedores que tenemos

```console
docker ps -a
```

Para saber el consumo de recursos de los contedores

```console
docker stats 
```

Para saber la direccion ip de un contenedor si es nesario


```console
sudo docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' task-manager-postgres
```
