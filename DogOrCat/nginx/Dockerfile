FROM nginx:1.19.1-alpine

COPY . /home/nginx

RUN cd /home/nginx; \
    mv docker-entrypoint.sh /usr/local/bin; \
    chmod +x /usr/local/bin/docker-entrypoint.sh; \
    mv nginx.conf /etc/nginx/nginx.conf; \
    rm /etc/nginx/conf.d/default.conf;

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]