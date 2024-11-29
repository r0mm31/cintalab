# Usa la imagen base de Go
FROM golang:1.21 AS builder-base

# Establece el directorio de trabajo
WORKDIR /src

# Copia los archivos necesarios al contenedor
COPY . .

# Construye la aplicación (ajusta según sea necesario)
RUN go build -o myapp .

# Imagen final para producción (puedes usar scratch o una imagen mínima)
FROM scratch

COPY --from=builder-base /src/myapp /myapp

ENTRYPOINT ["/myapp"]
