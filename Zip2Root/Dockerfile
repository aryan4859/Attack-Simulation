# Use PHP 5.6 with Apache
FROM php:5.6-apache

LABEL MAINTAINER="aryanstha <aryanstha4859@gmail.com>"

# Install mysqli extension
RUN docker-php-ext-install mysqli

# Copy your web files into the container
COPY www/ /var/www/html/

# Set proper permissions
RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html

# Enable mod_rewrite (optional, if needed)
RUN a2enmod rewrite

# Configure Apache to use index.php by default
RUN echo '<Directory /var/www/html>\n\
    Options -Indexes +FollowSymLinks\n\
    DirectoryIndex index.php index.html\n\
    AllowOverride All\n\
    Require all granted\n\
</Directory>' > /etc/apache2/conf-available/zip2root.conf \
    && a2enconf zip2root

# Expose port 80
EXPOSE 80

# Start Apache in foreground
CMD ["apache2-foreground"]
