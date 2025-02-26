# Use the official Nginx image
FROM nginx:alpine

# Copy the HTML file to the Nginx's default public folder
COPY index.html /usr/share/nginx/html/index.html

# Copy any additional files (e.g., images or memory dump files)
COPY 866d8156c467976aceda8e20f6bc7b83.jpg /usr/share/nginx/html/866d8156c467976aceda8e20f6bc7b83.jpg 

# Expose port 80 to the outside world
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
