pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {
                sh 'chmod +x deploy.sh && ./deploy.sh'
            }
        }
    }
}