# go-InstaCrawler

[![Build Status](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler.svg?branch=master)](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler)
[![Run Status](https://api.shippable.com/projects/5da439a382a9a900064c3542/badge?branch=master)]()
[![CircleCI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler.svg?style=svg)](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)



Microservicio que proporciona la URL de la última imagen publicada en Instagram de un hashtag dado.

Para otra información puedes ver el [antiguo README](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/README_2.md)

### Índice

<!-- TOC -->
- [go-InstaCrawler](#go-instacrawler)
    - [Pre-requisitos 📋](#pre-requisitos-📋)        
    - [Instalación 🔧](#instalación-🔧)    
    - [Ejecutando las pruebas (tests) ⚙️](#ejecutando-las-pruebas-tests-⚙️)    
    - [Integración Continua 📦](#integración-continua-📦)    
    - [Construido con 🛠️](#construido-con-🛠️)    
    - [Licencia 📄](#licencia-📄)    
    - [Gracias a... 🎁](#gracias-a-🎁)
<!-- /TOC -->

### Pre-requisitos 📋


Se recomienda instalar en primer lugar un gestor de versiones como [gvm](https://github.com/moovweb/gvm) para poder probar el mismo programa en diferentes versiones del lenguaje.

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```
Nos situamos en nuestra carpeta de trabajo y procedemos a instalar la versión que queramos de go.

```
gvm install go1.13.1
gvm use go1.13.1
```

Nota sobre versiones:

* go actualmente tiene la versión estable 1.13.1 y da soporte hasta a dos versiones "major" anteriores (1.11 en este caso)
* Más información en los comentarios del archivo [.travis.yml](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/.travis.yml) y en la [información de versiones](https://golang.org/doc/devel/release.html) oficial de go.

### Instalación y uso 🔧

Clonamos repositorio

```
git clone https://github.com/Jesus-Sheriff/go-InstaCrawler
cd go-InstaCrawler
```

Debemos definir unas variables de entorno. Estas variables pueden hacerse permanentes si se añaden al archivo `$HOME/.bash_profile` y se ejecuta `source $HOME/.bash_profile`.
(NOTA: $PATH y $GOPATH deben definirse donde se desee trabajar, estos son los valores por defecto recomendados en la documentación).


```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin
export INSTAGRAM_USERNAME=_tu_nombre_de_usuario_
export INSTAGRAM_PASSWORD=_tu_contraseña_
```
Para usar las órdenes del makefile se requiere también del programa `pmgo`.

Este programa es un gestor de procesos escrito en go básico pero funcional.

```
go get github.com/struCoder/pmgo
mv $GOPATH/bin/pmgo /usr/local/bin
```

Por último, en el archivo `.env` se define el puerto del servicio, por defecto es el 8080.

#### Uso del microservicio

Para ejecutar el servicio podemos hacer:

* Ejecución completa de tests+servicio

```
make
```

Por defecto al hacer `make` ejecuta la orden del makefile `all: test run` que, después de descargar y actualizar dependencias, ejecuta el test y si es correcto compila y ejecuta el programa.

* Ejecución solo del servicio

```
make run
```

La orden `make run` en el makefile es esta:

    run: deps
	    pmgo start github.com/Jesus-Sheriff/go-InstaCrawler/goinsta.v2/examples/show-latest-image/ app


Primero comprueba dependencias (`deps`) y después compila el código y ejecuta.

Esta orden obtiene las dependencias (las actualiza si es necesario) y ejecuta el programa.

Por defecto al ejecutarlo, muestra la última imagen con el hashtag #golang.

Un ejemplo de salida es:

```
2019/10/09 12:03:22 ultima foto:  https://scontent-mad1-1.cdninstagram.com/vp/7c8004a33e8ef83675e7c62a62c821d7/5E39388E/t51.2885-15/e35/70513351_167187977761265_1918517610523590583_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=105&se=8&ig_cache_key=MjE1MDc2MzEwMzMzMTk0ODE0Mw%3D%3D.2
```

Una vez arrancado podemos hacer las siguientes llamadas:

* `/status`
    
    Devuelve el status del servicio.
* `/latest`

    Devuelve el estado de la llamada y el URI de la última imagen almacenada.
* `/latest/{id}`

    Devuelve el estado de la llamada y el URI de la imagen número `id`. Tanto si el número especificado es correcto o no, se obtiene una respuesta.


## Ejecutando las pruebas (tests) ⚙️

Nota: el archivo de test de la clase principal con comentarios linea a linea está [aquí](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/goinsta.v2/tests/latest_image_test.go).

Actualmente están todos los tests en un solo fichero.

En él están los tests funcionales que chequean las llamadas al microservicio (`/status`, `/latest` y `/latest/{id}`) y los tests unitarios.

Para ejecutar todos los tests:

```
make test
```

Y debería dar como salida algo similar a:


```
go test -v ./...
?   	go-InstaCrawler/goinsta.v2/examples/show-latest-image	[no test files]
=== RUN   TestGetStatus
--- PASS: TestGetStatus (0.00s)
    latest_image_test.go:49: getStatus correcto: '{Status: "OK", URI: ""}'
=== RUN   TestGetImage
--- PASS: TestGetImage (0.00s)
    latest_image_test.go:66: getImage correcto: '{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}'
=== RUN   TestGetImageNumber
--- PASS: TestGetImageNumber (0.00s)
    latest_image_test.go:83: getImageNumber correcto: '{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}'
=== RUN   TestNotFound
--- PASS: TestNotFound (0.00s)
    latest_image_test.go:101: notFound correcto: '{Status: "NOT - OK: 404", URI: ""}'
=== RUN   TestImportAccount
--- PASS: TestImportAccount (6.13s)
    latest_image_test.go:139: URL is: https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/73287971_786782621780776_5746821102179489882_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=104&bc=1571337657&oh=3af3acad6bd6b04d0e1ec0ef53260dd5&oe=5E63B29C&ig_cache_key=MjE2Njc0NDk2Mzk4Mzg0Mjg4OQ%3D%3D.2
    latest_image_test.go:141: logged into Instagram as user 'apuntabienminombre'
PASS
ok  	go-InstaCrawler/goinsta.v2/tests	6.144s

```

La orden `make test` en el makefile es:

    test: deps
        $(GOTEST) -v ./...

Nuevamente requiere de las dependencias y después ejecuta los test en modo verbose `-v`.
El modo verbose muestra tanto los posibles fallos como las líneas de log que haya.

## Integración Continua 📦

**Los detalles de Integración Continua y su explicación para evaluación [aquí](CI.md)**

Actualmente están configurados y en funcionamiento:

[Travis-CI](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler) para los tests y ejecución.

En el archivo de configuración de Travis ( [.travis.yml](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/.travis.yml) ) están las distintas versiones usadas para testeo de la aplicación y su justificación.

[Circle-CI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler) para test y ejecución en la versión 1.13.1 de Go (la ejecución falla, se exlica [aquí](CI.md)).

[Shippable](https://app.shippable.com/github/Jesus-Sheriff/go-InstaCrawler/dashboard) para tests.

## Construido con 🛠️


* [gvm](https://github.com/moovweb/gvm) - Manejador de versiones de GO
* [make](https://es.wikipedia.org/wiki/Make) - Para la gestión de dependencias, variables de entorno, ejecución de test y compilación y ejecución.

buildtool: Makefile

El archivo de makefile actualmente funciona correctamente para los tests, ejecución y resolución de dependencias. Se está añadiendo una forma de poder definir las variables de entorno desde aquí.

## Licencia 📄

Este proyecto está bajo la Licencia GPLv3 - mira el archivo [LICENSE](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/LICENSE) para detalles

## Gracias a... 🎁

* A los creadores de [goinsta](https://github.com/ahmdrz/goinsta)

* A  Radomir Sohlich por su [plantilla de ejemplo de makefile para Go](https://sohlich.github.io/post/go_makefile/)

* AL traductor de esta plantilla README [Villanuevand](https://github.com/Villanuevand)



---
⌨️ Plantilla adaptada de [Villanuevand](https://github.com/Villanuevand) 😊
