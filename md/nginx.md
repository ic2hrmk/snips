# Configure NGINX with OpenSSL
## Generate self-signed certificates
To generate selfsigned key and it's certificate
~~~
sudo openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout /etc/ssl/private/nginx-selfsigned.key -out /etc/ssl/certs/nginx-selfsigned.crt
~~~

## Configure NGINX to use them
You have to put next configuration to NGINX conf. folder:
/etc/nginx/conf.d/my_config.conf
~~~
server {
    listen 80;
    server_name 127.0.0.1;
    return 301 https://$server_name$request_uri;
}
server {
    listen 443 ssl;
    server_name 127.0.0.1;

    ssl_certificate /etc/ssl/certs/nginx-selfsigned.crt;
    ssl_certificate_key /etc/ssl/private/nginx-selfsigned.key;

    location /api/ {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://localhost:8082;
    }

    location / {
        root /home/system/www/bweb/public/;
        try_files $uri $uri/ /index.html =404;
    }

    location ~* ^/(sys|collector|cashier|guard|)(/.+)?$ {
        root /home/system/www/bweb/public/;
        try_files $uri $uri/ /$1/index.html =404;
    }
}
~~~

