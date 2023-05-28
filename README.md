### Golang + MySQL + Gin + Gorm 

Requirements: Go, [Docker](https://www.docker.com/)

Steps:

```
# Clone the project
git clone https://github.com/agilsyofian/apploan.git

# Run containers
docker compose build
docker compose up -d

# Confirm it all went up correctly. Exit with Ctrl+C
docker compose logs
```


Testing with [Postman](https://www.postman.com) 

Steps:

```
# Download and install Postman
# Import postman_collection.json
```




### Feature

1 Registration

This feature is used to register a username and password as well as some other data and afterwards you will get a credit limit according to the data entered.


2 Log in

This feature is used to request access rights in the form of a token which will later be used to access other features. This token is used in the Authorization header and will be validated for each access. This token has an expiration time of 15 minutes.


3 Renew Tokens

This feature is used to update expired tokens.


3 Create Contracts

This feature is used to create a new contract or loan. During contact creation, a user's credit limit will be checked. If the limit has been met then it will fail and vice versa. This contract has a tenor according to the input by the user.


4 Pay Loans

This feature is used to pay bills that have been made before. Users can make single or multiple tenor payments. After the bill is fully paid, it will affect the credit limit that was previously cut.




### Architecture Application
![Alt text](https://github.com/agilsyofian/apploan/blob/master/Flow.png?raw=true "Optional Title")


### ERD (Entity Relationship Diagram) 
![Alt text](https://github.com/agilsyofian/apploan/blob/master/ERD.png?raw=true "Optional Title")
