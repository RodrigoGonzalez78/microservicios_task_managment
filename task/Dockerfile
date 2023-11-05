# Usar una imagen base de Go
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /tasks

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o tasks

# Exponer el puerto en el que se ejecuta la aplicación
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./tasks"]
