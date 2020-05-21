GoMPing
==== 
This is a tool I created for studying Golang.

![sample](sample.jpg)

## Description
This is a parallel pinging tool for Linux/MAC.
It's very easy to use because it has only minimal features.

## Usage
Please create the configuration.
```shell
vi ./gomping.yml

# example
target_group:
  - group_name: group1
    target_host:
      - hostname: host1
        ip: 192.168.1.1
      - hostname: host2
        ip: 192.168.1.2
      - hostname: Google DNS
        ip: 2001:4860:4860::8888
  - group_name: group2
    target_host:
      - hostname: host3
        ip: 192.168.1.3
  - group_name: group3
    target_host:
      - hostname: host4
        ip: 192.168.1.4
```

and execute gomping.(need root permission)
```shell
sudo gomping
```

## Install
Just download it and unzip it.

## Licence
GoMPing is [MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author
[knhk](https://github.com/knhk)
