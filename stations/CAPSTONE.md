#Capstone Project

## Summary

The purpose of this capstone project is to create a CLI tool to provide information on London's tube system (station name, zone etc.) and determine the closest station from the users location using longitude and latitude. I'll be retrieving data obtained from https://www.doogal.co.uk/london_stations, station data is in JSON format. The prohect will involving crating an API server for the CLI to interact with. 

I will demonstrate the following: 

1. Using HTTP GET methods to pull from the json file
2. Unmarshalling and store the data from the json file into a struct
3. Create a server 
4. Create a client that can interact with the server
5. Handle errors and use best practices
6. Testing of HTTP handlers 

## User Stories

### As a user, I would like to invoke the CLI app with the list of stations closest; starting with the closests, based on the longitude and latitude of the center of the city of London as default. As well, as giving the user the option to edit the longitutde and latitude or select a station on the list. 

**Acceptance Criteria**

When the program is running, the list of three stations will appear starting with the closest. At the bottom there will be a prompt to select a station or enter new coordinates. 

Example:
```
$ go run main.go
```
```
Closest Stations from the center of London:
1. ...
2. ...
3. ...

Press # to enter new coordinates or select a station.
```

### As a user, I would like to input in coordinates to find out what the closest stations are to my current location. 

**Acceptance Criteria**

When the program is running, the user will enter # and then will be prompted to enter new cooridnates. 

Example:

```

Enter logitutde & latitude:
....... .......
```

```
Closest Stations from the center of London:
1. ...
2. ...
3. ...

Press # to enter new coordinates or select a station.
```

### As a user, I would like to select a station and receive station information.  

**Acceptance Criteria**

When the program is running, once the user enters their coordinates they will be able to select a station. Once they have the station's name, distance, zone, marker-color will appear. A different message will appear if they selected the 2nd or 3rd closests. There will be an options to return to the list or to start over. 

Example:

```
Closest Stations from the center of London:
1. ...
2. ...
3. ...

Press # to enter new coordinates or select a station.
```

```
Closest Stations to your coordinates:
Name: Ampere Way
Distance: 5 miles
Zone: 3,4,5,6
Mark-Color: yellow

Press "b" to return to the list or press "r" to start a new search. 
```

```
Second Closest Stations to your coordinates:
Name: Anerley
Distance: 5.9 miles
Zone: 4
Mark-Color: GreenYellow

Press "b" to return to the list or press "r" to start a new search. 
```