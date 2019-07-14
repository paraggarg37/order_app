## Setup
Before you can clone and run this project make sure follwoing things are installed 
- GO
- MySQL 
- Docker 

## Running on local
### Update .env file 
1. Connection string in env file should be updated as `${user}:${password}@tcp(localhost:3306)/logistics`
2. Replace `${MAPS_KEY}` with your Key (Distance api should be enabled)
3. Update User and Password
4. Run :  `./start.sh` 

## Running TestCases 
1. Update .env file as stated above
2. Run : `./run_test.sh` 

## Running in Docker
1. Update `${MAPS_API_KEY}` in `docker-compose.yml` with the actual key
2. cd cmd 
3. Run : `make docker-dev-start` 
