pipeline {

  agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
  }

  stages {
    stage('Deploy App to Kubernetes') {     
      steps {
        container('kubectl') {
          withCredentials([file(credentialsId: 'mykubeconfig', variable: 'KUBECONFIG')]) {
            sh 'cat $KUBECONFIG > ./kube/config'
            // sh 'sed -i "s/<TAG>/${BUILD_NUMBER}/" deploy.yaml'
            sh 'kubectl get all'
          }
        }
      }
    }
  
  }
}