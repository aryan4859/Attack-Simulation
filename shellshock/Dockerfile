FROM php:7.4-apache

# Install required dependencies
RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    && docker-php-ext-configure gd \
    && docker-php-ext-install gd pdo pdo_mysql

# Enable Apache mod_rewrite
RUN a2enmod rewrite

# Allow Apache to execute .phtml and .php5 as PHP
RUN echo "<FilesMatch \"\.(php|php5|phtml)$\">" >> /etc/apache2/apache2.conf && \
    echo "SetHandler application/x-httpd-php" >> /etc/apache2/apache2.conf && \
    echo "</FilesMatch>" >> /etc/apache2/apache2.conf

# Set up the working directory
WORKDIR /var/www/html

# Copy PHP files into the container
COPY ./src/ /var/www/html/

# Create the uploads directory with correct permissions
RUN mkdir -p /var/www/html/uploads && chmod -R 777 /var/www/html/uploads

# Expose port 80
EXPOSE 80

# Restart Apache to apply changes
CMD ["apache2-foreground"]
