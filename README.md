# EdgeX native build on Linux x86/x64

### Objective

Create an end-to-end test environment that allows users to modify multiple services and test them using a Makefile. 
This setup is similar to Docker Compose but offers additional flexibility:

- **Modify Service Code**: Change the code of any service.
- **Edit Configuration**: Adjust the `configuration.yaml` or `configuration.toml` of any service.
- **Build and Execute**: Build the service and execute scenarios through the Makefile.
  - **Note**: If you need to modify an EdgeX service, clone it to your desired directory and update 
  the paths in the Makefile to point to your cloned path.
- **Run Services**: Start and manage the services.
- **EdgeX service config management**: Enable and disable features/configurations based on your use case.

### Additional Information

- You are not required to edit all services, but you have the option to do so.
- If you prefer a Docker Compose environment, please refer to the [Working in a Hybrid Environment](https://docs.edgexfoundry.org/3.1/getting-started/Ch-GettingStartedHybrid/) guide.

### Important Note

If your tasks require you to edit multiple codes, build, and execute them simultaneously, please follow the instructions below.

---

### Prerequisites

- Make sure the host machine as all [Required software](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#required-software) & [Prepare your environment](https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#prepare-your-environment) environment variables are set.
- Please `cd` to each edgex service folder and execute `make build` command & make sure all builds are completed.

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

### ⏳ View service status int two different ways:

If you want the view the status of the serviecs to be printed out to bash then use below command:
```bash
~/edgex-foundry/edgex-native-build-3.1-napa$ make edgex-services-show-status
USER         PID %CPU %MEM    VSZ   RSS   TTY      STAT START   TIME COMMAND
mpunix     54400  0.9  3.7 1438420 145792 pts/3  Sl   11:47   0:45 consul agent -ui -bootstrap -server -client 127.0.0.1 -bind 127.0.0.1 -advertise 127.0.0.1 -data-dir=tmp/consul -log-level=trace
mpunix     54410  0.0  0.6 1977516 23776 pts/3   Sl   11:47   0:01 ./core-metadata -cp=consul.http://127.0.0.1:8500 -registry -o
mpunix     54415  0.0  0.6 2125452 23960 pts/3   Sl   11:47   0:03 ./core-data -cp=consul.http://127.0.0.1:8500 -registry -o
mpunix     54420  0.0  0.6 2051068 25548 pts/3   Sl   11:47   0:01 ./core-command -cp=consul.http://127.0.0.1:8500 -registry -o
mpunix     54425  0.0  0.7 2134156 29256 pts/3   Sl   11:47   0:01 ./support-notifications -cp=consul.http://127.0.0.1:8500 -registry -o
mpunix     54430  0.0  0.7 2060080 29768 pts/3   Sl   11:47   0:04 ./support-scheduler -cp=consul.http://127.0.0.1:8500 -registry -o
mpunix     54440  0.0  0.7 1246844 28564 pts/3   Sl   11:47   0:04 ./app-service-configurable -cp=consul.http://127.0.0.1:8500 -registry -p=rules-engine -o
mpunix     54477  0.0  0.5 1242876 22520 pts/3   Sl   11:47   0:00 ./edgex-ui-server -o
mpunix     54466  0.0  1.1 1947564 43540 pts/3   Sl   11:47   0:00 ./kuiperd
mpunix     56483  0.0  0.7 1246888 28092 pts/3   Sl   11:54   0:03 ./device-virtual -cp=consul.http://127.0.0.1:8500 -registry -o
```
If you wish to view the service status from filtered htop view then use the below command:
```bash
~/edgex-foundry/edgex-native-build-3.1-napa$ make edgex-services-htop-status
::: Filtering EdgeX services and displaying with htop :::
>> you will be redirected to htop with filtered view of edgex services only
```
![image](https://github.com/user-attachments/assets/b8fbef57-6221-4727-a58f-f5b4a9d7e123)

You can download `htop` from [here](https://htop.dev/downloads.html)

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
