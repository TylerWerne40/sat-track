# Powershell stuff

function start_or_run {
    docker inspect peril_rabbitmq | Out-Null 2>&1

    if ( $? ) {
        Write-Host "Starting Peril RabbitMQ container..."
        docker start peril_rabbitmq
    }
    else {
        Write-Host "Peril RabbitMQ container not found, creating a new one..."
        docker run -d --name peril_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
    }
}

switch ($args[0]) {
    "start" {
        start_or_run
        }
    "stop" {
        Write-Host "Stopping Peril RabbitMQ container..."
        docker stop peril_rabbitmq
        }
    "logs" {
        Write-Host "Fetching logs for Peril RabbitMQ container..."
        docker logs -f peril_rabbitmq
        }
    Default {
        Write-Host "Usage: $0 {start|stop|logs}"
        exit 1
    }
}

