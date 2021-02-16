#!/bin/bash
#
# @Author : Randy Ardiansyah K
# Supported By Soleh
# chkconfig: 35 95 05
# description: Scheduler Proses Harian IBSGen2

# Run at startup: sudo chkconfig <service name> on

# Load functions from library
. /etc/init.d/functions

scriptFile=$(readlink -fn $(type -p $0))                # the absolute, dereferenced path of this script file
scriptDir=$(dirname $scriptFile)                        # absolute path of the script directory

# Configure variables below :
app="CronDashboardIBSGen2"
serviceLogName="service"
serviceName="CronDashboardIBSGen2"
homedir="/home/CronDashboardIBSGen2"

# STATIC VARIABLE
serviceUser="root"                                      # OS user name for the service
serviceGroup="root"                                     # OS group name for the service
etcInitDFile="/etc/init.d/$serviceName"       # symlink to this script from /etc/init.d
rcFileBaseName="rc$app"                # basename of the "rc" symlink file for this script
rcFileName="/usr/local/sbin/$rcFileBaseName"     # full path of the "rc" symlink file for this script

# Get all arguments
args=""
for i in "${@:2}"
do
   args="$args$i "
done

#echo "result:$args"

# Start the service
run() {
  echo -n $"Starting $app:"
  cd $homedir
  ./$app $args > $homedir/log/$serviceLogName.log 2> $homedir/log/$serviceLogName.err < /dev/null &

  sleep 1

  status $app > /dev/null
  # If application is running
  if [[ $? -eq 0 ]]; then
    # Store PID in lock file
    echo $! > /var/lock/subsys/$app
    success
    echo
  else
    failure
    echo
  fi
}

# Start the service
start() {
  status $app > /dev/null
  # If application is running
  if [[ $? -eq 0 ]]; then
    status $app
  else
    run
  fi
}

# Restart the service
stop() {
  echo -n "Stopping $app: "
  killproc $app
  rm -f /var/lock/subsys/$app
  echo
}

# Reload the service
reload() {
  status $app > /dev/null
  # If application is running
  if [[ $? -eq 0 ]]; then
    echo -n $"Reloading $app:"
    kill -HUP `pidof $app`
    sleep 1
    status $app > /dev/null
    # If application is running
    if [[ $? -eq 0 ]]; then
      success
      echo
    else
      failure
      echo
    fi
  else
    run
  fi
}

function installService {
   getent group $serviceGroup >/dev/null 2>&1
   if [ $? -ne 0 ]; then
      echo Creating group $serviceGroup
      groupadd -r $serviceGroup || return 1
      fi
   id -u $serviceUser >/dev/null 2>&1

   if [ $? -ne 0 ]; then
      echo Creating user $serviceUser
      useradd -r -c "user for $app service" -g $serviceGroup -G users -d $homedir $serviceUser
      fi
   ln -s $scriptFile $rcFileName
   ln -s $scriptFile $etcInitDFile

   chkconfig --add $serviceName  || return 1
   echo $serviceName installed.
   echo You may now use $app  to call this script.
   return 0;
}

function uninstallService {
   chkconfig --del $app
   rm -f $rcFileName
   rm -f $etcInitDFile
   echo $serviceName uninstalled.
   return 0;
}

# Main logic
case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  install)
    installService
    ;;
  uninstall)
    uninstallService
    ;;
  status)
    status $app
    ;;
  restart)
    stop
    sleep 1
    start
    ;;
  reload)
    reload
    ;;
  *)
    echo $"Usage: $0 {start|stop|restart|reload|status|install}"
    exit 1
esac
exit 0
