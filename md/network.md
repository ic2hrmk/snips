# Find IP by MAC
~~~
sudo arp-scan --interface=enp2s0 --localnet | grep 2a:71:38:fd:95:6e
~~~

# Get list of network devices in network

~~~
nmap -sP 192.168.23.1/24
~~~

# Add static entry to LOCAL machine

~~~
/etc/hosts

127.0.0.1 localhost sql1
~~~

# Add ssh-keys on remote host

~~~
[How to](https://www.digitalocean.com/community/tutorials/how-to-configure-ssh-key-based-authentication-on-a-linux-server)
~~~
