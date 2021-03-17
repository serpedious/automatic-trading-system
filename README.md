# automatic-trading-system

### you need to excute items below
- make files listed in .gitignore
- command "docker-compose build"
- command "docker-compose up"


### when you enter into postgres container
docker-compose exec db psql -U postgres test_db

### when you access postgres in prod via ec2
psql -h {rds-endpoint} -U {username} {dbname}

### activate aws resource when you deploy
- CloudFront
- RDS
- ECS Service

above resources cost so much. In usual I stop them.


### Trouble shooting
- It is possible CloudFront ID would change.
so you have to check ID specified at lambda function

- It is possible it is different from correct S3 domain
in CloudFront