pipeline {

  agent {
    kubernetes {
      yamlFile 'builder.yaml'
    }
  }

  stages {

    // stage('Kaniko Build & Push Image') {
    //   steps {
    //     container('kaniko') {
    //       script {
    //         sh '''
    //         /kaniko/executor --force --dockerfile `pwd`/Dockerfile \
    //                          --context `pwd` \
    //                          --destination=itsshashank/gin-sample:${BUILD_NUMBER}
    //         '''
    //       }
    //     }
    //   }
    // }

    stage('List pods') {
    withKubeConfig([credentialsId: 'mykubeconfig'
                    ]) {
      sh 'kubectl get pods'
    }
  }
  
  }
}