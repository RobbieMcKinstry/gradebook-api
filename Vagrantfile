# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

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


  config.vm.box = "ubuntu/xenial64"
  config.vm.network "forwarded_port", guest: 443, host: 443
  config.vm.network "forwarded_port", guest: 8000, host: 8000
  config.vm.network "forwarded_port", guest: 5000, host: 5000
  config.vm.network "forwarded_port", guest: 8080, host: 8080
  # config.vm.network "forwarded_port", guest: 3306, host: 3306
  config.vm.network "forwarded_port", guest: 5672, host: 5672

  # removes the default shared folder.
  config.vm.synced_folder ".", "/vagrant", disabled: true

  config.vm.synced_folder gradebook_path,       guest_gradebook_path
  config.vm.synced_folder gradebook_api_path,   guest_gradebook_api_path
  config.vm.synced_folder autograder_path,      guest_autograder_path
  config.vm.synced_folder jobs_path,            guest_jobs_path

  config.ssh.forward_agent = true
  config.vm.provider "virtualbox" do |v|
    v.memory = 4096 # this line is required to make the box large enough for all the docker deps
    v.cpus = 2
  end

  config.vm.provision "shell", path: ".deploy/as_root_ubuntu.bash"

  # config.vm.provision "shell", path: ".appdeps/provision.bash"
  # config.vm.provision "shell", path: ".deploy/install_go.bash"
  # config.vm.provision "shell", path: ".deploy/as_user.bash", privileged: false
  #config.vm.provision "docker" do |d|
  #  d.run "mysql",  image: "mysql",    args: "-p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=root -e MYSQL_DATABASE=alligrader"
  #  d.run "rabbit", image: "rabbitmq:3.6.0-management", args: "-p 8080:15672 -p 5672:5672"
  #end

  # Define a Vagrant Push strategy for pushing to Atlas. Other push strategies
  # such as FTP and Heroku are also available. See the documentation at
  # https://docs.vagrantup.com/v2/push/atlas.html for more information.
  # config.push.define "atlas" do |push|
  #   push.app = "YOUR_ATLAS_USERNAME/YOUR_APPLICATION_NAME"
  # end
end
