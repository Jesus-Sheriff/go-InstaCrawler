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
    - [Deployment 📦](#Deployment-📦)
    - [Contenedor del microservicio 🐳](#contenedor)
    - [Provisionamiento 🇻🅰️](#provisionamiento)
    - [Despliegue final](#desplieguefinal)
    - [Licencia 📄](#licencia-📄)    
    - [Gracias a... 🎁](#gracias-a-🎁)
    - [Enlaces de interés y guías de ayuda](#Enlaces-de-interés-y-guías-de-ayuda)
<!-- /TOC -->

### Pre-requisitos 📋


Se recomienda instalar en primer lugar un gestor de versiones como [g](https://github.com/stefanmaric/g) para poder probar el mismo programa en diferentes versiones del lenguaje.

```
curl -sSL https://git.io/g-install | sh -s
```

Nos situamos en nuestra carpeta de trabajo y procedemos a instalar la versión que queramos de go. Podemos instalar la última versión así:

```
g install latest
g run latest
```
O consultar las versiones disponibles así:
```
g list-all
```


Nota sobre versiones:

* go actualmente tiene la versión estable 1.13.4 (released 2019/10/31) y da soporte hasta a dos versiones "major" anteriores (1.11 en este caso)
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
make deps
make run
```
`make deps` descarga dependencias.
La orden `make run` en el makefile es esta:

    go run goinsta.v2/examples/show-latest-image/main.go


Compila el código y ejecuta.

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

[Circle-CI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler) para test y ejecución en la versión 1.13.1 de Go.

[Shippable](https://app.shippable.com/github/Jesus-Sheriff/go-InstaCrawler/dashboard) para tests.

## Construido con 🛠️


* [gvm](https://github.com/moovweb/gvm) - Manejador de versiones de GO
* [make](https://es.wikipedia.org/wiki/Make) - Para la gestión de dependencias, variables de entorno, ejecución de test y compilación y ejecución.
* [godep](https://github.com/tools/godep) - Manejador de dependencias. (Necesario para tener archivo go.mod para Heroku)
* [g](https://github.com/stefanmaric/g) - Como leemos en su documentación: "Simple go version manager, gluten-free."

buildtool: Makefile

El archivo de Makefile actualmente funciona correctamente para los tests, ejecución y resolución de dependencias. 

## Deployment 📦

Despliegue: https://goinstacrawler.herokuapp.com/

El despliegue se ha hecho en Heroku y se está trabajando en hacerlo en Google Cloud también.

Como se ha indicado en la sección "Uso del microservicio" tenemos las siguientes llamadas disponibles a la API:
* `/status`
    
* `/latest`

* `/latest/{id}`

[Más información sobre el despliegue para su corrección aquí.](docs/despliegue.md)

## Contenedor del microservicio 🐳

Contenedor: https://goinstacrawler.herokuapp.com/

Contenedor alternativo: https://goinstacrawler.azurewebsites.net/

-----------------
Detalles adicionales para corrección, documentación extendida y decisiones de diseño en este enlace: [docs/contenedorDocker.md](docs/contenedorDocker.md)

------------------

Se ha hecho el despliegue del contenedor Docker en Heroku y en Azure.

El contenedor sincronizado con los cambios de este repositorio está en Docker Hub y puedes hacerle pull a tu local aquí: [jesussheriff/goinstacrawler](https://hub.docker.com/r/jesussheriff/goinstacrawler)

Una vez que lo tienes en tu máquina local, ejecuta los siguientes comandos.

```
docker build  -t my-golang-app .
docker run -p 5000:5000 my-golang-app
```

El microservicio estará accesible desde tu navegador en http://localhost:5000/

## Provisionamiento 🇻🅰️

provision: provision/playbook.yml

-----------------
Detalles adicionales para corrección, documentación extendida y decisiones de diseño en este enlace: [docs/provisionamiento.md](docs/provisionamiento.md)

------------------

El provisionamiento tiene dos fases: creación de la máquina virtual (con Vagrant 🇻) y provisionamiento de esta máquina con lo necesario (con Ansible 🅰️ en nuestro caso).

Tienes dos formas de comprobar el funcionamiento en local:

### 1. Si NO tienes este repositorio clonado localmente

Para usar la máquina creada se ejecutan los siguientes comandos:

    vagrant init jesus-sheriff/go-InstaCrawler --box-version 1.0
    vagrant up

El primero descarga la imagen (box) y el segundo la inicia.

### 2. Si tienes este repositorio clonado localmente

Solo necesitas hacer la siguiente orden que iniciará el proceso de provisionado completo (VM + Provisionado):

    make vm

Una vez hecha alguna de las opciones anteriores, puedes conectarte por ssh a la máquina:

    vagrant ssh

## Despliegue final

Despliegue final: 104.214.228.152:5000

Documentación adicional para corrección en [docs/Iaas.md](docs/Iaas.md)

## Enlaces de interés y guías de ayuda 

[godep](https://github.com/tools/godep) - Documentación de godep

[go modules](https://github.com/golang/go/wiki/Modules#quick-start) - Aquí explica la documentación oficial del blog de Go cómo crear un archivo `go.mod` de dependencias.

[Heroku con Go](https://devcenter.heroku.com/articles/getting-started-with-go?singlepage=true) - Documentación de Heroku para empezar con tu proyecto de go.

[Heroku con Go explicado de otra forma](https://medium.com/@freeformz/hello-world-with-go-heroku-38295332f07b) - Explica cómo crear el archivo Godeps de dependencias y cómo configurar el puerto (extreadamente útil).

[Error de Heroku: `(Web process failed to bind to $PORT within 60 seconds of launch)`](https://stackoverflow.com/questions/15693192/heroku-node-js-error-web-process-failed-to-bind-to-port-within-60-seconds-of) - Pregunta de StackOverflow con la que entendí por qué me daba ese fallo Heroku.

[Deploy an app in Google Cloud - Documentación oficial](https://cloud.google.com/appengine/docs/standard/go113/testing-and-deploying-your-app)

[Deploy an app in Google Cloud - Medium](https://medium.com/google-cloud/deploying-your-go-app-on-google-app-engine-5f4a5c2a837)

[Exponer puerto para contenedor en local](https://forums.docker.com/t/how-to-expose-port-on-running-container/3252/6)

[Crear imágenes light de go para Docker y compilar el ejecutable sin linkeo dinámico](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)

[Esta pregunta de stackoverflow solucionó mis problemas de linkeo estático.](https://stackoverflow.com/questions/56832363/docker-standard-init-linux-go211-exec-user-process-caused-no-such-file-or-di) 
Error: standard_init_linux.go:211: exec user process caused "no such file or directory"

[Flags de compilación](https://stackoverflow.com/questions/22267189/what-does-the-w-flag-mean-when-passed-in-via-the-ldflags-option-to-the-go-comman)



[Aquí empecé a ver cómo se provisionaba con Ansible](https://medium.com/@Joachim8675309/vagrant-provisioning-with-ansible-6dba6bca6290)

[De aquí aprendí a coger las dependencias en Ansible (OJO! Tiene mucha información útil muy bien explicada)](https://www.adictosaltrabajo.com/2015/09/04/creacion-de-entornos-de-integracion-con-ansible-y-vagrant/)

[De aquí vi cómo instalar Go, pero al ser un fichero de script (nosotros usamos Ansible) me daba problemas. En concreto con las órdenes export.](http://clouds.freeideas.cz/subdom/clouds/2017/08/02/ansible-install-golang-and-set-env-variables/)

[Cómo clonar un repositorio en tu máquina virtual con Ansible](https://docs.ansible.com/ansible/latest/modules/git_module.html)

[Instalar git en Ansible](https://www.edureka.co/community/41267/install-git-using-ansible)

[Conexión ssh con Vagrant](https://www.hashbangcode.com/article/connecting-vagrant-box-without-vagrant-ssh-command)

[Cómo copiar, mover o renombrar con ```mv``` en Ansible](https://stackoverflow.com/questions/24162996/how-to-move-rename-a-file-using-an-ansible-task-on-a-remote-system)

[CLI de Azure para Linux](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-apt?view=azure-cli-latest) Cómo instalar el CLI de Azure.

[Azure Github](https://github.com/Azure/vagrant-azure) Documentación oficial en el GitHub de Azure para configurar Vagrant. (Recomendado como primera lectura.)

[Uso de Vagrant con Azure](https://blog.scottlowe.org/2017/12/11/using-vagrant-with-azure/) Esta guía está más o menos actualizada y me ha servido mucho.

[Invocar make en Ansible](https://docs.ansible.com/ansible/latest/modules/make_module.htmlv) En la documentación oficial de Ansible explica bien cómo invocar una orden de Makefile desde el playbook.



## Licencia 📄

Este proyecto está bajo la Licencia GPLv3 - mira el archivo [LICENSE](https://github.com/Jesus-Sheriff/go-InstaCrawler/blob/master/LICENSE) para detalles

## Gracias a... 🎁

* A los creadores de [goinsta](https://github.com/ahmdrz/goinsta)

* A  Radomir Sohlich por su [plantilla de ejemplo de makefile para Go](https://sohlich.github.io/post/go_makefile/)

* AL traductor de esta plantilla README [Villanuevand](https://github.com/Villanuevand)



---
⌨️ Plantilla adaptada de [Villanuevand](https://github.com/Villanuevand) 😊
