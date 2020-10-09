pipeline {
        environment {
            registry = "gits5622/eshop"
            registryCredential = 'docker-hub'
            dockerImage = ''
            }
        agent any
        stages {
                stage('Cloning our Git') {
                    steps {
                    git 'https://github.com/gitx5622/e-shop_back.git'
                    }
                }
                stage('Building our image') {
                    steps{
                        script {
                        dockerImage = docker.build registry
                        }
                    }
                }
                 stage('Deploy our image') {
                                    steps{
                                        script {
                                        docker.withRegistry( '', registryCredential ) {
                                        dockerImage.push()
                                            }
                                        }
                                    }
                                }

                stage ('Running tha Application'){
                    steps{
                        sh "docker compose up"
                    }
                }

    }
}