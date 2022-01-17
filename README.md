# automatic-trading-system

## Basic Information about System

### System Web URL
[URL](https://serpedious.link/)

### System details(Infra Architecture Design Image, DB Design Image)
[System description](https://checker-crime-252.notion.site/Automatic-Trading-System-eec52c8300a6425c814c8a148ef5206e)


## Help

### Exeute App(when you execute this app in your local environment, you need to excute items below)
- create files listed in .gitignore
- command "docker-compose up"

### Enter Postgres Container(Local)
docker-compose exec postgres psql -U postgres test_db

### Access to Production Postgres(Production)
psql -h {rds-endpoint} -U {username} {dbname}

### Costly resources
- CloudFront
- RDS
- ECS Service
- loadbalancer(Elastic IP)
- EC2(Verification)


