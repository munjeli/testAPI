#!/bin/bash
# Make a dir for the secret
# Can I mention again how uncomfortable I am
# with this? Where's my vault???
mkdir /etc/count/private
mv /home/ubuntu/secret.txt /etc/count/private/

# Make a dir to put certs.
mkdir /etc/count /etc/count/tls
mv /home/ubuntu/server.key /etc/count/tls/
mv /home/ubuntu/server.crt /etc/count/tls/

# Permissions.
chmod 400 /etc/count/tls/server.key
chmod 400 /etc/count/tls/server.crt
chmod 500 /etc/count/tls/
chmod 500 /etc/count/secret
chmod 400 /etc/count/secret/secret.txt