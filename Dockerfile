FROM devkitpro/devkitppc

# Install thirdparty ppc libs and other requirements
RUN apt-get update && \
	apt-get install -y --no-install-recommends git patch golang wget unzip cmake make && \
	git clone https://github.com/hoverguys/ppc-portlibs.git && \
	./ppc-portlibs/build.sh entityx && \
	rm -rf ./ppc-portlibs && \
	git clone https://github.com/hoverguys/EASTL.git && \
	apt-get remove -y unzip wget patch && \
	apt-get autoremove -y && \
	apt-get clean

ENV GOPATH=/go

VOLUME /rehover
WORKDIR /rehover/build

# Entrypoint
COPY ./docker-entrypoint.sh /
ENTRYPOINT "/docker-entrypoint.sh" && /bin/bash
