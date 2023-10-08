
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
2. Run the command ```go go build```
3. This will create a binary file called 'Fetch_API' and place it in the root of the project
4. Run the command ```go run build``` and this will run the project. You should see 'Listening and serving http on :8000'
5. You will need to open a seperate terminal to use curl commands

Here is an example of a curl command that will request data and activate the /add endpoint 

```bash
curl localhost:8000/add --include --header "Content-Type: application/json" -d '{"payer": "DANNON", "points": 300, "timestamp": "2022-10-31T10:00:00Z"}' --request "POST"
```

You should run that in a seperate terminal while the server is listening and you should get a 200 response if you have valid data.

