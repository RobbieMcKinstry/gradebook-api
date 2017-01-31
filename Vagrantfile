# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'fileutils'
require 'open-uri'
require 'tempfile'
require 'yaml'

Vagrant.require_version ">= 1.6.0"

$update_channel = "alpha"

kuber_ip                    = "192.168.5.2"
CLUSTER_IP="10.3.0.1"
NODE_IP = kuber_ip
NODE_VCPUS = 1
NODE_MEMORY_SIZE = 2048
USER_DATA_PATH = File.expand_path(".deploy/coreos-kubernetes/single-node/user-data")
SSL_TARBALL_PATH = File.expand_path("ssl/controller.tar")

Vagrant.configure(2) do |config|

  system("mkdir -p ssl && ./.deploy/coreos-kubernetes/lib/init-ssl-ca ssl") or abort ("failed generating SSL CA artifacts")
  system("./.deploy/coreos-kubernetes/lib/init-ssl ssl apiserver controller IP.1=#{NODE_IP},IP.2=#{CLUSTER_IP}") or abort ("failed generating SSL certificate artifacts")
  system("./.deploy/coreos-kubernetes/lib/init-ssl ssl admin kube-admin") or abort("failed generating admin SSL artifacts")

  gradebook_go_path     = '/src/github.com/alligrader/gradebook'
  gradebook_api_go_path = '/src/github.com/alligrader/gradebook-api'
  autograder_go_path    = '/src/github.com/alligrader/autograder'
  jobs_go_path          = '/src/github.com/alligrader/jobs'

  gradebook_path        = File.join(ENV['GOPATH'], gradebook_go_path)
  gradebook_api_path    = File.join(ENV['GOPATH'], gradebook_api_go_path)
  autograder_path       = File.join(ENV['GOPATH'], autograder_go_path)
  jobs_path             = File.join(ENV['GOPATH'], jobs_go_path)

  guest_gopath          = '/opt/gopath'

  guest_gradebook_path        = File.join(guest_gopath, gradebook_go_path)
  guest_gradebook_api_path    = File.join(guest_gopath, gradebook_api_go_path)
  guest_autograder_path       = File.join(guest_gopath, autograder_go_path)
  guest_jobs_path             = File.join(guest_gopath, jobs_go_path)


  config.vm.define "web", primary: true do |web|

      web.vm.box = "ubuntu/xenial64"
      web.vm.network "forwarded_port", guest: 443, host: 443
      web.vm.network "forwarded_port", guest: 8000, host: 8000
      web.vm.network "forwarded_port", guest: 5000, host: 5000
      web.vm.network "forwarded_port", guest: 8080, host: 8080
      # config.vm.network "forwarded_port", guest: 3306, host: 3306
      web.vm.network "forwarded_port", guest: 5672, host: 5672

      # removes the default shared folder.
      web.vm.synced_folder ".", "/vagrant", disabled: true

      web.vm.synced_folder gradebook_path,       guest_gradebook_path
      web.vm.synced_folder gradebook_api_path,   guest_gradebook_api_path
      web.vm.synced_folder autograder_path,      guest_autograder_path
      web.vm.synced_folder jobs_path,            guest_jobs_path

      web.ssh.forward_agent = true
      web.vm.provider "virtualbox" do |v|
        v.memory = 2048 # this line is required to make the box large enough for all the docker deps
        v.cpus = 2
      end
  end

  config.vm.define "kubernetes" do |kubernetes|
      kubernetes.vm.box = "apache"

      # always use Vagrant's insecure key
      kubernetes.ssh.insert_key = false

      kubernetes.vm.box = "coreos-%s" % $update_channel
      kubernetes.vm.box_version = ">= 1151.0.0"
      kubernetes.vm.box_url = "http://%s.release.core-os.net/amd64-usr/current/coreos_production_vagrant.json" % $update_channel

      kubernetes.vm.provider :virtualbox do |v|
        v.cpus = NODE_VCPUS
        v.gui = false
        v.memory = NODE_MEMORY_SIZE

        # On VirtualBox, we don't have guest additions or a functional vboxsf
        # in CoreOS, so tell Vagrant that so it can be smarter.
        v.check_guest_additions = false
        v.functional_vboxsf     = false
      end

      # plugin conflict
      if Vagrant.has_plugin?("vagrant-vbguest") then
        kubernetes.vbguest.auto_update = false
      end

      kubernetes.vm.network :private_network, ip: NODE_IP
      kubernetes.vm.network "forwarded_port", guest: 8001, host: 8001

      kubernetes.vm.provision :file, :source => SSL_TARBALL_PATH, :destination => "/tmp/ssl.tar"
      kubernetes.vm.provision :shell, :inline => "mkdir -p /etc/kubernetes/ssl && tar -C /etc/kubernetes/ssl -xf /tmp/ssl.tar", :privileged => true

      kubernetes.vm.provision :file, :source => USER_DATA_PATH, :destination => "/tmp/vagrantfile-user-data"
      kubernetes.vm.provision :shell, :inline => "mv /tmp/vagrantfile-user-data /var/lib/coreos-vagrant/", :privileged => true

  end

  # config.vm.provision "shell", path: ".deploy/as_root_ubuntu.bash"

  # config.vm.provision "shell", path: ".appdeps/provision.bash"
  # config.vm.provision "shell", path: ".deploy/install_go.bash"
  # config.vm.provision "shell", path: ".deploy/as_user.bash", privileged: false
  #config.vm.provision "docker" do |d|
  #  d.run "mysql",  image: "mysql",    args: "-p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=root -e MYSQL_DATABASE=alligrader"
  #end
end
