### APPLICATION WEB URL
~~[URL](https://serpedious.link/)(Locked by Basic Authentication, ask to developer when you access)~~ 
(now, application is stopped because of generating expense)

## System Description

### What Kind of Application?
- It enables you to maintain and promote transactions of trading about cryptocurrency by using external api

### What this application can do?
- Get notification about the good timing we should buy and sell bitcoin at Slack and dashboard
- Confirm and show transcations data which I have ever done before
- Download transaction history data with CSV file
- Show Bitcoin candle information by “Second”, “Miniute” and “Hour”
- Trade cryptocurrency from application GUI(BTC, XRP, ETH, XLM, MONA)
- Confirm my all assets data including cryptocurrency and JPY in my Bitflyer account
- Take memos that I need to keep in mind during trading
- Confirm and show data of history which was deposited and withdrawed JPY between Bitflyer account and my ordinary banking account
- Diagnosis situation about trader mental by giving some simple questions and print result as PDF
  
### Detail feature
- Send Notification mail when somebody create new user
- Create User
- Authenticate user
- Modify password of user
- Login by Guest user
- Validation(form)

### Used Technology
- Front: Nuxt.js+Vuetify
- BackEnd: Golang(Web Server: Chi, Indecator Library: Talib)+Python(Flask)
- Infrastructure(Cloud): AWS(Fargate, ECR, API Gateway, Lambda, ELB, S3, CloudFront, RDS, CodePipeline, CodeCommit, CodeBuild, CodeDeploy, Cloud Watch, Route53, ACM)
- Database: PostgresSQL
- CI / CD: Circle CI / CodePipeline
- External API Service: BitFlyer Lightning API, Google Chart API

More detail below

| Name           | Version  | Description                |
|----------------|----------|----------------------------|
| Nuxt.js        | 2.0.0    | UI Framework               |
| Vuetify        | 2.4.0    | UI Component               |
| Golang         | 1.16     | API Server                 |
| Testify        | 1.5.1    | Test for Golang            |
| Python         | 3.8.2    | API Server                 |
| PostgreSQL(RDS)| 13.3     | Database                   |
| Docker         | 20.10.10 | Container                  |
| Docker-Compose | 1.29.2   | Managing Containers        |
| Github         | 2.32.0   | Hosting Code               |
| Terraform      | 3.22.0   | Infrastructure as Code     |
| Circle CI      | 2.1      | CI CD pipeline             |
| Google Chart   | 0.3.3    | Candle Chart Component     |

#### DashBoard Image(Home)
<img width="1440" alt="Screen Shot 2022-01-17 at 21 57 45" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/c04abbfa-2f8c-41b8-a6aa-cf2fa3f04b32">

#### DashBoard Image(History)
<img width="1440" alt="Screen Shot 2022-01-30 at 16 37 20" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/02a71f6e-a535-4138-8dad-c486c6b9cfcb">

#### DashBoard Image(Memo)
<img width="1440" alt="Screen Shot 2022-01-30 at 16 43 15" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/6a158d96-83f1-44d4-bc02-2618fe8f9bec">

#### DashBoard Image(Alert)
<img width="1440" alt="Screen Shot 2022-01-30 at 16 44 42" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/7c2d02d4-135a-44aa-821a-f80dae586525">

#### DashBoard Image(Diagnosis)
<img width="1440" alt="Screen Shot 2022-01-30 at 16 48 08" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/b6a026e4-781d-4954-afa2-b1b11b863162">

#### Image Architecture for Infrastructure on AWS
<img width="897" alt="Screen Shot 2022-01-30 at 16 27 49" src="https://github.com/serpedious/automatic-trading-system/assets/40052091/c207a612-af64-4859-b7f7-004f53e4e69d">

#### Image Database Architecture on postgreSQL
![Automatic-Trading-System drawio (1)](https://github.com/serpedious/automatic-trading-system/assets/40052091/f5f1048c-0658-454e-a78e-01fd8add5781)

## HELP

### default guest user
- email: guest@gmail.com
- password: password

### Exeute App(when you execute this app in your local environment, you need to excute items below)
- create .env and set proper parameters(local is same with example env)
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
