#! /bin/bash
git pull origin master
#tar -xzvf beepkg.tar.gz
rm -rf conf/app.conf
cp conf/app.conf.bat conf/app.conf
supervisorctl restart wxapi
