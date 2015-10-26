Smoker
-----

Very simple tool to test a list of URLs for expected response codes and output the time the request took.

# Usage

Use **-file** to specify the file containing URLs to test with expected return code (Default: urls.json)

Use **-cookies** to set specific cookies


# Examples

Test urls listed in urls.json:

    $ ./smoker
    http://httpbin.org/ 200 181.786687ms
    https://httpbin.org/ 200 429.430469ms
    https://httpbin.org/status/404 404 343.816642ms
    https://httpbin.org/redirect-to?url=http%3A%2F%2Fexample.com%2F 302 85.686784ms
    https://httpbin.org/basic-auth/user/passwd 401 83.938694ms

Test urls listed in staging.json and send a couple of cookies:

    $ ./smoker \
      -file=staging.json \
      -cookies="foo=foo; bar=bar"

If all URLs respond as expected, then it will exit with status 0, otherwise it will exit 1


# Install

There is a simple Makefile which will build and install smoker for you to **/usr/local/bin**:

    $ make && sudo make install

