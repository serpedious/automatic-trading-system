# Automatic-Trading-System

## BASIC INFORMATION ABOUT THIS APPLICATION

### APPLICATION WEB URL
~~[URL](https://serpedious.link/)(Locked by Basic Authentication, ask to developer when you access)~~ 
(now, application is stopped because of generating expense)

### APPLICATION DETAILS(Including Design Image of Infrastructure Architecture, DB Design Image)
[System description](https://checker-crime-252.notion.site/Automatic-Trading-System-eec52c8300a6425c814c8a148ef5206e)


## HELP

### default guest user
- email: guest@gmail.com
- password: password

### Exeute App(when you execute this app in your local environment, you need to excute items below)
- create .env and set proper parameters
- command execute "docker-compose up"

### Enter Postgres Container(Local)
docker-compose exec postgres psql -U postgres test_db

### Access to Production Postgres(Production)
psql --host={rds-endpoint} --port=5432 --username={username} --password --dbname={dbname}

### Costly resources on AWS
- CloudFront
- RDS
- ECS Service
- Elastic Loadbalancer
- EC2


