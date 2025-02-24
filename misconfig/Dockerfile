# Use PHP with Apache image
FROM php:7.2-apache

# Enable mod_rewrite for URL rewriting (if needed)
RUN a2enmod rewrite

# Copy the application files into the container
COPY ./src/ /var/www/html/

# Expose the port to access the web application
EXPOSE 80
