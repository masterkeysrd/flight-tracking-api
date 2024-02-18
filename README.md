# backend-take-home

For backend candidates applying to moneda this is your take home assignment

## Objective

Develop a backend service for a flight tracking application that provides real-time flight information using ADS-B data. This service will support frontend features such as searching for flights by ICAO or IATA numbers and visualizing nearby air traffic. The backend should efficiently process and serve flight data to the frontend application.

## Project Requirements

Environment Setup:

- Set up a go environment: <https://go.dev/doc/install>. We recommend using brew for easier install
- Fork the repo/use the template from: [Take home repo](https://github.com/Moneda-Tech-Group/backend-take-home)

API Development:

- Implement a basic server using the provided skeleton code in main.go. You server should implement some kind of basic authentication to protect your endpoints
- flightdata.json is a json file that will act as your database. You should use this to help power your two endpoints
- Implement getFlightData endpoint
  - This should be a post request that can pass any of these parameters. At least one parameter should be passed or you should return an error if you have an empty parameter request

  ```text
  bbox              optional Bounding box (South-West Lat, South-West Long, North-East Lat, North-East Long)
  zoom              optional Map zoom level to reduce the number of flights to speed up rendering (0-11)
  hex               optional Filtering by ICAO24 Hex address. 
  reg_number        optional Filtering by aircraft Registration number.
  airline_icao      optional Filtering by Airline ICAO code.
  airline_iata      optional Filtering by Airline IATA code.
  flag              optional Filtering by Airline Country ISO 2 code from Countries DB.
  flight_icao       optional Filtering by Flight ICAO code-number.
  flight_iata       optional Filtering by Flight IATA code-number.
  flight_number     optional Filtering by Flight number only.
  ```

  - It should return a response similar to below which return all flight data based on the passed in parameters

  ```json
  [{
      "hex": "780695",
      "reg_number": "B-5545",
      "flag": "CN",
      "lat": 28.397377,
      "lng": 115.1008,
      "alt": 7078,
      "dir": 269,
      "speed": 775,
      "v_speed": -7.8,
      "squawk": "0205",
      "flight_number": "9429",
      "flight_icao": "CSH9429",
      "flight_iata": "FM9429",
      "airline_icao": "CSH",
      "airline_iata": "FM",
      "aircraft_icao": "B738",
      "updated": 1626153069,
      "status": "en-route"
      }, {
      ...
  }]
  ```

- Implement searchFlightInfo endpoint
  - This should be a post request that can pass any of these parameters

  ```text
  flight_icao     required Search by Flight ICAO code-number.
  flight_iata     required Or search by Flight IATA code-number.
  ```

  - It should return a response similar to below which should return a single json object of the flight is found

  ```json
  {
  "hex": "AAB812",
  "reg_number": "N790AN",
  "aircraft_icao": "B772",
  "flag": "US",
  "lat": 33.455017,
  "lng": -118.738312,
  "alt": 10668,
  "dir": 80,
  "speed": 942,
  "v_speed": 0,
  "squawk": "3726",
  "airline_icao": "AAL",
  "airline_iata": "AA",
  "flight_number": "6",
  "flight_icao": "AAL6",
  "flight_iata": "AA6",
  "duration": 434,
  "updated": 1626858778,
  "status": "en-route",
  }
  ```

As a stretch goal, provide a way to paginate your data.

### Security

Implement authentication for the API to secure access.

### Documentation and Code Submission

Provide a README.md file with any setup instructions, API documentation.
Include any assumptions made or important decisions in your design process.
Share your code via GitHub repository and ensure it's well-commented to explain key functionalities and design choices.

## How we will evaluate your submission

Architecture & Design: Clarity and scalability of the backend architecture.
Code Quality: Organization, readability, and use of Go best practices.
Problem-Solving Skills: Creativity and efficiency in solving challenges related to data integration and API performance.
Documentation: Completeness and clarity of the project documentation.

### Share github to these github users: exs6350, hterrero01

### Duration

Allocate no more than 4 hours to this assignment to focus on the core functionalities. The aim is to assess your approach to designing and implementing a scalable backend system rather than completing all features in detail.
