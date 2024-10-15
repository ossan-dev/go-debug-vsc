# Remote Docker Debug

## Commands

- docker build . --tag greeting-ms --file Dockerfile.debug
- docker run -p 8080:8080 -p 2345:2345 greeting-ms
- Start the Debugger in VSCode
- Sets some breakpoint up
- run this: `curl http://localhost:8080/greetings?name=ossan`