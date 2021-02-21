# Interval Merger
Simple REST Interface for merging intervals.

## Installation

```shell
# clone
git clone https://github.com/pomberer/go-interval-merger.git

# build
docker build -t interval-merger .

# run
docker run -p 8080:8080 interval-merger

# use
curl --location --request POST 'http://localhost:8080/api/v1/merge' \
--header 'Content-Type: application/json' \
--data-raw '{
    "intervals": [{
        "Low": 25,
        "High": 30
    },
    {
        "Low": 2,
        "High": 19
    },
    {
        "Low": 14,
        "High": 23
    },
        {
        "Low": 4,
        "High": 8
    }
    ]
}'
```

## Endpoints
* `POST /api/v1/merge`

sample-data:
```json
{
  "intervals": [{
    "Low": 25,
    "High": 30
  },
    {
      "Low": 2,
      "High": 19
    },
    {
      "Low": 14,
      "High": 23
    },
    {
      "Low": 4,
      "High": 8
    }
  ]
}
```

## Unit Test / Benchmark

```shell

# run unit tests
go test

# run benchmark (10000 intervals range 1 -100)
go test -bench=.
```

## Stats
```
* Memory Usage < 2MB
* Execution Time 10000 intervals < 1ms (MacBook Pro 2019 i9 2,3GHz)
```