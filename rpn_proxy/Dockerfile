FROM python:3.7-alpine
RUN apk add ruby
MAINTAINER Piotr Kaliński "pi.kalinski@gmail.com"
COPY . /app
WORKDIR /app
EXPOSE 5000
RUN pip install -r requirements.txt
CMD python ./app.py