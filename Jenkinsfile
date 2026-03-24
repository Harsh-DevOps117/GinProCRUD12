pipeline {
    agent any

    environment {
        APP_NAME = "gingo-app"
        DOCKER_IMAGE = "gingo-app:${BUILD_NUMBER}"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Harsh-DevOps117/GinProCRUD12.git'
            }
        }

        stage('Verify Go') {
            steps {
                sh 'go version'
            }
        }

        stage('Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build Binary') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE .'
            }
        }

        stage('Run Container (dev only)') {
            steps {
                sh '''
                docker stop gingo || true
                docker rm gingo || true
                docker run -d -p 8000:8000 --name gingo $DOCKER_IMAGE
                '''
            }
        }
    }
}
