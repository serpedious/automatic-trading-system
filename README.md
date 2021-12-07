# automatic-trading-system

### you need to excute items below
- make files listed in .gitignore
- command "docker-compose build"
- command "docker-compose up"


### when you enter into postgres container
docker-compose exec postgres psql -U postgres test_db

### when you access postgres in prod via ec2
psql -h {rds-endpoint} -U {username} {dbname}

### when you access postgres in prod via ec2
example-user
email: guest@gmail.com
password: password

### activate aws resource when you deploy
- CloudFront
- RDS
- ECS Service
- loadbalancer(Elastic IP)
- EC2(Verification)

above resources cost so much. In usual I stop them.


### Trouble shooting
- It is possible CloudFront ID would change.
so you have to check ID specified at lambda function

- It is possible it is different from correct S3 domain
in CloudFront

### when you add env variable, you need to add it at the file in S3
- golang env variable