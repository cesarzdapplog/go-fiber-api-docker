FROM golang:latest

WORKDIR /go/src/go-fiber-api-docker

# Copia solo el archivo go.mod y go.sum primero para descargar dependencias de forma eficiente
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Construye la aplicación
RUN go build -o bin/server cmd/main.go

# Define el comando para ejecutar la aplicación
CMD ["./bin/server"]
