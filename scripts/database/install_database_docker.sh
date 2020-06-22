#!/bin/bash

docker run --name mariadb_qcobtc -e MYSQL_ROOT_PASSWORD=supersecretpassword -d mariadb/server:10.3