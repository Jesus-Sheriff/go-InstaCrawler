
Vagrant.configure('2') do |config|
  config.vm.box = 'azure'
  config.vm.network "forwarded_port", guest: 5000, host: 5000
  # use local ssh key to connect to remote vagrant box
  config.ssh.private_key_path = '~/.ssh/id_rsa'
  config.vm.provider "azure" do |az, override|
    # Cogemos variables de entorno necesarias. Hay que hacer export antes de esto
    az.tenant_id = ENV['AZURE_TENANT_ID']
    az.client_id = ENV['AZURE_CLIENT_ID']
    az.client_secret = ENV['AZURE_CLIENT_SECRET']
    az.subscription_id = ENV['AZURE_SUBSCRIPTION_ID']
    
    # Specify VM parameters
    # NOmbre de la máquina virtual que queremos crear
    az.vm_name = 'goapp'
    
    # "Tamaño" de los recursos virtuales . Probé la B1ls (1 cpu y 500mb de RAM)
    # pero se quedaba corta de ram al provisionar y tuve que volver a esta que es la
    # más barata después de esa.
    az.vm_size = 'Standard_B1s'
    
    # La localización de nuestra máquina es importante si queremos 
    # cumplir con la normativa europea de protección de datos.
    # Debe ser la misma del grupo de recursos, si no fallará. (az group create -l westeurope -n goinsta2)
    az.location = 'westeurope'
    
    # La imagen a usar es Ubuntu Server 18.04 porque es la versión que uso en local y con la que se testeaba, además de ser de soporte extendido
    az.vm_image_urn = 'Canonical:UbuntuServer:18.04-LTS:latest'
    
    # Grupo de recursos donde se engloba esta máquina. (ya debe de haber sido creado)
    # Se crea con esta orden: az group create -l westeurope -n goinsta2
    az.resource_group_name = 'goinsta2'

    # Punto de entrada de nuestra app. Esta orden abre el puerto en el cortafuegos
    # y une el exterior con nuestra máquina.
    # (al principio creía que no era necesario y probé a abrir puertos directamente en el portal de Azure, pero
    # parece que es obligatorio)
    az.tcp_endpoints='5000'
  end
  config.vm.provision "ansible_local" do |ansible|
    ansible.playbook = "despliegue/playbook.yml"
    ansible.verbose = true
  end

end
