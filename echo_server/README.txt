docker build --rm -t echo .
docker run --rm -p 1234:1234/udp -t echo