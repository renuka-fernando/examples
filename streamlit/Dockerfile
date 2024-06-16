FROM python:3.9-slim

RUN pip3 install streamlit

USER 10014

WORKDIR /app

COPY MyApp.py MyApp.py 

EXPOSE 8501

# ENTRYPOINT ["streamlit", "run", "MyApp.py", "--server.port=8501", "--server.address=0.0.0.0"]
ENTRYPOINT ["streamlit", "hello", "--server.port=8501", "--server.address=0.0.0.0"]
