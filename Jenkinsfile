pipeline {
    agent any
    stages {
        stage('Generate Modules') {
            steps {
                echo 'Building...'
                sh 'cd $WORKSPACE/go-app && go mod init go-app'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'cd $WORKSPACE/go-app && go build -o ./dist/go-app'
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
                'sh cd $WORKSPACE/go-app/ansible && ansible-playbook -i hosts main.yml --extra-vars "env=staging"'
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

                'sh cd $WORKSPACE/go-app/ansible && ansible-playbook -i hosts main.yml --extra-vars "env=production"'
            }
        }
    }
}