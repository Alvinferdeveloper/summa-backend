# Summa - Backend

Summa es una plataforma de búsqueda de empleo diseñada para conectar a candidatos con empleadores, con un fuerte enfoque en la inclusión y la accesibilidad. El backend está construido en Go y sigue una arquitectura modular para facilitar el mantenimiento y la escalabilidad.

## 🌟 Arquitectura del Sistema

*El siguiente diagrama ilustra la arquitectura general del sistema, incluyendo los servicios principales, las bases de datos y cómo interactúan entre sí.*

*(Aquí puedes insertar la imagen de la arquitectura del sistema. Ejemplo:)*
`![Arquitectura del Sistema](assets/architecture.png)`

---

## 🚀 Tecnologías Principales

El backend utiliza un conjunto de tecnologías modernas y robustas para asegurar un rendimiento y una fiabilidad óptimos:

- **Lenguaje:** Go (Golang)
- **Framework Web:** Gin
- **Base de Datos:** PostgreSQL
- **ORM:** GORM
- **Contenerización:** Docker y Docker Compose
- **Mensajería:** RabbitMQ
- **Caché:** Redis
- **Observabilidad:**
  - **Logging:** Loki
  - **Visualización:** Grafana
  - **Recolección de Logs:** Promtail
- **Almacenamiento de Archivos:** Cloudflare R2 (compatible con S3)
- **Notificaciones en Tiempo Real:** WebSockets

---

## 📂 Estructura del Proyecto

El proyecto sigue una estructura organizada para separar responsabilidades, inspirada en la arquitectura MVC y las mejores prácticas de Go.

```
/summa-backend
├─ /cmd
│ └─ main.go            # Punto de entrada de la aplicación y configuración inicial.
├─ /config              # Conexión a la base de datos, migraciones y configuración de servicios.
├─ /controllers         # Manejadores HTTP que reciben peticiones y envían respuestas.
├─ /database/seeders    # Seeders para poblar la base de datos con datos iniciales.
├─ /dto                 # Data Transfer Objects para validación y estructuración de datos.
├─ /middlewares         # Middlewares para autenticación, logging, CORS, etc.
├─ /models              # Estructuras de GORM que representan las tablas de la base de datos.
├─ /routes              # Definición de todos los endpoints de la API, agrupados por recurso.
├─ /services            # Lógica de negocio principal, separada de los controladores.
├─ /templates           # Plantillas HTML para correos electrónicos.
├─ /utils               # Funciones auxiliares (ej: manejo de JWT, contraseñas).
├─ /websocket           # Lógica del hub de WebSockets para notificaciones en tiempo real.
├─ go.mod / go.sum      # Gestión de dependencias del proyecto.
├─ docker-compose.yml   # Orquestación de servicios para el entorno de desarrollo.
└─ Dockerfile           # Definición del contenedor para la aplicación Go.
```

---

## 🛠️ Cómo Empezar (Getting Started)

Sigue estos pasos para configurar y ejecutar el proyecto en tu entorno de desarrollo local.

### Prerrequisitos

- [Go](https://golang.org/dl/) (versión 1.22 o superior)
- [Docker](https://www.docker.com/get-started) y [Docker Compose](https://docs.docker.com/compose/install/)
- Un editor de código como [Visual Studio Code](https://code.visualstudio.com/)

### 1. Clonar el Repositorio

```bash
git clone <URL-del-repositorio>
cd summa-backend
```

### 2. Configurar las Variables de Entorno

Crea un archivo `.env` en la raíz del directorio `summa-backend` a partir del archivo `.env.example` (si existe) o créalo desde cero. Este archivo contendrá todas las credenciales y configuraciones necesarias.

```env
# Archivo: .env

# Configuración de la Aplicación
PORT=8080

# Configuración de la Base de Datos (PostgreSQL)
DB_HOST=localhost
DB_USER=tu_usuario
DB_PASSWORD=tu_contraseña
DB_NAME=summa_db
DB_PORT=5433

# Credenciales de Email (SMTP)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=tu_correo@gmail.com
GOOGLE_APP_PASSWORD=tu_contraseña_de_aplicacion

# Credenciales de Cloudflare R2 (S3)
R2_ACCESS_KEY_ID=tu_access_key
R2_SECRET_ACCESS_KEY=tu_secret_key
R2_ACCOUNT_ID=tu_account_id
R2_PUBLIC_URL=tu_public_url

# URL de RabbitMQ y Redis (usadas por Docker Compose)
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
REDIS_URL=redis://:1234@localhost:6379/0
```

**Nota:** El `docker-compose.yml` está configurado para usar estas variables. Asegúrate de que los puertos no entren en conflicto con otros servicios en tu máquina.

### 3. Iniciar los Servicios con Docker Compose

Este comando levantará todos los servicios necesarios: PostgreSQL, RabbitMQ, Redis, Loki, Grafana y Promtail.

```bash
docker-compose up -d
```

### 4. Instalar Dependencias y Ejecutar la Aplicación

Una vez que los servicios de Docker estén en funcionamiento, puedes ejecutar la aplicación de Go.

```bash
# Instalar/actualizar dependencias
go mod tidy

# Ejecutar la aplicación
go run ./cmd/main.go
```

El backend estará disponible en `http://localhost:8080`.

---

## ⚙️ Configuración

La configuración del backend se gestiona principalmente a través de **variables de entorno**. El archivo `main.go` carga estas variables desde un archivo `.env` al iniciar, lo que permite una configuración flexible para diferentes entornos (desarrollo, producción, etc.).

Los servicios clave que se configuran de esta manera son:
- Conexión a la base de datos.
- Credenciales del servicio de correo.
- Claves de acceso para el almacenamiento de objetos (R2).
- Conexión a Redis y RabbitMQ.

---

## 📊 Observabilidad (Logging y Monitoreo)

El proyecto está configurado con un stack de observabilidad para facilitar el monitoreo y la depuración:

- **Loki:** Centraliza los logs generados por la aplicación y otros servicios.
- **Promtail:** Recolecta los logs y los envía a Loki.
- **Grafana:** Permite visualizar y consultar los logs almacenados en Loki. Puedes acceder a Grafana en `http://localhost:3001`.

Para ver los logs de la aplicación, configura una fuente de datos Loki en Grafana apuntando a `http://loki:3100` y usa etiquetas como `{job="promtail"}` para filtrar los logs.
