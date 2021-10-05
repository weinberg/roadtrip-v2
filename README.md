# RoadTrip V2

Take a RoadTrip cross country with your friends... in your terminal!

RoadTrip is a low intensity casual background game requiring very little interaction and can be played at your own pace while working or doing other activities.

## Current State

What comes before Alpha? Currently you can run the server and client and create a car which will travel from Seattle to Cheyenne. No fancy features exist yet.

## Design

RoadTrip is a client server application. The client is written in Go for cross-platform support. The API server is written in node/typescript with a postgres database and prisma.
The server parts are built to run in containers using either Kubernetes or Docker Compose.
