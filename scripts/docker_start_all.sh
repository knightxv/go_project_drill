#!/usr/bin/env bash
#fixme This script is the total startup script
#fixme The full name of the shell script that needs to be started is placed in the need_to_start_server_shell array

#fixme Put the shell script name here

source ./style_info.cfg
source ./path_info.cfg
source ./function.sh

#fixme The 10 second delay to start the project is for the docker-compose one-click to start openIM when the infrastructure dependencies are not started

sleep 10
time=`date +"%Y-%m-%d %H:%M:%S"`
echo "==========================================================">>../logs/UserScore.log 2>&1 &
echo "==========================================================">>../logs/UserScore.log 2>&1 &
echo "==========================================================">>../logs/UserScore.log 2>&1 &
echo "==========server start time:${time}===========">>../logs/UserScore.log 2>&1 &
echo "==========================================================">>../logs/UserScore.log 2>&1 &
echo "==========================================================">>../logs/UserScore.log 2>&1 &
echo "==========================================================">>../logs/UserScore.log 2>&1 &


#service filename
service_filename=(
  #rpc
  user_score
  chain_up
)

#service config port name
service_port_name=(
  userScorePort
  chainUpPort
)

service_prometheus_port_name=(
  userScorePrometheusPort
  chainUpPrometheusPort
)

for ((i = 0; i < ${#service_filename[*]}; i++)); do

  #Check whether the service exists
  service_name="ps -aux |grep -w ${service_filename[$i]} |grep -v grep"
  count="${service_name}| wc -l"

  if [ $(eval ${count}) -gt 0 ]; then
    pid="${service_name}| awk '{print \$2}'"
    echo  "${service_filename[$i]} service has been started,pid:$(eval $pid)"
    echo  "killing the service ${service_filename[$i]} pid:$(eval $pid)"
    #kill the service that existed
    kill -9 $(eval $pid)
    sleep 0.5
  fi
  cd ../bin

  chmod +x ./*
  #Get the rpc port in the configuration file
  portList=$(cat $config_path | grep ${service_port_name[$i]} | awk -F '[:]' '{print $NF}')
  list_to_string ${portList}
  service_ports=($ports_array)

  portList2=$(cat $config_path | grep ${service_prometheus_port_name[$i]} | awk -F '[:]' '{print $NF}')
  list_to_string $portList2
  prome_ports=($ports_array)
  #Start related rpc services based on the number of ports
  for ((j = 0; j < ${#service_ports[*]}; j++)); do
    #Start the service in the background
    cmd="./${service_filename[$i]} -port ${service_ports[$j]} -prometheus_port ${prome_ports[$j]}"
    if [ $i -eq 0 -o $i -eq 1 ]; then
      cmd="./${service_filename[$i]} -port ${service_ports[$j]}"
    fi
    echo $cmd
    nohup $cmd >>../logs/UserScore.log 2>&1 &
    sleep 1
    pid="netstat -ntlp|grep $j |awk '{printf \$7}'|cut -d/ -f1"
    echo -e "${GREEN_PREFIX}${service_filename[$i]} start success,port number:${service_ports[$j]} pid:$(eval $pid)$COLOR_SUFFIX"
  done
done

sleep 1
# nohup ../bin/chain_event_listener >> ../logs/chain_event_listener.log 2>&1 & 

#fixme prevents the openIM service exit after execution in the docker container
tail -f /dev/null
