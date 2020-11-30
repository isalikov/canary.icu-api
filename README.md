[![Build Status](https://cloud.drone.io/api/badges/isalikov/canary.icu-api/status.svg)](https://cloud.drone.io/isalikov/canary.icu-api)

# canary.icu-api

Canarian news api

### Requirements
- redis

```bash
# docker run example

docker run --name redis \
    -v /tmp/data:/data \
    --network canary.icu-net \
    --rm  -d redis
```

### Running
```bash
docker run --rm --name canary.icu-api \
    --network canary.icu-net \
    -e REDIS_HOST=redis \
    -e REDIS_PORT=6379 \
    iknpx/canary.icu-api
```
