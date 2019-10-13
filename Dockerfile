FROM python:3.7

RUN mkdir /app
WORKDIR /app
ADD cyntest /app/
ADD src/templates /app/templates
RUN chmod +x /app/cyntest

EXPOSE 5000
CMD ["/app/cyntest"]