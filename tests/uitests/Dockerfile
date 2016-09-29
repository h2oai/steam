FROM pl31/chrome

RUN apt-get update && apt-get install -y \
	python python-pip curl unzip openjdk-8-jdk procps vim

RUN pip install h2o selenium

WORKDIR /root
RUN curl -O http://chromedriver.storage.googleapis.com/2.24/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN mv chromedriver /usr/bin/chromedriver
WORKDIR /root/steam

ENTRYPOINT ["./runtest.sh"]

