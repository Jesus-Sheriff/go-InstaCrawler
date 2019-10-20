# go-InstaCrawler

[![Build Status](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler.svg?branch=master)](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler)
[![Run Status](https://api.shippable.com/projects/5da439a382a9a900064c3542/badge?branch=master)]()
[![CircleCI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler.svg?style=svg)](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)



Microservicio que proporciona la URL de la última imagen publicada en Instagram de un hashtag dado.

Para otra información puedes ver el [antiguo README](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/README_2.md)

### Índice

<!-- TOC -->autoauto- [go-InstaCrawler](#go-instacrawler)auto        - [Índice](#índice)auto        - [Pre-requisitos 📋](#pre-requisitos-📋)auto        - [Instalación 🔧](#instalación-🔧)auto    - [Ejecutando las pruebas (tests) ⚙️](#ejecutando-las-pruebas-tests-⚙️)auto    - [Integración Continua 📦](#integración-continua-📦)auto    - [Construido con 🛠️](#construido-con-🛠️)auto    - [Licencia 📄](#licencia-📄)auto    - [Gracias a... 🎁](#gracias-a-🎁)autoauto<!-- /TOC -->

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

### Instalación 🔧

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
        $(GORUN) goinsta.v2/examples/show-latest-image/main.go

Primero comprueba dependencias (`deps`) y después compila el código y ejecuta.

Esta orden obtiene las dependencias (las actualiza si es necesario), ejecuta los test y si todo es correcto, ejecuta el programa.

Por defecto al ejecutarlo, muestra la última imagen con el hashtag #golang.

Un ejemplo de salida es:

```
2019/10/09 12:03:22 ultima foto:  https://scontent-mad1-1.cdninstagram.com/vp/7c8004a33e8ef83675e7c62a62c821d7/5E39388E/t51.2885-15/e35/70513351_167187977761265_1918517610523590583_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=105&se=8&ig_cache_key=MjE1MDc2MzEwMzMzMTk0ODE0Mw%3D%3D.2
```

## Ejecutando las pruebas (tests) ⚙️

Nota: el archivo de test de la clase principal con comentarios linea a linea está [aquí](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/goinsta.v2/tests/latest_image_test.go).

Para ejecutar todos los tests:

```
make test
```

Y debería dar como salida algo similar a:


```
=== RUN   TestImportAccount
--- PASS: TestImportAccount (5.23s)
    latest_image_test.go:38: URL is: https://scontent-mad1-1.cdninstagram.com/vp/42471a4ab5bc8a7db6936fb3d097da7d/5E22E36B/t51.2885-15/e35/p1080x1080/70194953_158216195247611_8124119613573040881_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=111&ig_cache_key=MjE1MDgxNTAyMTYxODEzNjkxMQ%3D%3D.2
    latest_image_test.go:40: logged into Instagram as user 'apuntabienminombre'
PASS
ok  	command-line-arguments	5.244s
```

La orden `make test` en el makefile es:

    test: deps
        $(GOTEST) -v ./...

Nuevamente requiere de las dependencias y después ejecuta los test en modo verbose `-v`.
El modo verbose muestra tanto los posibles fallos como las lías de log que haya.

## Integración Continua 📦

Los detalles de Integración Continua y su explicación para evaluación [aquí](CI.md)

Actualmente están configurados y en funcionamiento:

[Travis-CI](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler) para los tests.

En el archivo de configuración de Travis ( [.travis.yml](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/.travis.yml) ) están las distintas versiones usadas para testeo de la aplicación y su justificación.

[Circle-CI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler) para test y ejecución en la versión 1.13.1 de Go.

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
