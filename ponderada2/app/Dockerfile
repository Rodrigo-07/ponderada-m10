    FROM python:3.12-slim

    ENV PYTHONUNBUFFERED=1
    WORKDIR /code
    COPY requirements.txt .
    RUN pip install --no-cache-dir -r requirements.txt

    COPY . .
    # CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080", "--workers", "1"]
    # CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080"]
    CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080", "--loop", "uvloop", "--http", "h11", "--keep-alive", "120", "--backlog", "2048"]