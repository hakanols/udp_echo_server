UDP Echo Go server
==================

Needed a online UDP echo server for some experimentaion with Nordic nRF9160 and LTE-m and NB-IoT.
The echo server is ofcourse generic. This repo also contain a test client, deployment instructions and
some experimentation with Distroless.

[Docker Distroless](https://github.com/GoogleContainerTools/distroless)
-------------------
I am a strong beliver in keeping complexity down. I belve that distroless in and
[Multisage build](https://docs.docker.com/develop/develop-images/multistage-build/) succeed in that.
* Platfor independednt build and run envirement (Linux, Windos, Mac & cloud)
* Immutable environment
* No unesesasy "stuff" in run environment
* Relative uncomplicaded config (8 line Dockerfile for both build & run)
* Only Docker as dependanscy

One valid critic is that perhaps one type of komplexity is just changed for another.
Some times that is the case. E.g. How to debug?
Mostly solvabal but with added complexity. There are no silver bullets but for my case
I would argue that it is a perfect match.

See [distroless-go-hello] for an minimal example.

Dependencies for examples below
--------------------------------
* [GO Lang](https://golang.org/)
* [Docker](https://www.docker.com/)


Local echo server without Docker
--------------------------------
Terminal 1:

    cd echo_server
	go run main
	
Terminal 2:

    cd test_client
	go run main 0.0.0.0

Local echo server with Docker
--------------------------------
Terminal 1:

    cd echo_server
    docker build --rm -t echo .
    docker run --rm -p 1234:1234/udp -t echo
	
Terminal 2:

    cd test_client
	go run main
	
Echo server on Google Could
---------------------------
[Start server on Goolge cloud](GoogleCloud.md)


Terminal:

    cd test_client
	go run main [SERVER_IP]
