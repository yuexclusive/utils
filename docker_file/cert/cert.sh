docker build . -t certbot

docker run -it --rm --name certbot -v ${PWD}:/letsencrypt -v ${PWD}/certs:/etc/letsencrypt certbot bash

certbot certonly --webroot

docker run -itd --rm --name nginx -v ${PWD}/nginx.conf:/etc/nginx/nginx.conf -v ${PWD}:/letsencrypt -v ${PWD}/certs:/etc/letsencrypt -p 80:80 -p 443:443 nginx

## Renewal

# To do a dry run of cert renewal:

certbot renew --dry-run

# Reload our NGINX web server if the certs change:

docker exec -it nginx sh -c "nginx -s reload"


# Checkout the Certbot [docs](https://certbot.eff.org/instructions) for more details