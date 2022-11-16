# Crawler in Go to get Data from Coords, segments, stores, category

 
 You can pass parameters from command line to define what you want to get from James.

# Steps Install
    1. https://golang.org/doc/install ( Install Go By your OS )
    2. open directory "james" and Write:
    ` go install `
    3. Up PostgreSql
    ` docker-compose up -d `
    
# Commands

## Get All ( Grande São Paulo )

` go run main.go `

## Get from Other Coords

` go run main.go -lat "-23.584293" -lng "-46.674584" `

## Get from Other Coords and a specific segment

` go run main.go -lat "-23.584293" -lng "-46.674584" -segment "mercados" `

## Get from a specific store

` go run main.go -store "Carrefour" `

## Get from a specific Category

` go run main.go -category "Bebidas" `

## Set a name or number for this scan

` go run main.go -scan "1" `

## get multiples coordenates from file

` go run main.go -file="leitura.csv" -segment="farmacia" `

## Example args

` go run main.go -lat "-23.584293" -lng "-46.674584" -segment "mercados" -store "Carrefour" -category "Bebidas" -scan "1" ` 

## Clean Database 

` make clean ` 

## Upload S3

`bucket : eiprice.delivery /james`




