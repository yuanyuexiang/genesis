pipeline {
  agent any
  stages {
    stage('getsource') {
      steps {
        git 'https://github.com/yuanyuexiang/genesis.git'
      }
    }
    stage('build') {
      steps {
        sh "'/root/go/bin/bee' pack"
        archiveArtifacts 'secandJob.tar.gz'
      }
    }
  }
}
