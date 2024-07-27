#!/bin/bash
pids=$(ps aux | grep -E 'consul|core-metadata|core-data|core-command|support-notifications|support-scheduler|app-service-configurable|edgex-ui-server|kuiperd|device-virtual' | grep -v grep | awk '{print $2}' | tr '\n' ',' | sed 's/,$//')
if [ -n "$pids" ]; then
    htop -p $pids
else
    echo "No matching EdgeX services found."
fi
