pipeline {
    agent any
    tools {go '1.19'}
    stages {
        stage('Build') {
            steps {
                sh 'go build backend/src/main.go -o main'
            }
        }
        stage('Test') {
            // test the following files : backend/src/DB/exams_test.go, backend/src/DB/news_test.go, backend/src/DB/users_test.go
            steps {
                sh 'go test backend/src/DB/exams_test.go'
                sh 'go test backend/src/DB/news_test.go'
                sh 'go test backend/src/DB/users_test.go'
            } 
        }
        stage('Deploy') {
            steps {
                sh 'chmod +x deploy.sh && ./deploy.sh'
            }
        }
    }
}