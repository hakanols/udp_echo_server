UDP Echo Go server
==================

Needed a online UDP echo server for some experimentation with Nordic nRF9160 and LTE-m and NB-IoT.
The echo server is of course generic. This repo also contain a test client, deployment instructions and
some experimentation with Distroless.

Docker Distroless
-----------------
I am a strong believe in keeping complexity down. I believe that
[Distroless](https://github.com/GoogleContainerTools/distroless) and
[Multistage build](https://docs.docker.com/develop/develop-images/multistage-build/) succeed in that.
* Platform independent build and run environment (Linux, Windows, Mac & cloud)
* Immutable environment
* No unnecessary "stuff" in run environment
* Relative uncomplicated configuretion (8 line Dockerfile for both build & run)
* Only Docker as dependency

One valid critic is that perhaps one type of complexity is just changed for another.
Some times that is the case. E.g. How to debug?
Mostly solvable but with added complexity. There are no silver bullets but for my case
I would argue that it is a perfect match.

See (distroless-go-hello) for an minimal example.

Dependencies for examples below
--------------------------------
* [GO Lang](https://golang.org/)
* [Docker](https://www.docker.com/)

Local echo server without Docker
--------------------------------
Terminal 1:

	cd echo_server
	go run main.go
	
Terminal 2:

	cd test_client
	go run main.go 0.0.0.0

Local echo server with Docker
--------------------------------
Terminal 1:

	cd echo_server
	docker build --rm -t echo .
	docker run --rm -p 1234:1234/udp -t echo
	
Terminal 2:

	cd test_client
	go run main.go
	
Echo server on Google Could
---------------------------
[Start server on Google cloud](GoogleCloud.md)


Terminal:

	cd test_client
	go run main.go [SERVER_IP]
