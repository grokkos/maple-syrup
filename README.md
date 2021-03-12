<br />
<p align="center">
  <a href="https://github.com/gerokkos/clerk">
    <img src="https://i.ibb.co/pZvbmrs/901de892-68c5-44a2-9810-2dc6b9498931-200x200.png" alt="Logo" width="80" height="80">
  </a>
  <h3 align="center">Maple Syrup API</h3>
</p>

# Table of Contents

* [Getting Started](#getting-started)
* [Project Structure](#project-structure)
* [Integration Guide ](#integration-guid)
* [Use the API](#use-the-api)
* [Clone the Application](#clone-the-application)
* [Run with Docker](#run-with-docker)
* [Run the Application Locally](#run-the-application-locally)
* [Clean tests cache if cached](#clean-tests-cache-if-cached)
* [Contributing](#contributing)





# Project Structure


    ├── Maple Syrup                    
    │   ├── api             # Functionality/features
    │   ├── tests           # Testing endpoints  
    │   └── main            # Run the application

    The database used is PostgreSQL



## Integration Guide 

* A Batch can be either dispatched or un-dispatched
* A Batch is dispatched if it resulted in a Transaction being sent to the bank.
* A Batch is un-dispatched if it didn't send a Transaction to the bank
* There is exactly one un-dispatched Batch at any point of time
* A Batch contains an accrued amount, with an initial value of 0
* A User can send low volume amounts to a Batch and the accrued amount of the Batch increases
* Once the accrued amount of the Batch goes above the threshold of 100, it gets dispatched. The created Transaction has a an amount equal to the accrued  amount of the Batch
* An history of all Batches, dispatched or un-dispatched must be kept

``
Step 1
``
When the application is firing up, it creates the first undispatched Batch and the User which is attached to


``
Step 2
``
Create a new User with only Name and keep the generated id


``
Step 3
``
Create a new Roundup with Amount and User id, and will automatically check the Undispatched Batch Summary


``
Step 4
``
If the Summary per Batch exceed 100, the Batch is being Dispatched and a Transaction with the Summary is created

``
Step 5
``
A new Undispatched Batch with Summary 0 will be created which the next Roundups will be attached to, until the limit exceeds again

### Use the API

| Endpoint         |                              |   
| -------------    | -----------------------------|
| /user            | POST                         |
| /roundup         | POST                         |
| /users           | GET                          |
| /batchlist       | GET ?id="userid"             |
| /roundups        | GET                          |
| /batches         | GET                          |
| /transactions    | GET                          |





### Clone the Application

``
git clone github.com/grokkos/maple-syrup.git
``


### Run with Docker

Start up the application services by running:

``
docker-compose up
``

Call the API from http://localhost:8080


To stop the services run: 

``
docker-compose down
``


Run the tests with Docker:

``
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
``




### Run the Application Locally


``
go run main.go
``


### Clean tests cache if cached


``
go clean -testcache
``


## Contributing

1. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
