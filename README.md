
<div align=center>

# Fetch API Intern Project 

</div>


Hello my name is Clay Truelove from the University of Alabama and I made this application as a part
of the take home assignment from Fetch Rewards

## Endpoints

The idea of this project is run a local server and create 3 seperate endpoints that use POST and GET requests.
We then decode this data and store it in structs, slices, and maps to manipulate it.
I have a seperate folder for every function and a "type" folder for global data strucutes.

### Things you need to run this

You must have Go installed
You can download it here https://golang.org/dl/

### Basic Steps to run this project

1. Clone the repository and go to the root Directory called 'Fetch_API' 
2. Run the command ```bash go build```
3. This will create a binary file called 'Fetch_API' and place it in the root of the project
4. Run the command ```bash go run Fetch_API``` and this will run the project. You should see 'Listening and serving http on :8000'
5. You will need to open a seperate terminal to use curl commands

### Using the Add endpoint


Here is an example of a curl command that will activate the /add endpoint 

```bash
curl localhost:8000/add --include --header "Content-Type: application/json" -d '{"payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z"}' --request "POST"
```

### Using the Balance endpoint


Here is an example of a curl command that will request data and activate the /balance endpoint 

```bash
curl localhost:8000/balance
```
This makes a GET request and returns the Payer and their associated balance from our local data structure

### Using the Spend endpoint


Here is an example of a curl command that will activate the /spend endpoint 

```bash
curl localhost:8000/spend --include --header "Content-Type: application/json" -d '{"points": 300}' --request "POST"
```

I find it easier to store the JSON data in a seperate .json file and include it in the body of the request.
Here is an example of a .json file and a command you can run that uses the .json file in the body

**`request.json`**
```json
{
"payer" : "DANNON",
"points" : 5000,
"timestamp" : "2020-11-02T14:00:00Z"
}
```
Then we include this in our bash command to make it easier

```bash
curl localhost:8000/add --include --header "Content-Type: application/json" -d @request.json --request "POST"
```

### Thank You for taking the time to look at my project
I have been wanting to learn Go for back end developent so this project was really enjoyable for me. Regardless of weather you choose to go forward with my interview process I learned a lot from this project and really enjoyed this project. 

Thanks, 
Clay