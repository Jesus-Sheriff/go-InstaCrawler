Información preliminar relativa a Go
==================================================

Go actualmente tiene la versión estable 1.13.1 y da soporte hasta a dos versiones "major" anteriores (1.11 en este caso). 


Travis CI
    [![Build Status](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler.svg?branch=master)](https://travis-ci.com/Jesus-Sheriff/go-InstaCrawler)
    [![Dificultad](https://img.shields.io/badge/Dificultad-2%2F10-brightgreen)]()
=====================

TravisCI funciona con el manejador de versiones en Go `gimme`. Más información en el [GitHub de gimme](https://github.com/travis-ci/gimme). Se le han indicado las versiones de esta forma:

    go:
    - "1.11"
    - "1.11.x"
    - "1.12.x" 
    - "1.13.x"
    - master
Se ha testeado la versión 1.11 por ser la más antigua "mayor" con soporte. 
Las versiones terminadas en "x" indican la última "minor release" de esa versión "mayor". 

La versión 1.13.x testea en estos momentos la versión 1.13.1 que es la última minor release de 1.13. 
`master` es la última versión de go en desarrollo. Es importante/conveniente ejecutar en la última versión en desarrollo para ver si nuestro código va a seguir siendo válido en próximas releases del lenguaje.

En mi caso, hay que indicar las variables de entorno de usuario y contraseña de Instagram. Esto se ha hecho en la interfaz web. También he comprobado que en los logs no aparece el usuario ni la contraseña en texto legible.

Shippable
    [![Run Status](https://api.shippable.com/projects/5da439a382a9a900064c3542/badge?branch=master)]()
    [![Dificultad](https://img.shields.io/badge/Dificultad-4%2F10-green)]()
===============================================================

Shippable funciona con el manejador de versiones de Go `gvm`. Este manejador no acepta versiones relativas del lenguaje, hay que indicarle exactamente en qué versión queremos que testee. No se cómo indicar que testee en la versión de desarrollo.

    go:
    - "1.11"
    - "1.11.13"
    - "1.12.10" 
    - "1.13.1"

Para indicar las variables de entorno se usa un pequeño formulario de la interfaz web que nos las cifra para poder integrarlas en el archivo de configuración:

    env:
    - secure: XyUZf5yMxvoRjs8gdBMzXV3GM3anMIylfpTVsLhtlGuOVGznYmtu1ZIcXFAkP4uikoLF6eyvG83SSakbKkwA5kNAaYXL3DogjP8H4kwJ0HF5TiAke/xZJWAn9MKYYXw0PYgRIgGqBX2/eHf1pfrr1IR6cSTdByVZmOjXmFYpgsdG93yOGHBJoe3CXJ8RjRaqyVI1Bp8g0qWUuIKvTRCsIMFdYvqgqKylJozU2kP/xdmYPC2sFmC6ViATXMEVsjcwT9VA5VjWfB5GPXWsBznY9zMxYXMsb4REMsK2wyeTOgyl2UILGPtFEoBJrsUX78z4BOl9eQYm81TaxsgNV4iuzA==

CircleCI 
    [![CircleCI](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler.svg?style=svg)](https://circleci.com/gh/Jesus-Sheriff/go-InstaCrawler)
    [![Dificultad](https://img.shields.io/badge/Dificultad-9%2F10-red)]()
=============================================

CircleCI es un poco distinto de configurar. Su archivo de configuración es muy distinto al de Travis y Shippable. Se basa en YAML y un simple espacio puede hacer que no funcione su lectura (y fallen todos los test).

Ejemplo de test:

    docker:
        - image: circleci/golang:1.13.1
        steps:
        - checkout
        - run: make test

Lo he configurado para que ejecute los test en la versión 1.13.1 y además lo construya y ejecute haciendo uso de las órdenes de makefile.

Las variables de entorno se definen en la web al igual que en Travis.
