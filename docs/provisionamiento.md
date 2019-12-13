# Provisionamiento 🇻🅰️

## Por qué Vagrant 🤔

Se ha usado Vagrant porque con él podemos crear una máquina de VirtualBox por línea de comandos y es integrable con otros sistemas para provisionado, en concreto Ansible.

Esto hace reproducible de forma automática la creación de máquinas virtuales.

## Por qué Ansible 🤔

Ansible es una plataforma libre que gestiona todas las órdenes que hay que llevar a cabo para tener la máquina virtual provisionada.

Su configuración se basa en el uso de un fichero de scripting en formato YAML y nos recuerda a los ficheros de Makefile.

## Qué se ha hecho 🤷‍♂️

Se ha creado con Vagrant una máquina virtual (en adelante VM) en local. A esta VM se le ha instalado Ubuntu y se le han configurado puertos y direcciones. Aparte se le indica que se hará provisionado con Ansible, proporcionando la ruta donde está el fichero de configuración respectivo.

En el fichero de Ansible indicamos lo que queremos instalar/descargar, dónde queremos las cosas, limpieza a realizar de archivos descargados y, en definitiva, cualquier comando que haya que ejecutar dentro de esta VM para que funcione correctamente nuestro microservicio.

## Vagrantfile 🇻

En el archivo Vagrantfile especificamos:

    config.vm.box = "ubuntu/trusty64"

Esto nos descarga la versión de Ubuntu 18.04 estable de 64 bits.

Se ha elegido esta porque es la versión de Soporte Extendido (LTS) hasta 2022 y además es la estable. Aparte se ha elegido Ubuntu por ser lo que se ha usado hasta ahora en mi equipo local.

Después he indicado la configuración de red. Aquí asocio el puerto externo 5000 con el interno 5000.

    config.vm.network "forwarded_port", guest: 5000, host: 5000

Indicamos dónde está el fichero de provisionado:

    config.vm.provision "ansible_local" do |ansible|
        ansible.playbook = "provision/playbook.yml"
        ansible.verbose = true
    end

Vagrant además coloca el puerto 22 interno en el 2222 externo (aunque esto no nos debe preocupar si usamos ```vagrant ssh```)

## playbook.yml (fichero de Ansible 🅰️)

El fichero playbook.yml, colocado en la carpeta provision, realiza lo siguiente:

- Descarga la versión de Go necesaria (la 1.31.1 ha sido una de las testeadas en Travis, por eso se ha escogido):

      - name: Go 1.31.1
        command: curl -O https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
- Descomprime:

      - name: unzip go
        command: tar -xvf go1.13.1.linux-amd64.tar.gz
- Copia la carpeta descomprimida a /usr/local/go

      - name: move go lib
        copy: remote_src=True src=go dest=/usr/local
- Elimina el archivo de Go comprimido y la carpeta descomprimida:

      - name: delete downloaded source file
        command: rm  go1.13.1.linux-amd64.tar.gz
      - name: delete extracted go folder
        command: rm -R  go

- (Aunque no se pide en este hito) Instalamos git y clonamos nuestro repositorio con el microservicio. El ```force=yes``` actualiza el local con actualizaciones del remoto si hubiese:

        - name: Install git 
            apt:
                name: git
                state: present
        - name: Repo clone
            git:  repo=https://github.com/Jesus-Sheriff/go-InstaCrawler  force=yes dest=/home/vagrant/go-InstaCrawler
- Creamos enlace simbólico entre /usr/local/bin/go y /usr/local/go/bin/go para que la orden go sea reconocida en el sistema (sin tener que usar los export para indicar GOPATH y actualizar el PATH):

        - name: Symlink Go binary
            file:
            src: "/usr/local/go/bin/go"
            dest: "/usr/local/bin/go"
            state: link

----------------------
Nota: El fichero de ansible tiene muchas partes comentadas de pruebas que se han ido haciendo. El hecho de que estén es por si para el siguiente hito ayudan o hay que volver a ver esas opciones. También hay algún enlace que puede ayudar.

-----------------------

## Makefile

El Makefile tiene ahora cuatro nuevas  entradas:
   
    vm: vm-up vm-provision
    vm-up:
        vagrant up
    vm-provision:
        vagrant provision
    get-vm:
        vagrant init jesus-sheriff/go-InstaCrawler --box-version 1.0

- ```make vm``` levanta la máquina y provisiona.
- ```make vm-up``` levanta la máquina solo.
- ```make vm-provision``` provisiona solamente.
- ```make get-vm``` descarga la imagen (box) desde los servidores de Vagrant Cloud. Esta opción se incluye por si se desea descargar fácilmente la imagen.

## Conexión a la máquina en ejecución

Podemos conectarnos con ```vagrant ssh```. Esto nos conectará de forma transparente con la máquina por ssh.

También se puede acceder por ssh "puro" con la siguiente orden:
    ssh -p 2222 vagrant@127.0.0.1 -i ~/.vagrant.d/insecure_private_key

Se hace uso de una clave privada que ha creado Vagrant automáticamente. La contraseña por defecto del usuario "vagrant" es "vagrant".


## Vagrant Cloud 🆙☁️

Vagrant Box: https://app.vagrantup.com/jesus-sheriff/boxes/go-InstaCrawler

El asistente de Vagrant Cloud online permite subir una imagen de tipo "box" generada por Vagrant con la orden:
    
    vagrant package

Por defecto crea la imagen ```package.box```.

Se inicia el asitente gráfico desde Vagrant y es sencillo, lo único que puede resultar más complicado es que pide un "provider" que en nuestro caso es "virtualbox" y una suma de verificación.

Para hacer la suma se ejecuta:

    md5sum package.box

Una vez que está subida y publicada la imagen, se puede usar tanto en un archivo Vagrantfile como descargar en local. 

## Enlaces

Consultar sección de Enlaces en el README principal.