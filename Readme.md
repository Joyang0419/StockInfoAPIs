# StockInfoAPIs

## Project Purpose
- Provide data of Taiwan stock

## Project Installation
- golang version: go 1.19
- install docker daemon first
```bash
# check work dir: ./StockInfoCrawler
cd deployments/
docker-compose up --force-recreate
```

## Project main used libraries
- Go-gin, API
- Go-swagger, provide convenient way to test API.
- Go-gorm, operate DataBase.

## project architecture
- 4 Tier architecture

![TierArchitecture](./docs/tier.png "TierArchitecture")


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
