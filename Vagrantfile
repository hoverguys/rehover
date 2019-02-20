# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
	config.vm.provider "virtualbox" do |v|
		v.memory = 2048
	end

	config.vm.box = "ubuntu/xenial64"
	config.vm.provision :docker
	config.vm.provision :docker_compose, yml: "/vagrant/docker-compose.yml", rebuild: true, run: "always"
end