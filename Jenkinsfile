pipeline {

  agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
  }

  node {
    stage('List pods') {
      withKubeConfig([credentialsId: 'mykubeconfig'
                      ]) {
        sh 'kubectl get all'
      }
    }
  }
}