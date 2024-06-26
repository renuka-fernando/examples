# Use the official Node.js image
FROM node:16.13-alpine
RUN apk update && apk upgrade

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the port the app runs on
EXPOSE 3000

USER 10014

# Command to run the application
CMD ["node", "server.js"]
