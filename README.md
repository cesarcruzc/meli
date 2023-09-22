# Servicio de consulta de datos de MercadoLibre

## Descripción

Este proyecto consiste en un microservicio en Go que expone un endpoint para leer un archivo CSV, consultar APIs de
MercadoLibre y guardar los datos en una base de datos MongoDB.

## Construido con

- Go 1.21+
- Gin
- Godotenv
- MongoDB
- Docker

## Requisitos

- Docker
- Docker Compose
- Postman (opcional)
- Insomnia (opcional)

## Estructura del proyecto

```
proyecto-root/
│
├── api/
│   └── urls.go
│       - Define las rutas o endpoints y las asocia con sus controladores correspondientes.
│
├── config/
│   ├── bootstrap.go
│       - Maneja la inicialización y configuración de componentes esenciales, como la conexión a la base de datos y el logger.
├── core/
│   ├── csvprocessor.go
│       - Define la lógica para procesar archivos CSV.
│   ├── jsonlprocessor.go
│       - Define la lógica para procesar archivos JSONL.
│   ├── meliapi.go
│       - Define la lógica para interactuar con la API de MercadoLibre.
│   ├── meliclient.go
│       - Define la lógica para interactuar a nivel de cliente con la API de MercadoLibre.
│   ├── models.go
│       - Define la estructura `Item` que representa un ítem en la base de datos.
│   ├── process.go
│       - Define una función que inicia un proceso para manejar archivos.
│   ├── repository.go
│       - Define la lógica para interactuar con la base de datos MongoDB.
│   └── txtprocessor.go
│       - Esqueleto para el procesamiento de archivos TXT (sin implementación real).
│
├── v1/
│   ├── controllers.go
│       - Define los controladores que manejan las solicitudes HTTP.
│   └── middlewares.go
│       - Define middlewares para inyectar la base de datos y el logger en el contexto de Gin.
│
├── main.go
│   - Punto de entrada del proyecto. Inicializa componentes esenciales y arranca el servidor web.
│
└── data.csv
    - Una muestra de datos en formato CSV.

```

## Instalación

1. Clona el repositorio

```
git clone https://github.com/cesarcruzc/meli
```

2. Entra en la carpeta del proyecto

```
cd meli
```

## Ejecución

1. El proyecto utiliza variables de entorno para configurar el servicio. Estas variables se encuentran en el
   archivo `.env`

    - `MONGO_URI`: URI de conexión a la base de datos MongoDB
    - `MONGO_DB`: Nombre de la base de datos MongoDB
    - `API_URL`: URL del API de MercadoLibre
    - `API_TOKEN_URL`: URL de servicio externo para obtener token de acceso cada 5 horas
    - `X_API_KEY`: API Key para autenticación en el servicio externo
    - `FILE_PATH`: Ruta del archivo a procesar
    - `FILE_TYPE`: Tipo de archivo a procesar
    - `FILE_SEPARATOR`: Separador de columnas del archivo a procesar

2. Ejecuta el comando para construir la imagen de Docker y levantar los contenedores

  ```
  docker-compose up --build
  ``` 

3. El servicio estará disponible en la siguiente URL: `http://127.0.0.1:8888/`

## Swagger

La documentación de la API se encuentra disponible en la siguiente URL: `http://127.0.0.1:8888/api/v1/docs/index.html`

## API

La API expone los siguientes endpoints:

---

- Endpoint: `/`
- Método: `GET`
- Descripción: Inicio de la API

---

- Endpoint: `/health`
- Método: `GET`
- Descripción: Verifica el estado de la API

---

- Endpoint: `/api/v1/process-file`
- Método: `POST`
- Descripción: Procesa un archivo CSV y guarda los datos obtenidos del API de MercadoLibre en la base de datos

---

- Endpoint: `/api/v1/items`
- Método: `GET`
- Descripción: Obtiene todos los items
- Parámetros:
    - `page`: Página de items a obtener
    - `pageSize`: Cantidad de items por página


- Endpoint: `/api/v1/token`
- Método: `POST`
- Descripción: Obtiene un token de acceso para consumir el API de MercadoLibre de un servicio externo autenticado
  mediante apiKey

---

## Paso a paso API
1. Ejecutar el servicio de consulta de datos de MercadoLibre (meli)

```
curl --request POST \
  --url http://127.0.0.1:8888/api/v1/process-file \
  --header 'User-Agent: insomnia/2023.5.8'
```
2. Consultar los items guardados en la base de datos
```
curl --request GET \
  --url 'http://127.0.0.1:8888/api/v1/items?page=1&pageSize=10' \
  --header 'User-Agent: insomnia/2023.5.8'
```

## Contacto

- Autor: César Cruz
- Email: cc.cruz.caceres@gmail.com
- Phone: +57 315 275 8073