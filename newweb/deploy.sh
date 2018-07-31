#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull https://github.com/itman312/newweb.git
cd webserver/
./webserver &