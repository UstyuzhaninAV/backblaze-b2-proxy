#### nginx b2 reverse proxy

this is for setting up nginx to be a caching reverse proxy for backblaze b2 __private__ bucket

it will maintain the Authorization header for backend requests in the vhost config and reload nginx

prereqs:

* application key, needs read access to target bucket (https://secure.backblaze.com/app_keys.htm)
* bucketId for target bucket (https://secure.backblaze.com/b2_buckets.htm)

setup:

1. copy b2.conf to your nginx vhosts folder
2. copy .env.example to .env and fill in your b2 info + nginx b2.conf vhost path
3. setup executable on cronjob to run every 5-6 days. key will expire after 1 week

in b2.conf, you will need to update some things:

* on `proxy_cache_path` line - change your cache size from `18g` to whatever you want
* on `proxy_pass` line - change "myprivatebucket" to your private bucket name

and of course, make sure `/var/cache/nginx/` exists and is writable by nginx

crontab example: `00 00 * * */5   root    /root/b2key`
