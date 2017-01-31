# Golang-project architecture
Defining the good golang architecture

#Journey
This project is for trying to define the better architecture for golang project. I strongly believe this "Any fool can write code that a computer can understand. Good programmers write code that humans can understand" -Martin Fowler. statement and do the work in same passion. So I have added comments describing my code wherever possible.

#Getting Started
Fork and clone the repository.
run
```
go run main.go
```
Note: Import golang vendors if required.

#Architecture
```
app
   api
   controllers
   models
   microservices
   helpers
   validations
config
databases
main.go  
  
```
![Architecture](http://s3-sa-east-1.amazonaws.com/todovapersonal/golang-architecture.png)

##api
   api package contains all the REST API routes where validations will be done in validations package with same file name and same function names as in api package. This validations will be done through MIDDLEWARE concept.
It is simple to take the request,validate,pass params to controllers

##controllers
   api routes will call the controllers where file name and function names are same as api package.
   This is the main center where function logic will be done like calling to microservices ,for loops ,connecting third parties etc. 
   
##models
   models are simply table models and functions to query and fetch from database. This layer is simply connected to database and do all functions. It does not have any controller logics. It is simply for take query params from controllers and query database and return response.
   
##microservices
   microservices are for third party services like amazon,sockets,mailchimp, some background cron jobs etc.
   
##helpers   
   helpers are simply utility functions for environment variables, password encryptions,jwt token.
   
##validations
   All route validations will be done here as a middleware to routes. so the routes will be CLEAN and SIMPLE.
   
##config 
   config folder will contain keys.jso, push notification certificates file etc.
   
##databases   
   databases will have connection to databases. Can also write functionalitites like indexing etc when server starts.
   
##main.go
main.go is the main file where everything will be connected and starts running server.

#REST Routes
##Register User
POST->user/register
```
 {
  "email":"bathula@gmail.com",
  "password":"dlf"  
}
```

##Login User
POST->user/login
```
{
  "email":"bathula@gmail.com",
  "password":"dlf"
 
}
```



##Love :heart: to hear feedback from you
RT Bathula-weirdo,coffee lover
battu.network@gmail.com
