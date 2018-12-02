#! /bin/bash
git checkout .
git pull origin master

rm -rf conf/app.conf
cp conf/app.conf.bat conf/app.conf

chmod -R 755 reload.sh
chmod -R 755 restart.sh
