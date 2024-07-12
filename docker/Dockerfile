FROM ubuntu:20.04
#
# Setup redishoneypot
#
COPY ftp /opt/ftp/
COPY ftpserver.json /opt/ftp/

# Start redishoneypot
WORKDIR /opt/ftp

CMD ./ftp

# CMD tail -f
# docker build -t hary654321/ftp  .

