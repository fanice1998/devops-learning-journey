pipeline {
    agent any
    stages {
        stage('Build Docker') {
            steps {
                sh 'docker build --no-cache -t hello-devops:latest'
                echo '✅ Docker build 成功！'
              }
          }
      }
  }
