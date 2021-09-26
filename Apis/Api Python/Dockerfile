FROM python:3
WORKDIR /app/src
COPY . /app/src
RUN pip install --upgrade google-cloud-pubsub
RUN pip install flask
RUN pip install pymongo
RUN pip install pyodbc
RUN pip install waitress
RUN pip install python-dotenv
EXPOSE 5000
CMD ["python", "Api.v2.py"]