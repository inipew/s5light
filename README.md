# S5light

![GitHub all releases](https://img.shields.io/github/downloads/hang666/s5light/total)
![GitHub](https://img.shields.io/github/license/hang666/s5light)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/hang666/s5light)

![GitHub Repo stars](https://img.shields.io/github/stars/hang666/s5light?style=social)
![GitHub forks](https://img.shields.io/github/forks/hang666/s5light?style=social)
![GitHub followers](https://img.shields.io/github/followers/hang666?style=social)

### A lightweight socks5 proxy server and install script.

This software supports Windows/MacOS/Centos/Debian/Ubuntu and, in theory, all Linux.


Configuration file using YAML language.

Please fill in the template according to "config.yaml.example", and save it as "config.yaml".

You can also use the "--config" flag to customize the path to the config file.


## Features

 - Full TCP/UDP support
 - Multiple validation
 - IP address whitelisting
 - Specify out local address


## Install Script:

```bash
wget --no-check-certificate https://raw.githubusercontent.com/inipew/s5light/main/script/install.sh -O install_s5.sh && bash install_s5.sh
```
