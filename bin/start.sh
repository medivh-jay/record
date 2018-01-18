#!/bin/sh

port=9999
host=10.26.4.152:27017
username=tracker
password="e5d\$e(Gs%epN3nDb"

echo listen port is : ${port}
echo mongo url is : ${host}
rm -f ./nohup.out
kill -9 $(pidof ./record_run)

cp ./record ./record_run
nohup ./record_run -port ${port} -host ${host} -username ${username} -password ${password} &