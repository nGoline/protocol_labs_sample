# `/scripts`

Scripts to ease development and deployment

## `/database`

Scripts related to database

### Install Database Docker

To install MariaDB on a docker container be sure to have docker installed and then:

```bash
./install_database_docker.sh
```

The database will be accessible on your container IP:3306. To get your container IP just run:

```bash
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mariadb_qcobtc
```

Be sure to export the contents of [database.conf](database/database.conf) file and to update this line `SQL_ADDRESS=<DATABSE_INTERNAL_IP>:3306` to the IP acquired on the previous step.

### Initialize Databse

In order to run the workers it's necessary to populate the database using the [database.sql](database/database.sql) to create the tables and to create the initial registries needed to fetch data from the exchanges and blockstream.

to acces the container to run this creation from inside it run:

```bash
docker exec -it bash /bin/sh -c "[ -e /bin/bash ] && /bin/bash || /bin/sh"
mysql -p
```

And use the default password `supersecretpassword`.
