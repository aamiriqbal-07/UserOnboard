if [ $# -eq 0 ]; then
    echo "No port specified. Using default port 8080"
    PORT=8080
else
    PORT=$1
fi

export HOST_PORT=$PORT
export APP_PORT=$PORT

echo "Stopping any existing containers..."
docker-compose down

echo "Starting application on port $PORT..."
docker-compose up --build