# Coding Exercise

The objective of this take-home exercise is to develop a simple in-memory key-value store, exposed through a REST-ful HTTP interface.

The server may be implemented using any language/HTTP framework, as long as the project README explains how to run it. When you are done, push your code to this GitHub repository.

Don't worry too much about extensive documentation or tests, just focus on writing clean code.

## Requirements

1. HTTP endpoints:
    * `PUT /<key>` — Set the value of a key (value in request body)
    * `GET /<key>` — Retrieve the value of a key
    * `DELETE /<key>` — Delete a key
    * `GET /<key>/history` - Return a JSON-formatted representation of a key's history.
2. The server should store its keys and values in memory and must not rely on any other process (e.g. a database).
3. Keys will be between 1 and 255 bytes long and will only use the ASCII characters `a-z`, `A-Z`, `0-9`.
4. Values will be arbitrary byte strings with a maximum length of 1024 bytes.
5. Timestamps should be expressed as Unix Epoch timestamps with milliseconds.
6. The history endpoint should return all writes to the given key for the past two minutes, in descending timestamp order. Delete events to a key should be indicated with a `null` value; deletes should not clear the key's earlier history.

Try to keep the implementation **as simple as possible** while keeping the server robust.
We are much more concerned with reliability than performance.

## Example requests

These examples assume the HTTP server is running on `localhost`, port `3000`.

1. Write a value to the key `foo`:

    ```sh
    $ curl -i -X PUT 'http://localhost:3000/foo' -H 'Content-Type: application/octet-stream' --data-binary 'hello world!'
    HTTP/1.1 204 No Content
    ```

2. Fetch the value of `foo`:

    ```sh
    $ curl -i 'http://localhost:3000/foo'
    HTTP/1.1 200 OK
    Content-Type: application/octet-stream
    Content-Length: 12

    hello world!
    ```

3. Fetch the history of `foo`:

    ```sh
    $ curl -i 'http://localhost:3000/foo/history'
    HTTP/1.1 200 OK
    Content-Type: application/json
    Content-Length: 52

    [{"value":"hello world!","timestamp":1605097943904}]
    ```

4. Delete the key `foo`:

    ```sh
    $ curl -i -X DELETE 'http://localhost:3000/foo'
    HTTP/1.1 204 No Content
    ```

5. Fetch the history of `foo` again, after the delete:

    ```sh
    $ curl -i 'http://localhost:3000/foo/history'
    HTTP/1.1 200 OK
    Content-Type: application/json
    Content-Length: 93

    [{"value":null,"timestamp":1605098003047},{"value":"hello world!","timestamp":1605097943904}]
    ```

6. Fetch the value of `qux` (which does not exist):

    ```sh
    $ curl -i 'http://localhost:3000/qux'
    HTTP/1.1 404 Not Found
    ```
