# Movie Festifal App

# How To run This Project
```bash
# clone repository
git clone git@github.com:IbnAnjung/movie_fest.git

cd movie_fest

cp .env.example .env

mkdir public
mkdir public/files

#if You'r host already install mysql and redis, just run database/sql_dump.sql on you'r sql client.

#or you can use docker for mysql and redis,
# run this command if you not use your mysql and redis on your host. 
docker-compose up -d 

#adjust you'r .env file

#now you can run this project
go run main.go

```

# Endpint List                         
- GET    /videos/*filepath         
- HEAD   /videos/*filepath         
- POST   /admin/movie/upload       
- PUT    /admin/movie/meta         
- GET    /admin/movie/votes/        
- GET    /admin/movie/votes/most    
- GET    /admin/movie/views/        
- GET    /admin/movie/views/most    
- GET    /admin/genres/views/most   
- GET    /admin/genres/votes/most   
- GET    /movie/                    
- POST   /movie/vote                
- POST   /movie/unvote              
- GET    /movie/voted               
- POST   /movie/start               
- POST   /movie/playback            
- GET    /movie/history             
- GET    /movie/:id                 
- POST   /auth/register             
- POST   /auth/login                
- POST   /auth/logout