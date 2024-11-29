# Usar una imagen base oficial de Python
FROM python:3.11-slim

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar todos los archivos del proyecto al contenedor
COPY . .

# Instalar dependencias (si tienes alguna)
# Si tienes dependencias, puedes crear un requirements.txt más adelante
# RUN pip install --no-cache-dir -r requirements.txt

# Definir el comando para ejecutar la aplicación
# Cambia "your_script.py" por el nombre de tu script principal
CMD ["python", "your_script.py"]

