pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh "'/root/go/bin/bee' pack"
        archiveArtifacts '*.tar.gz'
      }
    }
  }
}
