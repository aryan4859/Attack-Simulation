FROM php:7.4-apache

COPY public/ /var/www/html/
COPY flag.txt /flag.txt

RUN chmod o+r /var/log/apache2/access.log

EXPOSE 80
