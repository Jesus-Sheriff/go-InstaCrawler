

[![Build Status](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler.svg?branch=master)](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

# go-InstaCrawler
Instagram tag crawler based on golang

This microservice will provide images that matches an especified tag on Instagram. It would be based on [goinsta](https://github.com/ahmdrz/goinsta). I'll storage the data in a MySQL database. I'll use papertrails for logging and REST technology.

Este microservicio proporcionará imágenes que compartan un tag de Instagram. Estará basado en [goinsta](https://github.com/ahmdrz/goinsta). Los datos se almacenarán en una base de datos MySQL. Se usará papertrails para logs y tecnología REST. Go se usa para la lógica del servicio en sí (la API que en este caso se conecta  a Instagram) y para las comunicaciones clientes servidor con el [paquete http](https://golang.org/pkg/net/http/).

## Por qué un proyecto sobre Instagram

Está de moda y puedo hacer algo que me resulte útil.

## Por qué uso [goinsta](https://github.com/ahmdrz/goinsta)

Primero, porque está basado en Go y quería aprender un poco cómo funciona un lenguaje "nuevo" (al menos para mi) que 
está especialmente diseñado para aplicaciones cloud.

Además es un proyecto bastante activo en GitHub.

## Por qué papertrails

Es necesario mantener un registro de logs y este es "gratis".
No conozco otro sistema y creo que puede ser sencillo de configurar.

## Por qué MySQL

Porque la conozco y he visto que se integra bien con GO.

## Por qué REST

Porque es la interfaz que necesito, no preciso de más. 

## ¿Es segura la comunicación con el servidor?

Hay que usar un timeout para evitar Hijacking como se dice [aquí](https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779).


## STATUS

Actualmente estoy haciendo pruebas con los ejemplos y se han ejecutado correctamente en local.
Se producen fallos actualmente (quizás por conexiones consecutivas a Instagram).
