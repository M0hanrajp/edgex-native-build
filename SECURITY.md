## EdgeX native build Security Policy

This project aims to improve the native build of EdgeX using Makefile it is not responsible for security issues or the support of EdgeX services. The services already present in the repository are just an example to showcase how the code can be cloned to your local machine and how provided Makefile can improve the development workflow.

In order to report any security issues with EdgeX services, please follow the below links.

### Supported Versions

Not every EdgeX release receives security support.
Please consult the EdgeX Foundry Wiki for information on the
[EdgeX Long Term Support Policy](https://wiki.edgexfoundry.org/pages/viewpage.action?pageId=69173332).


### Reporting Vulnerabilities

Instructions to report a vulnerability may be found on the
[security page of the EdgeX Foundry Wiki](https://wiki.edgexfoundry.org/display/FA/Security).


### Questions

Questions on the EdgeX security policy can be raised in the
[EdgeX Security Working Group](https://wiki.edgexfoundry.org/display/FA/Security+Working+Group?src=contextnavpagetreemode)
or to the 
[EdgeX Technical Steering Committee](https://wiki.edgexfoundry.org/pages/viewpage.action?pageId=329436&src=contextnavpagetreemode).

---

Below are the current package versions being used to explore the project.

## Supported Versions

Please follow the below versions of packages required to build and run edgex as intended.

### System Information
```sh
~$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.4 LTS
Release:        22.04
Codename:       jammy
```
- All below mentioned basic requirements where downloaded from [Build and Run on Linux on x86/x64](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/)

### Go Version
```sh
~$ go version
go version go1.22.4 linux/amd64
```

### GCC Version
```sh
~$ gcc --version
gcc (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0
Copyright (C) 2021 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

### Make Version
```sh
~$ make --version
GNU Make 4.3
Built for x86_64-pc-linux-gnu
Copyright (C) 1988-2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

### Consul Version
```sh
~$ consul --version
Consul v1.19.0
Revision bf0166d8
Build Date 2024-06-12T13:59:10Z
Protocol 2 spoken by default, understands 2 to 3 (agent will automatically use protocol >2 when speaking to compatible agents)
```

### Redis Version
```sh
~$ redis-server --version
Redis server v=7.2.5 sha=00000000:0 malloc=jemalloc-5.3.0 bits=64 build=d2f534f69a26fea
```

### ZMQ Library
```sh
~$ ldconfig -p|grep zmq
        libzmq.so.5 (libc6,x86-64) => /lib/x86_64-linux-gnu/libzmq.so.5
        libzmq.so (libc6,x86-64) => /lib/x86_64-linux-gnu/libzmq.so
```

### Git Version
```sh
~$ git --version
git version 2.34.1
```

All below docker related packages were downloaded from [Install Docker Engine](https://docs.docker.com/engine/install/) & [Install the Compose plugin
](https://docs.docker.com/compose/install/linux/)

### Docker Version
```sh
~$ docker --version
Docker version 27.0.1, build 7fafd33
```

### Docker Service Status
```sh
~/edgex$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
     Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
     Active: active (running) since Tue 2024-06-25 23:20:48 IST; 3min 44s ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 6213 (dockerd)
      Tasks: 27
     Memory: 45.1M
     CGroup: /system.slice/docker.service
             └─6213 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.436621695+05:30" level=warning msg="error locating sandbox id 9e316a365f2306ea3727e6b9655acdd6c596e>
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.437495118+05:30" level=info msg="Loading containers: done."
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456612139+05:30" level=warning msg="WARNING: No blkio throttle.read_bps_device support"
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456663540+05:30" level=warning msg="WARNING: No blkio throttle.write_bps_device support"
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456672034+05:30" level=warning msg="WARNING: No blkio throttle.read_iops_device support"
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456675641+05:30" level=warning msg="WARNING: No blkio throttle.write_iops_device support"
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456698467+05:30" level=info msg="Docker daemon" commit=ff1e2c0 containerd-snapshotter=false storage>
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.456778650+05:30" level=info msg="Daemon has completed initialization"
Jun 25 23:20:48 LIN-MP22QN1X dockerd[6213]: time="2024-06-25T23:20:48.531878377+05:30" level=info msg="API listen on /run/docker.sock"
Jun 25 23:20:48 LIN-MP22QN1X systemd[1]: Started Docker Application Container Engine.
```
