FROM python:3.8.2
WORKDIR /app
COPY . .
RUN pip3 install --upgrade pip
RUN pip3 install --upgrade setuptools
RUN pip3 install -r requirements.txt
CMD ["python", "main.py"]