#!/usr/bin/env bash

mongod -f /etc/mongod.conf --quiet &
sleep 30
mongo localhost:27017 /etc/users.js
sleep 10
mongod --shutdown
mongod -f /etc/mongod.conf
