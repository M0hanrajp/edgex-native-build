# ⚙️ EdgeX native build on Linux x86/x64

This repository is for building and running EdgeX on Linux natively, the current build is performed over Ubuntu 22.04 & all EdgeX core services built are cloned from v3.1.1 (Napa build).

* Objective: Make EdgeX native build execution process convenient !
    * This will be done through Makefile & will have similar functionalities to ease up execution process.
    * will include commands for the user to enable and disable features based on their use case.

---
## EdgeX Foundry Services Management

This Makefile provides a set of commands to manage various EdgeX Foundry services. It includes commands to start and stop services such as Consul, core services, support services, and more.

### Prerequisites

- Make sure the host machine as all [Required software](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#required-software) & [Prepare your environment](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#prepare-your-environment) environment variables are set.

### Directory Path

The directory path `~/edgex-foundry/edgex-native-build-3.1-napa` used in the Makefile will depend on where you clone this repository. You may need to update the Makefile to reflect your chosen directory.

For the Consul agent:

```makefile
# consul agent
start-consul-agent:
        @nohup consul agent -ui -bootstrap -server -client 127.0.0.1 -bind 127.0.0.1 -advertise 127.0.0.1 -data-dir=tmp/consul > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/edgex-consul-agent/nohup.out 2>&1 &
        @echo "::: EdgeX consul-agent is running... :::"
```
Make sure to update the paths in the Makefile according to your directory structure.

### Getting Started
To get started, ensure you have `make` installed on your system. Clone this repository and navigate to the directory containing the Makefile.
```bash
git clone https://github.com/M0hanrajp/edgex-native-build.git
cd edgex-native-build
```
### Commands
Note: In this guide, you will be building and running EdgeX in "non-secure" mode.
Prior to building and running EdgeX, set this environment variable to false. (By default, this variable is set to true).
```bash
$ export EDGEX_SECURITY_SECRET_STORE=false
```
### ▶️ Start Services
Start all EdgeX services with a single command:
```sh
make edgex-services-start
```
### ⏹️ Stop Services
Stop all EdgeX services gracefully:
```sh
make edgex-services-stop
```
This command will stop all running EdgeX services.
### 📊 Show Status
Check the status of all EdgeX services:
```sh
make edgex-services-show-status
```
### If you wish to enable/disable a service by name then please execute the below commands.
### ▶️ Start Services
```bash
make start-<service-name>
// example:
make start-core-metadata
```
### ⏹️ Stop Services
```bash
make kill-<service-name>
// example:
make kill-core-metadata
```
---
### 📜 Log Storage
Logs for each service will be stored in their respective service folders within the `edgex-service-logs` directory, located in the root directory where this repository is cloned.

### What is Makefile_advanced ?
This file is being implemented with better info (as start, running, failed!) on edgex services with better formatted output such as loading animations & checkmarks. This file is at it's earliest stage and will be improved with multiple tests while running the services.

Below are a few snapshots of the current Makefile_advanced:
```bash
# When starting a service
~/edgex-foundry/edgex-native-build-3.1-napa$ make -f Makefile_advance start-all-services
{Loading animation} Container edgex-core-consul: started {cursor}
...
# When the service is reported running
~/edgex-foundry/edgex-native-build-3.1-napa$ make -f Makefile_advance start-all-services
{checkmark logo} Container edgex-core-consul ::: Running
...
# When the service fails to run
~/edgex-foundry/edgex-native-build-3.1-napa$ make -f Makefile_advance start-all-services
{Loading animation} Container edgex-core-consul: failed!
```
### 📒 Notes
- Ensure to start the services in the order specified in the [EdgeX Foundry documentation](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#run-edgex).
- The kill commands use `SIGTERM` `(signal 15)` to stop the services gracefully. However, after stopping the service, the Consul UI will not indicate if the service is reachable or not.
- You can also use `SIGKILL` `(signal 9)`, which will forcefully stop the service and the Consul UI will indicate if a service is reachable or not.

This is WIP readme, will be updated over time.

### 🤝 Contributing

Contributions are welcome! Please fork this repository and submit a pull request with your changes. Let's make EdgeX services management even better together! 💪

---
Thanks 😊
