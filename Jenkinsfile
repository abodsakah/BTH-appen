pipeline {
    agent any
    tools {go '1.18'}
    stages {
        stage('Build') {
            steps {
                sh 'go build backend/api/main.go -o main'
            }
        }
        stage('Test') {
            // test the following files : backend/src/DB/exams_test.go, backend/src/DB/news_test.go, backend/src/DB/users_test.go
            steps {
                sh 'go test -p 1 ./...'
            } 
        }
        stage('Deploy') {
            steps {
                sh 'chmod +x deploy.sh && ./deploy.sh'
            }
        }
    }
}