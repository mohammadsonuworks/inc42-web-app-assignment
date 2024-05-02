# Use the official Node.js image
FROM node:20-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json .

# Install dependencies
RUN npm install

# Copy source-code
COPY . .

# Expose the default Next.js development port
EXPOSE 3000

# Set the default command to start the Next.js development server
CMD ["npm", "run", "dev"]
