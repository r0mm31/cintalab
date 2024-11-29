# Usar una imagen base oficial de Ubuntu
FROM ubuntu:22.04

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos necesarios al contenedor
COPY requirements.txt ./

# Instalar las dependencias necesarias
RUN apt-get update && \
    apt-get install -y python3 python3-pip && \
    pip3 install --no-cache-dir -r requirements.txt

# Copiar el código de la aplicación
COPY . .

# Definir el comando para ejecutar la aplicación
CMD ["python3", "app.py"]

# Configuración de reinicio automático en caso de fallo
HEALTHCHECK CMD curl --fail http://localhost:5000/ || exit 1

