# Connect to banjo DB

~~~
$ sudo -i -u postgres
$ psql
~~~

# Get list of DB

~~~
postgres=# \list
~~~

# Get list of Tables

~~~
postgres=# select * from information_schema.tables where table_schema = 'public';
~~~

# Allow access from remote hosts

Path to config files
~~~
/etc/postgresql/9.4/main
~~~

## Configuring postgresql.conf

~~~
listen_addresses = 'localhost' --> listen_addresses = '*'
~~~

## Configuring pg_hba.conf

~~~
host    all             all              0.0.0.0/0                       md5
host    all             all              ::/0                            md5
~~~

# PGDUMP
~~~
PGPASSWORD=123456 /usr/bin/pg_dump --host banjo.loc --port 5432 --username "system" --role "system" --no-password --format custom --blobs --verbose --file "/home/ic2h/banjo_dev.backup" "banjo_dev"
~~~