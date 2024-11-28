# Declarar el argumento GOLANG_IMAGE al inicio
ARG GOLANG_IMAGE=docker.io/library/golang:1.21@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c947de4506981aa

# Usa la imagen base de Go
FROM --platform=${BUILDPLATFORM} ${GOLANG_IMAGE} as builder-base

# Establece el directorio de trabajo
WORKDIR /src

# Copia los archivos necesarios al contenedor
COPY . .

# Ejecuta las pruebas y verifica el código
RUN --mount=type=bind,target=/src --mount=target=/root/.cache,type=cache --mount=target=/go/pkg/mod,type=cache \
    go vet ./... && go test -v ./...

# Construye la aplicación (ajusta según sea necesario)
RUN go build -o myapp .

# Imagen final para producción (puedes usar scratch o una imagen mínima)
FROM scratch

COPY --from=builder-base /src/myapp /myapp

ENTRYPOINT ["/myapp"]

