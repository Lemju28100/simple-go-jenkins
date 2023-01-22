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
                sh 'eval $(ssh-agent) && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_stage && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_staging && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_prod && \
                    ansible-playbook -i ansible/hosts ansible/main.yml --extra-vars "env=staging" --extra-vars "workdir=$WORKSPACE" --extra-vars "commit_id=$env.BUILD_NUMBER"'
            }
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

                sh 'eval $(ssh-agent) && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_stage && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_staging && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_prod && \
                    ssh-add /var/lib/jenkins/.ssh/id_rsa_production && \
                    ansible-playbook -i ansible/hosts ansible/main.yml --extra-vars "env=production" --extra-vars "workdir=$WORKSPACE" --extra-vars "commit_id=${env.BUILD_NUMBER}"'
            }
        }

    }

    post {
            always {
                cleanWs()
            }
        }
}