# go-InstaCrawler

[![Build Status](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler.svg?branch=master)](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler)
[![Run Status](https://api.shippable.com/projects/5da439a382a9a900064c3542/badge?branch=master)]()
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)



Microservicio que proporciona la URL de la última imagen publicada en Instagram de un hashtag dado.

Para otra información puedes ver el [antiguo README](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/README_2.md)

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

Para ejecutar:

```
make run
```

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
ok  	command-line-arguments	5.393s
```

Si además de los errores queremos que nos muestre un log, se puede usar la opción "-v" y la salida sería como esta:

```
go test -v gopkg.in/ahmdrz/goinsta.v2/tests/latest_image_test.go 
=== RUN   TestImportAccount
--- PASS: TestImportAccount (5.23s)
    latest_image_test.go:38: URL is: https://scontent-mad1-1.cdninstagram.com/vp/42471a4ab5bc8a7db6936fb3d097da7d/5E22E36B/t51.2885-15/e35/p1080x1080/70194953_158216195247611_8124119613573040881_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=111&ig_cache_key=MjE1MDgxNTAyMTYxODEzNjkxMQ%3D%3D.2
    latest_image_test.go:40: logged into Instagram as user 'apuntabienminombre'
PASS
ok  	command-line-arguments	5.244s
```

## Integración Continua 📦

Actualmente está configurado y en funcionamiento [Travis-CI](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler) para los tests.

En el archivo de configuración de Travis ( [.travis.yml](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/.travis.yml) ) están las distintas versiones usadas para testeo de la aplicación y su justificación.

La adaptación de [Circle-CI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler) está en proceso.

## Construido con 🛠️


* [gvm](https://github.com/moovweb/gvm) - Manejador de versiones de GO
* [dep](https://github.com/golang/dep) - Manejador de dependencias de GO



## Licencia 📄

Este proyecto está bajo la Licencia GPLv3 - mira el archivo [LICENSE](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/LICENSE) para detalles

## Gracias a... 🎁

* A los creadores de [goinsta](https://github.com/ahmdrz/goinsta)

* AL traductor de esta plantilla README [Villanuevand](https://github.com/Villanuevand)



---
⌨️ Plantilla adaptada de [Villanuevand](https://github.com/Villanuevand) 😊
