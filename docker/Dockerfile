FROM ubuntu:20.04
#
# Setup redishoneypot
#
COPY mysqld /opt/mysql/
COPY dist/* /opt/mysql/

# Start redishoneypot
WORKDIR /opt/mysql

CMD ./mysqld

# CMD tail -f
# docker build -t hary654321/mysql  .

