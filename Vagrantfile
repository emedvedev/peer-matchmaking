# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.define "mgmt_server" do |server|
    server.vm.box = "ubuntu/trusty64"
    server.vm.network "public_network"
    server.vm.network "private_network", virtualbox__intnet: "mgmt", ip: "192.168.50.10"
    server.vm.provision "shell", inline: <<-SHELL
    SHELL
  end
  config.vm.define "client_server_1" do |server|
    server.vm.box = "ubuntu/trusty64"
    server.vm.network "private_network", virtualbox__intnet: "client_1", ip: "172.30.90.10"
    server.vm.provision "shell", run: "always", inline: <<-SHELL
        ip route add 172.30.91.0/24 via 172.30.90.2
    SHELL
  end
  config.vm.define "client_server_2" do |server|
    server.vm.box = "ubuntu/trusty64"
    server.vm.network "private_network", virtualbox__intnet: "client_2", ip: "172.30.91.10"
    server.vm.provision "shell", run: "always", inline: <<-SHELL
        ip route add 172.30.90.0/24 via 172.30.91.2
    SHELL
  end
  config.vm.define "isp" do |router|
    router.vm.box = "higebu/vyos"
    router.vm.network "private_network", virtualbox__intnet: "ixp", ip: "192.168.90.10"
    router.vm.network "private_network", virtualbox__intnet: "isp", ip: "192.168.91.2"
    router.vm.provision "shell", privileged: false, inline: <<-SHELL
      source /opt/vyatta/etc/functions/script-template
      set protocols bgp 65000 neighbor 192.168.91.10 remote-as 65001
      set protocols bgp 65000 neighbor 192.168.91.11 remote-as 65002
      set protocols bgp 65000 parameters router-id 192.168.91.2
      commit
      save
      exit
    SHELL
  end

  config.vm.define "src" do |router|
    router.vm.box = "higebu/vyos"
    router.vm.network "private_network", virtualbox__intnet: "ixp", ip: "192.168.90.20"
    router.vm.network "private_network", virtualbox__intnet: "isp", ip: "192.168.91.10"
    router.vm.network "private_network", virtualbox__intnet: "client_1", ip: "172.30.90.2"
    router.vm.network "private_network", virtualbox__intnet: "mgmt", ip: "192.168.50.20"
    router.vm.provision "shell", privileged: false, inline: <<-SHELL
      source /opt/vyatta/etc/functions/script-template
      configure
      set protocols bgp 65001 neighbor 192.168.91.2 remote-as 65000
      set protocols bgp 65001 redistribute connected
      set protocols bgp 65001 parameters router-id 192.168.91.10
      set system flow-accounting interface eth2
      set system flow-accounting sflow server 192.168.50.10
      set system flow-accounting sflow sampling-rate 100
      commit
      save
      exit
    SHELL
  end

  config.vm.define "dst" do |router|
    router.vm.box = "higebu/vyos"
    router.vm.network "private_network", virtualbox__intnet: "ixp", ip: "192.168.90.30"
    router.vm.network "private_network", virtualbox__intnet: "isp", ip: "192.168.91.11"
    router.vm.network "private_network", virtualbox__intnet: "client_2", ip: "172.30.91.2"
    router.vm.provision "shell", privileged: false, inline: <<-SHELL
      source /opt/vyatta/etc/functions/script-template
      set protocols bgp 65002 neighbor 192.168.91.2 remote-as 65000
      set protocols bgp 65002 redistribute connected
      set protocols bgp 65002 parameters router-id 192.168.91.11
      commit
      save
      exit
    SHELL
  end
end
