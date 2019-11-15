# Despligue (Heroku)

## Por qué Heroku

Heroku es gratis para hacer lo básico que necesito y es "muy fácil" de hacer funcionar tu app en él. Aunque la documentación, de primeras, parecía buena y completa para Go, me he encontrado con muchos inconvenientes. 

No obstante, una vez que se ha configurado parece que funciona relativamente bien.

### ATENCIÓN: SPOILER! (Errores que han ido surgiendo)

Para empezar Heroku no detectaba el lenguaje. Esto era porque necesitaba un archivo `go.mod` que indica las dependencias que hay que conseguir. Este archivo es de "nueva implantación" en Go y entraba en conflicto con la orden `go get` que es la que yo usaba. Si usaba el archivo, los test fallaban desde la versión 1.13.* de Go. Esto lo he solucionado cambiando el valor de esta variable de entorno a off en los 3 sitios de CI : `GO111MODULE=off`. [(Más info sobre CI aquí)](CI.md)

Ahora había que crear el archivo `go.mod` indispensable para que Heroku detectase el lenguaje, pero para generar este archivo necesitamos las dependencias. Para esto he usado `godep`.

Cuando ya tenemos todo esto, ya se empieza a ver la luz al final del túnel y la aplicación comienza a funcionar después de hacer alguna cosilla más.

## Instalación de las herramientas de Heroku y primeros pasos

Instalar cli de Heroku e iniciar sesión:

```
sudo snap install heroku --classic
heroku login
```

Creamos la aplicación con el buildpack de Go: Un buildpack son unas herramientas que se nos dan para un lenguaje que permiten su compilación, ejecución, carga de scripts...

```
heroku create -b https://github.com/heroku/heroku-buildpack-go.git
```
Ahora hay que resolver dependencias y crear el archivo `go.mod` por lo que he dicho en la sección anterior.

```
godep save  ./...
go mod init pruebaheroku
```

## Script de configuración: Procfile

Ahora ya tenemos todo casi listo. Hay que crear un archivo Procfile con la siguiente línea de texto:
```
web:  make heroku
```
Esto indica que la aplicación va a ser de tipo web y que para lanzarla hay que ejecutar el comando `make heroku`

Y en el Makefile la orden `heroku` es:

```
heroku: 
	./bin/show-latest-image
```
Heroku tiene una instrucción esecífica en el makefile porque el buildpack construye el binario y lo coloca en un carpeta concreta con un nombre concreto. 

Ahora actualizamos cambios y hacemos push a Heroku:

```
git commit  -am "Pruebas con heroku"
git push heroku master
```
Ya estaría desplegada la app.

## Manejo de Logs: Papertrail

He usado Papertrails para tener el registro de logs a mano y centralizado.

```
heroku addons:create papertrail:choklad
heroku addons:open papertrail
```
`choklad` es una opción que indica que queremos la versión gratuita.

## Configuración de Push en Github + Deploy

La configuración debe hacerse desde la interfaz web. 
Vamos a Deploy y en Deployment Method escogemos GitHub. Indicamos repositorio y rama. En mi caso, por tener configurados los CI, he indicado que no haga despliegue si algún test falla (bad shit happens). Y pinchamos en automatic deploy.

## Funcionamiento correcto y ejemplos

[https://goinstacrawler.herokuapp.com/](https://goinstacrawler.herokuapp.com/) - Raiz

    {"status":"OK","URI":""}
[https://goinstacrawler.herokuapp.com/status](https://goinstacrawler.herokuapp.com/status) - Status

    {"status":"OK","URI":""}
[https://goinstacrawler.herokuapp.com/latest](https://goinstacrawler.herokuapp.com/latest) - Última imagen

    {"status":"OK","URI":"https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com\u0026_nc_cat=109\u0026bc=1571337657\u0026oh=121bcbbc0ebee792f067f0d9cfcd5549\u0026oe=5E2D09EB\u0026ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}

[https://goinstacrawler.herokuapp.com/latest/1](https://goinstacrawler.herokuapp.com/latest/1) - Imagen número 1

    {"status":"OK","URI":"https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/s1080x1080/72415541_145699950150997_7117851733228337833_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com\u0026_nc_cat=101\u0026bc=1571337657\u0026oh=4a5b93972d502b2e13b4627fb112f319\u0026oe=5E4A614E\u0026ig_cache_key=MjE2NjMyNTg0MjY4NjA0MzY5Mw%3D%3D.2"}

[https://goinstacrawler.herokuapp.com/latest/669697888](https://goinstacrawler.herokuapp.com/latest/669697888) - Imagen no existe con ese identificador

    {"status":"NOT - OK: 416 Requested Range Not Satisfiable","URI":""}

[https://goinstacrawler.herokuapp.com/holacaracola](https://goinstacrawler.herokuapp.com/latest/holacaracola) - URL no válida

    {"status":"NOT - OK: 404","URI":""}