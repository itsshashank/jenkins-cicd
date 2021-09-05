pipeline {

  agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
  }

  node {
    stage('List pods') {
      withKubeConfig([credentialsId: '<credential-id>',
                      caCertificate: '<ca-certificate>',
                      serverUrl: '<api-server-address>',
                      contextName: '<context-name>',
                      clusterName: '<cluster-name>',
                      namespace: '<namespace>'
                      ]) {
        sh 'kubectl get pods'
      }
    }
  }
}