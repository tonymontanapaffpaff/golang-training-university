pipeline {
    agent any
    tools {
        go 'go1.16'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u -mod=readonly golang.org/x/lint/golint'
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build ./cmd/main.go'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vet'
                    sh 'go vet ./...'
                    echo 'Running fmt'
                    sh 'go fmt ./...'
                    echo 'Running linter'
                    sh 'golint ./...'
                    echo 'Running test'
                    sh 'make test'
                }
            }
        }
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"

        }
    }
}