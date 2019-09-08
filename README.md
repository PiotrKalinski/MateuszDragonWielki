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

* Golang
* Python, Flask
* Ruby

Solution is mostly oriented for Linux-related users.  

## How to start project

1. Go with your terminal to the directory of the project

    /cd/to/your/path

2. Go to /golang directory and type:

    go run main.go

3. Go to /rpn_proxy directory and type:

    python ./app.py

4. Type in browser

    http://localhost:4000/

## Functionalities

Api contains one method:

## Expression to process

URL for method

    http://localhost:5000/rpn/

Body

* expression int[] - Array of expressions to process

## Issues, important notes

In terms of showing processed expression, I had to think about how to split technologies between three parts. Flask server suited well for small project as this one, Ruby for calculating expression and Golang as procedural language suited well for interfaces with typed struts.

If I had a possibility to change, or improve something I'd rather:

* Rewrite python logging.
* Perhaps split calculation class to separate logging method.
* There was a try to close implementation in docker, but I had a problem with network. Golang app couldn't connect to flask api server.
* If I had any earlier experience with Python I would definitely try to make tests.
