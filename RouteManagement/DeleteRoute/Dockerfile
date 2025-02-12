# Base Image
FROM node:16

# Set the working directory
WORKDIR /app

# Copy files
COPY package.json package-lock.json ./
RUN npm install

COPY . .

# Expose port
EXPOSE 6004

# Start the app
CMD ["node", "app.js"]
