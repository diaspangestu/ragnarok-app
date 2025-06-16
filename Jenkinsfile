pipeline {
  agent any

  environment {
    DOCKER_IMAGE = 'npangestu/test-deploy:dev'
    KUBE_CONTEXT = 'minikube'
  }

  stages {
    stage('Build Docker Image') {
      steps {
        sh 'docker build -t $ DOCKER_IMAGE .'
      }
    }

    stage('Push to Docker Hub') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
          sh '''
            echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
            docker push $DOCKER_IMAGE
          '''
        }
      }
    }

    stage('Deploy to Minikube') {
      steps {
        sh 'kubectl config use-context $KUBE_CONTEXT'
        sh 'kubectl apply -f .'  // assumes your YAMLs are in ./k8s
      }
    }
  }
}
