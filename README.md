# EdgeX native build on Linux x86/x64

This repository is for building and running EdgeX on Linux natively, the current build is performed over Ubuntu 22.04 & all EdgeX core services built are cloned from v3.1.1 (Napa build).

* Objective: Make EdgeX native build execution process convenient !
    * This will be done through Makefile & will have similar functionalities to ease up execution process.
    * will include commands for the user to enable and disable features based on their use case.

---

### EdgeX Foundry Service Management

This Makefile provides a set of commands to manage various EdgeX Foundry services. It includes commands to start and stop services such as Consul, core services, support services, and more.

### Prerequisites

- Make sure the host machine as all mentioned [Required software](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#required-software) & [Prepare your environment](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#prepare-your-environment) 

### Usage

Sure! Here's the common syntax for executing the services along with an example:

- **Start a service:** 
```bash
make start-<service-name> 
// example:
make start-core-metadata
```
- **Stop a service:**
```bash
make kill-<service-name>
// example:
make kill-core-metadata
```
You can mention it like this:

### Log Storage

Logs for each service will be stored in their respective service folders within the `edgex-service-logs` directory, located in the root directory where this repository is cloned.

### Notes

- Ensure to start the services in the order specified in the [EdgeX Foundry documentation](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#run-edgex).
- The kill commands use `SIGTERM` `(signal 15)` to stop the services gracefully. However, after stopping the service, the Consul UI will not indicate if the service is reachable or not. 
- You can also use `SIGKILL` `(signal 9)`, which will forcefully stop the service and the Consul UI will indicate if a service is reachable or not.

This is WIP readme, will be updated over time.

---
Thanks 😊
