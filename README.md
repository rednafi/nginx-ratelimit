# nginx-ratelimit


A simple hello-world JSON API service incorporates rate limiting via Nginx. It allows 
50 requests per second and returns a 429 error when the limit is exceeded.

See the blog post [here].

## Installation

```sh
git clone git@github.com:rednafi/nginx-ratelimit
```

## Take it for a spin

-   Start the service:

    ```sh
    docker compose up -d
    ```

-   Test the rate limiting by sending many requests:

    ```sh
    seq 200 | xargs -n 1 -P 100 bash -c 'curl -s localhost/greetings|jq'
    ```

    This returns:

    ```json
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


[here]: https://rednafi.com/go/rate_limiting_via_nginx
