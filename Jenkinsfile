properties([pipelineTriggers([githubPush()])])

pipeline {
    agent any
    stages {
        stage ('Checkout') {
            checkout scm
        }
        stage('Lint HTML') {
            steps {
                sh 'tidy -q -e *.html'
            }
        }
        stage('Lint Dockerfile'){
            steps {
                sh 'hadolint Dockerfile'
            }
        }
        stage('Build Container Image') {
            sh 'docker image build -t mansong/go-image-classifier:latest .'
        }
        stage ('Scan Container Image') {
            aquaMicroscanner imageName: 'mansong/go-image-classifier:latest', notCompliesCmd: 'exit 1', onDisallowed: 'fail', outputFormat: 'html'
        }
        stage('Push Container Image') {
            withCredentials([
                usernamePassword(credentialsId: 'docker-credentials',
                        usernameVariable: 'USERNAME',
                        passwordVariable: 'PASSWORD')]) {
                            sh 'docker login -p "${PASSWORD}" -u "${USERNAME}"'
                            sh 'docker image push ${USERNAME}/go-image-classifier:latest'
                        }
        }
        stage ('Analysis') {
            sh 'docker run -i kubesec/kubesec:c5a4ff5 scan /dev/stdin < deployment.yaml | jq --exit-status '.score > 10' >/dev/null'
        }
        stage('Deploy') {
            withCredentials([
                file(credentialsId: 'kube-config',
                        variable: 'KUBECONFIG')]) {
                            sh 'kubectl apply -f deployment.yaml'
                    }
                }
            }
        stage('Post Deploy Test') {
            steps {
                echo 'Testing Deployment'
                script {
                    String webSite = 'https://mansong-jenkins-udacity.s3.us-east-1.amazonaws.com/index.html'
                    def returnCode = sh(returnStdout: true, script: "curl -sLI -o /dev/null -w '%{http_code}' ${webSite}").trim() as Integer
                    if (returnCode != 200) {
                        currentBuild.result = 'FAILURE'
                    }
                }
            }
        }
    }
    /* Cleanup workspace */
    post {
       always {
           deleteDir()
       }
    }
}