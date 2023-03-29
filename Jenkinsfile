pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t myproject .'
            }
        }
        
        stage('Test') {
            steps {
                sh 'docker run --rm myproject python manage.py test'
                sh 'docker run --rm myproject pytest'
            }
        }
        
        stage('Security Scan') {
            steps {
                script {
                    def zap_commands = readFile 'zap_commands.txt'
                    sh "docker run -d -p 8080:8080 --name zap owasp/zap2docker-stable"
                    sh "docker exec zap zap-cli -p 8080 status -t 120 && echo 'ZAP is up!'"
                    sh "docker exec zap $zap_commands"
                    sh "docker stop zap"
                }
            }
        }
                
        stage('Load Test') {
            steps {
                script {
                    def jmeter_testplan = readFile 'testplan.jmx'
                    sh "docker run --rm -v $(pwd):/mnt -w /mnt -e HEAP='-Xms1g -Xmx1g' jmeter -n -t /mnt/${jmeter_testplan} -l /mnt/results.jtl"
                }
            }
        }
                

        stage('Deploy Microservices') {
                steps {
                    script {
                        for (service in ["admin", "logger", "stores", "analytics", "main_network"]) {
                            dir(service) {
                                sh 'docker-compose up -d'
                                sh 'docker-compose run --rm owasp zap-baseline.py -t http://localhost:8002'
                                sh 'docker-compose run --rm jmeter -n -t /jmeter/scripts/register_order.jmx -Jbase_url=http://localhost:8002/service/store/concurrent'
                            }
                        }
                    }
                }
            }
        }
    }
