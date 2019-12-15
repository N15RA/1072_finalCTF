FROM ubuntu:16.04
LABEL maintainer="roy@nisra.net"

# Apt update && upgrade
RUN dpkg --add-architecture i386
RUN apt-get update --fix-missing && apt-get -y upgrade
RUN apt-get install -y libc6:i386 libstdc++6:i386
RUN apt-get install -y libc6-dev-i386
RUN apt-get install -y ucspi-tcp
RUN apt-get install -y gcc make

# create ctf-user
RUN useradd -u 1337 -U -m -s /bin/bash ctf
WORKDIR /home/ctf/

COPY ./entry.sh /home/ctf/
COPY ./makefile ./main.c /home/ctf/
COPY ./flag /home/ctf/

# Perm
RUN chmod 664 /home/ctf/flag
RUN chmod 775 /home/ctf/*.sh

# Compile and cleanup
RUN make
RUN rm -f ./main.c ./makefile

USER ctf

# Set problem env var
ENV PROB=ret2sc
# entry
CMD ["/home/ctf/entry.sh"]