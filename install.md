reddis https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-20-04-id

sudo apt update
sudo apt install redis-server
sudo nano /etc/redis/redis.conf
supervised systemd
sudo systemctl restart redis.service
sudo systemctl status redis