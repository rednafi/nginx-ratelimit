# nginx-ratelimit

Rate limiting through reverse proxy. Allow 50 req/sec and return 429 error when the limit
is exceeded.

## Installation

```sh
$ git clone git@github.com:rednafi/nginx-ratelimit
```

## Take it for a spin

-   Start the service:

    ```sh
    docker compose up -d
    ```

-   Test the rate limiting by sending many requests:

    ```sh
    $ seq 200 | xargs -n 1 -P 100 bash -c 'curl -s 34.138.11.32/greetings|jq'
    ```

This returns:

    ```txt
    {
        "message": "Hello World"
    }
    {
        "message": "Hello World"
    }

    ...

    {
        "status": 429,
        "message": "Too Many Requests"
    }
    {
        "status": 429,
        "message": "Too Many Requests"
    }
    {
        "message": "Hello World"
    }
    {
        "status": 429,
        "message": "Too Many Requests"
    }
    ```
