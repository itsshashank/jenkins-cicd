pipeline {

  agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
  }

  node {
    stage('List pods') {
      withKubeConfig([credentialsId: 'mykubeconfig', variable: 'KUBECONFIG']) {
        sh 'cat $KUBECONFIG > ~/.kube/config'
        sh 'kubectl get all'
      }
    }
  }
}