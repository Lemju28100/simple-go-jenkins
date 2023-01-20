pipeline {
    agent any
    stages {
        stage('Generate Modules') {
            steps {
                echo 'Building...'
                sh 'go mod init go-app'
            }
        }
        stage('Tidy Modules') {
            steps {
                echo 'Testing...'
                sh 'go mod tidy'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build -o ./dist/go-app'
                sh 'pwd'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
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
                    ansible-playbook -i ansible/hosts ansible/main.yml --extra-vars "env=staging"'
            }
        }
        stage('Deploy to Production') {
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
                    ansible-playbook -i ansible/hosts ansible/main.yml --extra-vars "env=production"'
            }
        }

    }

    post {
            always {
                cleanWs()
            }
        }
}