# Reverse Polish Notation application
 
 
## Info
 
This project is separated in three modules each responsible for different aspect of project

 
## Solution details
 
When I started this project, I got to figure out how to split functionalities and choose best technology for each module.

* Endpoint interface for gathering input data (stdin) and output results (stdout) - Golang
* API for getting requests, sending to the worker and logging input and timings. Logs should be saved to a file. - Python
* Worker responsible for the process of calculation - Ruby
 
## Prerequisites
 
Before you start a tour with this project, be sure that you have installed:
 
   * Docker
 
Verify that you have successfully installed docker by typing in terminal:
 
`docker run hello-world`
   
Solution is mostly oriented for Linux-related users.  
 
## How to start project?
 
1. Go with your terminal to the directory of the project
 
   
    /cd/to/your/path


2. This project uses three containers

* nginx - reverse proxy to allow communication between containers
* rpn_proxy - python-flask based api with ruby script to run Reverse Polish Notation algorithm script
* golang - interface and simple web interface
 
  
3. When you are in the main directory of project, write:
 
   
    docker-compose up


## Functionalities

#

Api contains one method:
 
### Expression to process:
 
URL for method
 
    http://localhost:5000/rpn/
 
Body
   
   * expression int[] - Array of expressions to process

 
## Issues, important notes

 
In terms of showing processed expression, I had to think about how to split technologies between three parts. Flask server suited well for small project as this one, Ruby for calculating expression and Golang as procedural language suited well for interfaces with typed structs.
 
If I had a possibility to change, or improve something I'd rather:
 
   * Rewrite python logging.
   * Perhaps split calculation class to separate logging method.
   * There is surely a place for tool that creates some cool documentation like Swagger - if I had a bigger experience with it I'd definitely used it.
   * If I had any earlier experience with Python I would definitely try to make tests.
 