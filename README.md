# CLI Orbital-Emergency-Control

cli programm to conrol orbital emergency station

## Features

- Authorization
- Check history
- Change the room
- Every room (check zone report, fix zone)
- Check zone list
- Get some resourses from storage
- Send emergy signal
- Check inventory

## How to run 

terminal:

``` bash
go run ./cmd/app
```

## Project Structure

``` text
/cmd/app - application entry point
/internal - all business logik
/auth-data.env - store for security data
```

# Scenario 

### Main scenario

User is admin, who is working in the Orbital Station. Your main purpose is fixing all zones and send SOS signal.

### Plan for fast win 

1. Authorise
2. Go to the storage
3. Take 3 stuff without medicine
4. Come back at control room and get access card
5. Go to the reactor and fix it
6. Go to the Life Support and fix it
7. Get medicine from storage
8. Go to the Communication and fix it
9. Sent SoS signal from control room

- Congratulations! You **Win !**

### Main rules 

almost every step spend some energy and oxygen. So be carefull to don't let them fall less than zero.

- Fixed Reactor gives you 100 energy point for the hull game. 
- Fixed Life Support gives you 100 oxygen points for the hull game.


If u wanna leave without win or lose, enter in the control room:

``` bash
0
```
It is developer exit


## auth-data.env
``` bash
LOGIN=admin
PASSWORD=12345
```