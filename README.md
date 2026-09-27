# Curso de Go

Curso práctico de Go orientado a la construcción de una API HTTP aplicando los
principios de Clean Architecture. El proyecto comienza con un servicio de
consulta de usuarios en memoria y evoluciona mediante ejercicios incrementales.

## Objetivos

- Instalar y verificar el entorno de Go.
- Entender la estructura de un módulo y la organización por paquetes.
- Separar dominio, casos de uso y entrega HTTP.
- Trabajar con interfaces, errores, `context.Context` y JSON.
- Ejecutar, probar y mejorar una API utilizando las herramientas estándar de Go.

## Requisitos

- Ubuntu o una distribución compatible.
- Go 1.22 o superior.
- Git, recomendado para guardar el avance de cada etapa.

## Instalación

En Ubuntu se puede instalar Go desde los repositorios oficiales de Snap:

```bash
sudo snap install go --classic
```

Comprueba la instalación:

```bash
go version
```

## Estructura del curso

```text
curso-golang/
├── README.md
├── setup/
│   └── setup.go                 # Genera la estructura inicial
└── sources/
	 ├── go.mod                   # Módulo Go: go-clean-arch
	 ├── cmd/api/main.go          # Composición y arranque de la API
	 └── internal/
		  ├── domain/user.go       # Entidades, errores e interfaces
		  ├── usecase/
		  │   └── user_usecase.go  # Reglas de aplicación
		  └── delivery/http/
				└── user_handler.go  # Adaptador HTTP
```

### Etapas

1. **Fundamentos de Go**: módulos, paquetes, structs, métodos, interfaces y errores.
2. **Dominio**: definir `User`, el error `ErrUserNotFound` y los contratos de repositorio y caso de uso.
3. **Casos de uso**: implementar `GetUser` sin acoplarlo a HTTP ni a una base de datos.
4. **Entrega HTTP**: registrar rutas, leer parámetros, devolver estados HTTP y serializar JSON.
5. **Composición**: conectar el repositorio en memoria, el caso de uso y el handler desde `main`.
6. **Pruebas y mejoras**: añadir tests, validar entradas, mejorar errores y sustituir el repositorio en memoria.

## Preparar el proyecto

Desde la raíz del repositorio, genera o regenera los archivos base:

```bash
cd curso-golang
go run setup/setup.go
```

Después entra en el módulo y comprueba que el código compila:

```bash
cd sources
go test ./...
go build ./...
```

> El script de configuración sobrescribe los archivos que genera. Guarda tu trabajo antes de ejecutarlo de nuevo.

## Ejecutar la API

Desde `sources/`:

```bash
go run ./cmd/api
```

El servidor escucha en `http://localhost:8080`. En otra terminal, consulta el usuario disponible:

```bash
curl http://localhost:8080/users/1
```

Respuesta esperada:

```json
{"id":1,"name":"Gopher","email":"gopher@go.dev"}
```

Una consulta para un usuario inexistente devuelve `404 Not Found`:

```bash
curl -i http://localhost:8080/users/2
```

Para detener el servidor, pulsa `Ctrl+C`.

## Cómo completar cada etapa

1. Lee primero el código de `internal/domain`: identifica qué conoce el dominio y qué detalles externos desconoce.
2. Sigue el flujo de una petición desde `cmd/api/main.go` hasta `user_handler.go`, `user_usecase.go` y el repositorio en memoria.
3. Ejecuta los comandos de compilación y la API antes de modificar el código.
4. Completa un ejercicio cada vez y verifica el resultado con tests o `curl`.
5. Ejecuta `gofmt` sobre cada archivo Go modificado:

	```bash
	gofmt -w ruta/al/archivo.go
	```

6. Antes de pasar a la siguiente etapa, ejecuta:

	```bash
	go test ./...
	go vet ./...
	```

## Ejercicios propuestos

- Añadir tests unitarios para el caso de uso con un repositorio falso.
- Añadir tests del handler usando `net/http/httptest`.
- Devolver `400 Bad Request` cuando el ID no sea numérico o no esté presente.
- Diferenciar `404 Not Found` de otros errores del repositorio.
- Añadir un endpoint para listar usuarios.
- Implementar un repositorio en memoria con varios usuarios.
- Extraer la configuración del puerto a una variable de entorno.
- Añadir un `Dockerfile` y documentar cómo ejecutar la API en un contenedor.

## Comandos habituales

| Comando | Uso |
| --- | --- |
| `go run ./cmd/api` | Arrancar la API |
| `go test ./...` | Ejecutar todos los tests |
| `go test -v ./...` | Ejecutar tests mostrando detalles |
| `go vet ./...` | Detectar problemas habituales |
| `go build ./...` | Compilar todos los paquetes |
| `gofmt -w archivo.go` | Formatear un archivo Go modificado |

## Referencia rápida de la API

| Método | Ruta | Resultado |
| --- | --- | --- |
| `GET` | `/users/1` | Devuelve el usuario de ejemplo |
| `GET` | `/users/{id}` | Devuelve el usuario o `404` si no existe |
