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

    stage('Deploy App to Kubernetes') {     
      steps {
        container('kubectl') {
          withCredentials([file(credentialsId: 'mykubeconfig', variable: 'KUBECONFIG', 'serverUrl': 'http://172.20.0.2')]) {
            sh 'sed -i "s/<TAG>/${BUILD_NUMBER}/" deploy.yaml'
            sh 'kubectl config set-cluster mylocal '
            sh 'kubectl get all -o wide'
            // sh 'kubectl apply -f deploy.yaml'
          }
        }
      }
    }
  
  }
}