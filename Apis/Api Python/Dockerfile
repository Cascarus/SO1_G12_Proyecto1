FROM python:latest
WORKDIR /app/src
COPY . .
RUN pip install --upgrade google-cloud-pubsub
RUN pip install flask
RUN pip install pymongo
RUN pip install mysql-connector-python
RUN pip install waitress
RUN pip install python-dotenv
EXPOSE 5000
CMD ["python", "api.py"]