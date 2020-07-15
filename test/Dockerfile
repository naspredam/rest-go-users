FROM golang:buster

ADD . /app/
ADD test/run_all_tests.sh /app/

WORKDIR /app

ENTRYPOINT ["sh", "run_all_tests.sh"]