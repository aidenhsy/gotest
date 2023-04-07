pipeline {
    agent any
    environment {
        SSH_USER = 'ecs-assist-user'
        REMOTE_HOST = '106.14.158.113'
    }
    stages {
        stage('Execute command on remote host') {
            steps {
                sshagent(credentials: ['sshkey']) {
                    sh 'ssh-keyscan 106.14.158.113 >> ~/.ssh/known_hosts'
                    sh """
                        ssh $SSH_USER@$REMOTE_HOST '
                        cd /var/www/tshelpers
                        git pull https://gitee.com/aidenhsy/tshelpers.git
                        ls
                        '
                    """
                }
            }
        }
    }
}