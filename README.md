# Logviewer 
This reads a log file for docker logs and displays it on a front end vuejs app. Think of it as a poorman's logging system for the lower cost app servers who don't want to spend money on logdna, logz, datadog etc..

It runs on a docker container itself for both the front end and backend portion.

## Screenshot


## Quick Start
- copy .env.example to .env ```cp .env.example .env```
- edit .env as needed
- The project requires the frontend directory to contain a dist folder with it built out. 
- Start by running ```cd frontend``` ```npm i``` ```npm run build```
- Now you can build out the project in root folder. To run the project you can use make command ```make local```
- if you do not have make installed just use docker-compose ```docker-compose up```

Missing docker-compose? [how to install docker-compose](https://docs.docker.com/compose/install/)


