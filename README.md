# GO Rest API Template
---
### This project starter schema can help develop Go project especially while building a REST API
Project Starter Flow Visualization
![alt text](https://i.imgur.com/uHjiu7S.png)
---
> Directory root explaination
> This is explaination about root diretory
## `/internal`
Internal folder contain 5 (sub) folders:
- `/internal/api`:
this folder contain `routes.go` file to regist all of your api using **gin** 
- `/internal/db` :
in here you can initialize you databases connection on `db.go`, the connection that you can used can be variative depends what you used
- `/internal/handlers`: 
`"models"_handler.go` can be utilize to handlers the endpoint services, this module makes the project more structure and easy to maintain
- `/internal/services`:
service module contain a lot of service for the user request and responses, in this you can implement query to interact to your own db and in this module would help validation data with `/internal/utils` module
- `/internal/utils`:
this module will contain a lot of function that will help handler module or service module with this module can make the maintenance more easily