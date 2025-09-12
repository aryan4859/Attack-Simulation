# Use PHP 5.6 with Apache
FROM php:5.6-apache

LABEL MAINTAINER="aryanstha <aryanstha4859@gmail.com>"

# Install mysqli extension
RUN docker-php-ext-install mysqli

# Copy web files
COPY www/ /var/www/html/

# Set permissions
RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html

# Enable mod_rewrite
RUN a2enmod rewrite

# Apache configuration to use index.php
RUN echo '<Directory /var/www/html>\n\
    Options -Indexes +FollowSymLinks\n\
    DirectoryIndex index.php index.html\n\
    AllowOverride All\n\
    Require all granted\n\
</Directory>' > /etc/apache2/conf-available/zip2root.conf \
    && a2enconf zip2root

# Copy wait-for-db script
COPY wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

EXPOSE 80

# Start Apache after waiting for DB
CMD ["/wait-for-db.sh", "db", "apache2-foreground"]
