pipeline {
    agent any
    stages {
        stage("Build") {
            steps {
                echo "Building..."
                sh "docker build -t 2464410/simple-go-jenkins:${env.BUILD_NUMBER} ."
            }
        }

        stage("Push to Docker Hub") {

            steps {
            withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    echo "Pushing to Docker Hub..."
                    sh "docker login -u ${USERNAME} -p ${PASSWORD}"
                    sh "docker push 2464410/simple-go-jenkins:${env.BUILD_NUMBER}"
            }
            }


        }

        stage('Deploy to Staging') {
            steps {
                echo 'Deploying to Staging...'
                ansiblePlaybook(
                    credentialsId: 'StagingPrivateKey',
                    inventory: 'ansible/hosts',
                    playbook: 'ansible/main.yml',
                    extraVars: [
                        "env": "staging",
                        "workdir": "${WORKSPACE}",
                        "commit_id": "${env.BUILD_NUMBER}"
                    ]
                )
        }
        stage('Deploy to Production') {

            when {
                branch 'main'
            }

            steps {
                echo 'Deploying to Production...'

                // This is where I want to use the input step to ask for a confirmation
                // before deploying to production
                input(message: 'Deploy to Production?', ok: 'Yes', parameters: [
                    string(defaultValue: 'No', description: 'Are you sure you want to deploy to production?', name: 'confirm')
                ])

                ansiblePlaybook(
                    credentialsId: 'ProductionPrivateKey',
                    inventory: 'ansible/hosts',
                    playbook: 'ansible/main.yml',
                    extraVars: [
                        "env": "production",
                        "workdir": "${WORKSPACE}",
                        "commit_id": "${env.BUILD_NUMBER}"
                    ]
                )
            }
        }

    }

    post {
            always {
                cleanWs()
            }
        }
}