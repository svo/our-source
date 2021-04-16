# frozen_string_literal: true

# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"
PLAYBOOK = "infrastructure/ansible/playbook-development.yml"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.define "virtualbox" do |virtualbox|
    virtualbox.vm.box = "bento/debian-9"

    virtualbox.vm.hostname = "our-source-virtualbox"

    virtualbox.vm.provision "ansible" do |ansible|
      ansible.playbook = PLAYBOOK
      ansible.compatibility_mode = "2.0"
      ansible.extra_vars = { ansible_python_interpreter: "/usr/bin/python" }
    end

    virtualbox.cache.scope = :machine if Vagrant.has_plugin?("vagrant-cachier")

    virtualbox.vm.provider :virtualbox do |vb|
      vb.cpus = 2
      vb.memory = 1024
    end
  end
end
