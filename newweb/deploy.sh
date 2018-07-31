#! /bin/sh

kill -9 $(pgrep aven)
cd /home/wys/aven/
git pull https://github.com/itman312/aven.git
./aven &