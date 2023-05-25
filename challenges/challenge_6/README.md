<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-3.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #5. Kafka
</h1>

In this 6th challenge we are going to integrate Kafka in our application

## Instructions
* Add a new field into Ad named type (value examples: automobile/laptop/house)
* For each inserted ad send a kafka message containing the ad_type 
* The consumer will receive the message and increment in a table ad_statistics(ad_type, number_of_ads) the number 
of that ad type


## Resources
1. Docker file example https://github.com/mica-tw/leboncoin-enablement/blob/main/backend/kafka-exercise/docker-compose.yml
2. Consumer/Producer example  https://github.com/mica-tw/leboncoin-enablement/tree/main/backend/kafka-exercise

