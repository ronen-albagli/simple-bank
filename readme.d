 ##     Go Simple Bank

 #### Intro
 ###### In this porject we will implemenets a simple Bank account manager
 &nbsp;
 ###### Technologies will be used in this project:
  * Golang
  * MongoDB
  * RabbitMQ

###### This project will be based on Gin (gin-gonic)
&nbsp;
##### Architecture
###### We will take the Clean-Architecture approche (Robert C. Martin (“Uncle Bob”)) - hexagonal architecture

###### Each Part of the application is stand alone, No dependencies between application parts

&nbsp;
#### Project Structure
* API - Responsible for all http communications
* Repository - Responsible for the Data access points - integrate with mongo,sql,etc...
* Domain - Responsible for the business logic
* Config - Keeps the app config





