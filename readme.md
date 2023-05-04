build:
 docker build -t sports .

run:
 docker run -p 3000:3000 sports
 o
 docker run -e PORT=8000 -p 8000:8000 sports

 [![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)