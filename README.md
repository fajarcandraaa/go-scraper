# Scraping data from some marketplace sites using golang and gocolly libraries

<!-- ABOUT THE PROJECT -->
## About The Project

This is a little project about how to scraping data from marketplace using golang and gocolly library. 

## About code structure
I'll sort by flow :
1. Service -> is about business logic
2. Repository -> is a layer between the Service and data layers of your application with an interface to perform create, read, update and delete CRUD operations. By using repositories in your application

Than for this directory, not always go with the flow
3. DTO -> is a layer that used to transform data to an object
4. Presentation -> is used to present some paylod or response
5. Entity -> a layer that representation from database
6. Helpers -> is helper, used to handel reusable function
7. Config -> is layer that handle configuration of our database

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)
* [PostgreSQL](https://www.postgresql.org)
* [Colly](http://go-colly.org)

<!-- GETTING STARTED -->
## Getting Started
Before we get started, it's important to know that  before you run this code you have to make sure that `Redis` is already exist and ready to run on your device. Than this code use a custom command to execute it with makefile to make more simple command like :
1. make update
2. make tidy
3. make start

So, let start it.
1. After clone this repository, just run `make update`.
2. Setup your `.env` file such as database connection and redis connection based on default setting on you device.
3. To make sure that all dependency is run well, than run `make tidy`.
4. Finally, you can run your project with command: `make start`.
5. Go to postman and set url like `http://localhost:8080/`, for information that port to run this project depend on configuratin on `.env`

## Afterword
Hopefully, it can be easily understood and useful. Thank you~