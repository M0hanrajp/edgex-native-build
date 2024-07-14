# Simple Makefile for easy execution of tasks

# consul agent
start-consul-agent:
	@nohup consul agent -ui -bootstrap -server -client 127.0.0.1 -bind 127.0.0.1 -advertise 127.0.0.1 -data-dir=tmp/consul > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/edgex-consul-agent/nohup.out 2>&1 &
	@echo "::: EdgeX consul-agent is running... :::"

# This command forcefully kills consul, there is no information in edgex documentation on how to stop consul once we start it
kill-consul-agent:
	@ps aux | grep 'consul agent' | grep -v grep | awk '{print $$2}' | xargs kill -15
	@echo "::: EdgeX consul-agent has stopped working :::"

# ==== EdegeX core services ====
# Start core-common-config-bootstrapper service
# This service will exit once it has seeded the Configuration Provider with the common config. ( Based on documentation )

start-core-common-config-bootstrapper:
	@cd edgex-go-3.1.1/cmd/core-common-config-bootstrapper && nohup ./core-common-config-bootstrapper -cp=consul.http://127.0.0.1:8500 -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/core-common-config-bootstrapper/nohup.out 2>&1 &
	@echo "::: core-common-config-bootstrapper has pushed configuraitons :::"

# edgex core-metadata service
start-core-metadata:
	@cd edgex-go-3.1.1/cmd/core-metadata/ && nohup ./core-metadata -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/core-metadata/nohup.out 2>&1 &
	@echo "::: EdgeX core-metadata is running... :::"
	
kill-core-metadata:
	@ps aux | grep "core-metadata" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX core-metadata has stopped working :::" 
# It's important to start the services as listed in the order from https://docs.edgexfoundry.org/3.1/getting-started/native/Ch-BuildRunOnLinuxDistro/#run-edgex

# edgex core-data service
start-core-data:
	@cd edgex-go-3.1.1/cmd/core-data/ && nohup ./core-data -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/core-data/nohup.out 2>&1 &
	@echo "::: EdgeX core-data is running... :::"
	
kill-core-data:
	@ps aux | grep "core-data" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX core-data has stopped working :::"

# edgex core-command service
start-core-command:
	@cd edgex-go-3.1.1/cmd/core-command/ && nohup ./core-command -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/core-command/nohup.out 2>&1 &
	@echo "::: EdgeX core-command is running... :::"
	
kill-core-command:
	@ps aux | grep "core-command" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX core-command has stopped working :::" 

# edgex support-notifications service
start-support-notifications:
	@cd edgex-go-3.1.1/cmd/support-notifications/ && nohup ./support-notifications -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/support-notifications/nohup.out 2>&1 &
	@echo "::: EdgeX support-notifications is running... :::"
	
kill-support-notifications:
	@ps aux | grep "support-notifications" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX support-notifications has stopped working :::" 

# edgex support-scheduler service
start-support-scheduler:
	@cd edgex-go-3.1.1/cmd/support-scheduler/ && nohup ./support-scheduler -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/support-scheduler/nohup.out 2>&1 &
	@echo "::: EdgeX support-scheduler is running... :::"
	
kill-support-scheduler:
	@ps aux | grep "support-scheduler" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX support-scheduler has stopped working :::"

# edgex app-service-configurable service
start-app-service-configurable:
	@cd app-service-configurable-3.1.1/ && nohup ./app-service-configurable -cp=consul.http://127.0.0.1:8500 -registry -p=rules-engine -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/app-service-configurable/nohup.out 2>&1 &
	@echo "::: EdgeX app-service-configurable is running... :::"
	
kill-app-service-configurable:
	@ps aux | grep "app-service-configurable" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX app-service-configurable has stopped working :::"

# edgex device-virtual service
start-device-virtual:
	@cd device-virtual-go-3.1.1/cmd && nohup ./device-virtual -cp=consul.http://127.0.0.1:8500 -registry -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/device-virtual/nohup.out 2>&1 &
	@echo "::: EdgeX device-virtual is running... :::"
	
kill-device-virtual:
	@ps aux | grep "device-virtual" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	@echo "::: EdgeX device-virtual has stopped working :::"

# edgex ekuiper service (rules engine)
start-ekuiper:
	@export CONNECTION__EDGEX__REDISMSGBUS__PORT=6379
	@export CONNECTION__EDGEX__REDISMSGBUS__PROTOCOL=redis
	@export CONNECTION__EDGEX__REDISMSGBUS__SERVER=localhost
	@export CONNECTION__EDGEX__REDISMSGBUS__TYPE=redis
	@export EDGEX__DEFAULT__PORT=6379
	@export EDGEX__DEFAULT__PROTOCOL=redis
	@export EDGEX__DEFAULT__SERVER=localhost
	@export EDGEX__DEFAULT__TOPIC=rules-events
	@export EDGEX__DEFAULT__TYPE=redis
	@export KUIPER__BASIC__CONSOLELOG="true"
	@export KUIPER__BASIC__RESTPORT=59720
	@echo ":::All environment variables set :::"
	@cd ekuiper-1.14.0/_build/kuiper--linux-amd64/bin && nohup ./kuiperd > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/ekuiper/nohup.out 2>&1 &
	@echo "::: EdgeX ekuiper is running... :::"
	
# Pid is a variable here which will contain pid of the process running at 9081.
# The condition will only be entered when pid will not be empty.
#     1. **`@pid=`**:
#    - This part assigns the output of the following command to the variable `pid`.
#    2. **`$$(sudo lsof -t -i :9081)`**:
#    - `$$`: In a `Makefile`, `$$` is used to escape the dollar sign, so it gets passed to the shell as a single `$`.
#    - `sudo`: Runs the command with superuser privileges.
#    - `lsof`: Stands for "list open files". It's a command used to find out which files are open by which processes.
#    - `-t`: This option tells `lsof` to output only the process IDs (PIDs) of the processes using the specified files.
#    - `-i :9081`: This option tells `lsof` to look for processes using the network port 9081.
#    So, the command `sudo lsof -t -i :9081` lists the PIDs of processes using port 9081. The `$$` ensures that the output
#    of this command is assigned to the `pid` variable in the `Makefile`.
kill-ekuiper:
	@pid=$$(sudo lsof -t -i :9081); \
	if [ -n "$$pid" ]; then \
		echo "Killing process $$pid using port 9081"; \
		sudo kill -15 $$pid; \
	else \
		echo "No process using port 9081"; \
	fi
	@echo "::: EdgeX ekuiper has stopped working :::" 

# edgex ui-server service
# The EdgeX graphical user interface (GUI) provides an easy to use visual tool to monitor data passing through EdgeX services.
start-edgex-ui-server:
	@cd edgex-ui-go-3.1.0/cmd/edgex-ui-server && nohup ./edgex-ui-server -o > ~/edgex-foundry/edgex-native-build-3.1-napa/edgex-service-logs/edgex-ui-server/nohup.out 2>&1 &
	@echo "::: EdgeX ui-server is running... :::"

# the output from executing the command is redirected so it wont print any output
kill-edgex-ui-server:
	@ps aux | grep "edgex-ui-server" | grep -v grep | awk '{print $$2}' | xargs kill -15 > /dev/null 2>&1 &
	$ @echo "::: EdgeX ui-server has stopped working :::"

# ALl service status:
edgex-service-status:
	@echo "USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND"
	@ps aux | grep consul.agent | grep -v grep
	@ps aux | grep core-metadata | grep -v grep
	@ps aux | grep core-data | grep -v grep
	@ps aux | grep core-command | grep -v grep
	@ps aux | grep support-notifications | grep -v grep
	@ps aux | grep support-scheduler | grep -v grep
	@ps aux | grep app-service-configurable | grep -v grep
	@ps aux | grep device-virtual | grep -v grep
	@ps aux | grep edgex-ui-server | grep -v grep
	@ps aux | grep kuiperd | grep -v grep
