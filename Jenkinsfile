pipeline {
    agent any
    tools {go '1.18'}
    stages {
        stage('Build') {
            steps {
                sh 'go build backend/api/main.go -o main'
            }
        }
        stage('Deploy') {
            steps {
                sh 'chmod +x deploy.sh && ./deploy.sh'
            }
        }
    }
}