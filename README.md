# Multicompose

Multicompose es una herramienta de línea de comandos (CLI) que permite automatizar la ejecución de múltiples comandos `docker-compose` de forma secuencial. Es útil para flujos de trabajo que requieren levantar contenedores en un orden específico y ejecutar scripts de inicialización entre los pasos.

## Descripción

Este programa CLI facilita la ejecución de los siguientes pasos de manera automatizada:

1. Levantar contenedores usando un archivo `docker-compose.yml` con el comando `docker-compose up --build -d`.
2. Ejecutar un script de inicialización, `init.sh`, que puede contener configuraciones adicionales o pasos de preparación para los contenedores.
3. Levantar contenedores adicionales con un segundo archivo `docker-compose-post.yml` usando `docker-compose up --build -d`.

## Instalación

Para instalar y usar `multicompose`, sigue estos pasos:

### Requisitos previos

- [Go](https://golang.org/dl/) 1.18 o superior.
- Docker y Docker Compose instalados en tu máquina.

### Clonación del repositorio

Clona el repositorio de `multicompose`:

```bash
git clone https://github.com/tu_usuario/multicompose.git
cd multicompose
```

### Construcción de la aplicación

Construye la aplicación usando Go:

```bash
go build -o multicompose .
```

### Instalación global (opcional)

Si deseas instalar la herramienta globalmente en tu sistema:

```bash
go install
```

Este comando instalará la herramienta en el directorio `$GOPATH/bin`.

## Uso

La herramienta `multicompose` se ejecuta desde la línea de comandos y tiene el siguiente uso básico:

```bash
multicompose up
```

### Comando `up`

Este comando ejecuta el siguiente flujo de trabajo:

1. Levanta los contenedores usando `docker-compose.yml`:
    ```bash
    docker-compose up --build -d
    ```
2. Ejecuta el script de inicialización `init.sh`:
    ```bash
    bash init.sh
    ```
3. Levanta los contenedores adicionales usando `docker-compose-post.yml`:
    ```bash
    docker-compose -f docker-compose-post.yml up --build -d
    ```

## Contribuciones

Si deseas contribuir al proyecto, puedes hacerlo de las siguientes maneras:

1. **Informar errores**: Si encuentras un bug o tienes alguna sugerencia, abre un [issue](https://github.com/tu_usuario/multicompose/issues).
2. **Realizar cambios**: Si tienes una mejora, haz un fork del proyecto, realiza tus cambios y abre un pull request.

## Licencia

Este proyecto está bajo la Licencia MIT - consulta el archivo [LICENSE](LICENSE) para más detalles.

## Contacto

- Correo electrónico: tu_email@dominio.com
- Repositorio: [https://github.com/tu_usuario/multicompose](https://github.com/tu_usuario/multicompose)

