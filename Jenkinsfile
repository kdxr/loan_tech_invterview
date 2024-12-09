pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.23.4'
    }
    triggers {
        pollSCM 'H/10 * * * *'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        GIT_REPO = 'https://github.com/kdxr/loan_tech_invterview.git' // URL ของ GitHub repo
        BUILD_DIR = 'backend/manage_customers/' // โฟลเดอร์ที่ต้องการ build
    }
    stages {
        stage('Checkout') {
            steps {
                echo 'Cloning repository...'
                git branch: 'main', url: "${GIT_REPO}" // เปลี่ยน branch ได้ตามต้องการ
            }
        }

        // stage("unit-test") {
        //     steps {
        //         echo 'UNIT TEST EXECUTION STARTED'
        //         sh 'make unit-tests'
        //     }
        // }
        // stage("functional-test") {
        //     steps {
        //         echo 'FUNCTIONAL TEST EXECUTION STARTED'
        //         sh 'make functional-tests'
        //     }
        // }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                dir(BUILD_DIR) {
                    sh 'go get ./...'
                    // sh 'docker build . -t shadowshotx/product-go-micro'
                    sh 'go build ./...'
                }
            }
        }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         echo 'Delivery...'
        //         sh ''
        //         sh 'go run ./...'
        //     }
        // }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push shadowshotx/product-go-micro'
        //         }
        //     }
        // }
    }
}
