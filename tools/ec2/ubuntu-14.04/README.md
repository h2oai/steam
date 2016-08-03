# CI Steam Host 

This readme details how the "daily" CI steam host has been set up at https://steam.h2o.ai/.

Use `ssh -i ~/.ec2/keys/steam-engine.pem ubuntu@52.91.182.20` to ssh into that EC2 box.

The `bounce-steam.sh` script in this git directory is invoked by crontab to periodically pull the latest master build and bounce steam.

## Ubuntu 14.04 installation notes

Install postgres:

```
sudo apt update
sudo apt upgrade
sudo apt install postgresql
```

Set up `steam` user/role:

```
sudo -i -u postgres 
createuser --interactive -P steam
# Enter password for new role: pa$$word
# Enter it again: pa$$word
# Shall the new role be a superuser? (y/n) n
# Shall the new role be allowed to create databases? (y/n) y
# Shall the new role be allowed to create more new roles? (y/n) n
createdb steam steam
cd path/to/steam/var/master/scripts
./create-database.sh
exit
```

Set up `.pgpass` for postgres:

- Edit `~/.pgpass` and append the line `*:*:*:steam:pa$$word`
- Run `chmod 600 ~/.pgpass` 

## Starting steam

See `bounce-steam.sh` 
