# Movie Festifal App

# How To run This Project
```bash
# clone repository
git clone git@github.com:IbnAnjung/movie_fest.git

cd movie_fest

cp .env.example .env

#if You'r host already install mysql and redis, just run database/sql_dump.sql on you'r sql client.

#or you can use docker for mysql and redis,
# run this command if you not use your mysql and redis on your host. 
docker-compose up -d 

#adjust you'r .env file

#now you can run this project
go run main.go

```