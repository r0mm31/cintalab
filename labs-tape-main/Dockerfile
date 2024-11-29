# Usar una imagen base oficial de Python
FROM python:3.11-slim

# Establecer el directorio de trabajo
WORKDIR /app

# Establecer PYTHONUNBUFFERED para deshabilitar el buffering
ENV PYTHONUNBUFFERED=1

# Copiar solo el archivo requirements.txt primero para aprovechar el caché

# Instalar dependencias (si tienes alguna)
RUN pip install --no-cache-dir -r 

# Copiar el resto del código al contenedor
COPY . .

# Definir el comando para ejecutar la aplicación
CMD ["python", "-u", "your_script.py"]  # Cambia "your_script.py" por el nombre real de tu script principal

# Opcionalmente, puedes agregar un HEALTHCHECK si tu aplicación lo requiere.
HEALTHCHECK CMD curl --fail http://localhost:5000/ || exit 1
