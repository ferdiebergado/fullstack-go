FROM nginx:1.27.1-alpine3.20
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./public /app
