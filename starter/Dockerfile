# Team submission image. The scorer runs:
#   docker run <image> sanitize --in /in --out /out --map /map/mapping.json
#   docker run <image> restore  --in /in --out /out --map /map/mapping.json
FROM python:3.12-slim
WORKDIR /app
COPY . /app
# RUN pip install --no-cache-dir -r requirements.txt
ENTRYPOINT ["python", "/app/gateway.py", "--detector", "regex"]
