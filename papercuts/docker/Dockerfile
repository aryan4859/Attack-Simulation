FROM php:7.4-apache

# Copy the entire directory to the container
COPY public/ /var/www/html/
COPY src/ /var/www/html/src/   
COPY flag.txt /flag.txt

# Make sure Apache logs are world-readable
RUN chmod o+r /var/log/apache2/access.log
