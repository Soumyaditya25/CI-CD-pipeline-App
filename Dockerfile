# Use a base image from a trusted source
FROM golang:1.17



# Install any additional packages or dependencies your application needs

# Set the working directory inside the container
WORKDIR /app

# Copy your application files into the container
COPY login /app


RUN go build -o app
# Expose any necessary ports
EXPOSE 3001

# Define the command to run your application
CMD ["./app"]
