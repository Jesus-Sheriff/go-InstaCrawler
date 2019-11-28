# Despliegue de contenedor Docker (Heroku) 🐳

## Por qué Heroku 🤔

Como ya había hecho el despliegue en Heroku, tenía todo el sisttema montado y el CLI instalado, así que no tenía nada que perder. Aparte al leer documentación de otros sitios como Google Cloud creo que podía ser más laborioso.

## Qué se ha hecho 🤷‍♂️

El proceso ha sido adaptar el microservicio que funcionaba como aplicación en la nube a un contenedor que "contiene" este microservicio. Desde fuera el funcionamiento es exactamente el mismo pero se han hecho algunos cambios internos para dockerizar la app. 

Se ha construido (build) la aplicación y el contenedor resultante ocupa en la actualidad unos 6MB gracias a las capacidades que tiene GO para generar ejecutables estáticos.

## Contenedor desplegado 📦👌

Desplegado en https://goinstacrawler.herokuapp.com/

Voy a intentar subirlo a Google Cloud (u otra nube) ya que en la práctica anterior no pude hacerlo.

## Descripción de los pasos necesarios 🔜

Primero descargamos la imagen Docker:

    docker pull jesussheriff/goinstacrawler

Cambiamos el tag de la imagen descargada al que nos indica Heroku en su documentación para incluirlo al registro:

    docker tag jesussheriff/goinstacrawler registry.heroku.com/goinstacrawler/web

Lo empujamos al registro: 

    docker push registry.heroku.com/goinstacrawler/web

Le decimos a Heroku finalmente que lo active:

    heroku container:release --app goinstacrawler  web

Todos estos pasos podrían incluirse en un script que queda como opción para mejora.

## Dockerfile 📄

------------
Nota: Se ha usado un pequeño truco aunque muy conocido: usar una imagen completa de alpine con golang como `builder` y generar un binario estático (sin referencias a librerías externas). 

Una vez creado, se copia a una imagen Docker "sin nada" que se denomina `scratch`. De esta forma resulta imperioso usar la orden `ENTRYPOINT` en lugar de `CMD`.

---------------

Primero hay que buscar un SO liviano que sustente nuestra aplicación. En este caso se ha usado alpine para golang. Se le pone de etiqueta `builder`.

    FROM golang:alpine AS builder

Se instala git, es necesario para el fetching de dependencias. Se le pasa argumento de no-cache para que no use version antigua cacheada.

    RUN apk update && apk add --no-cache git

Se especifica el directorio de trabajo y se copian los archivos desde él hacia nuestra imagen Docker. 

    WORKDIR $GOPATH/src/pruebadocker/
    COPY . .

(Sí, puede parecer feo el usar el `COPY . .` porque copia todo el contenido del directorio actual a la imagen, pero no le he dado mucha importancia porque como he dicho, esta fase contruye una imagen temporal.)

Se descargan depedencias con `go get`. Se podría usar `go mod`, de ahí que empezase usando el `COPY . .`, porque necesitaba varios archivos de la carpeta. 

    RUN go get -v github.com/gorilla/mux

Se compila el binario estático. Se especifica para reducir tamaño que el sistema de destino va a ser linux, que va a ser cgo (compiled go) y que se omitan marcas para depuración con `-ldflags="-w -s"`. Para más información de estos últimos flags: [stackoverflow.com](https://stackoverflow.com/questions/22267189/what-does-the-w-flag-mean-when-passed-in-via-the-ldflags-option-to-the-go-comman)

    RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/main goinsta.v2/examples/show-latest-image/main.go

En caso de que el ejecutable no sea estático, nos aparecerá el siguiente error en tiempo de ejecución:

    standard_init_linux.go:211: exec user process caused "no such file or directory"

Ahora vamos a la segunda fase:

Definimos imagen scratch de Docker.

    FROM scratch

Copiamos solamente el ejecutable de la imagen `builder` a esta imagen.

    COPY --from=builder /go/bin/main /go/bin/prueba

Especificamos el `ENTRYPOINT` y ponemos puerto por defecto. No haría falta realmente especificar el puerto aquí porque Heroku lo pone en una variable de entorno que se lee desde el ejecutable.

    ENTRYPOINT ["/go/bin/prueba", "--port", "5000"]

Se debe especificar `ENTRYPOINT` en lugar de `CMD` porque el primero indica a Heroku directamente dónde está el ejecutable. El segundo, en cambio, invoca a un shell de la imagen, lo cual falla ya que la imagen está vacía porque únicamente tiene a nuestro ejecutable.

## Fallos, problemas y pequeños detalles para quien quiera leer 
### (en realidad es para que mi yo futuro recuerde fallos típicos)

- No especificar el `ENTRYPOINT` en lugar de `CMD` para cualquier imagen scratch da fallo. 

- Cerciórate de que el ejecutable es estático (no malgastes dos tardes enteras en ver por qué "no encuentra el ejecutable". No es que no encuentre el ejecutable, es que no encuentra un shell en la imagen scratch porque no hay).

- En Heroku se lee de la variable `PORT` en tiempo de ejecución el puerto que te da en ese momento (no debes de preocuparte de eso ahora en el Dockerfile si en el PaaS ya funcionaba bien).

- Si tenías la aplicación en Heroku y quieres desplegar en Heroku, lo más probable es que se sobrescriba la aplicación de antes con la de Docker. Tenlo en cuenta.

- El comando `go mod` de go va muy bien "según dicen". A mi me parece un latazo y ahora mismo es más sencillo usar `go get`. Quizás en el futuro sea obligatorio o cambie en algo, tenlo en cuenta también.

