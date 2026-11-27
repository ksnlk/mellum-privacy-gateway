#!/bin/bash

ab -n 20000 -c 100 -H "Authorization: Bearer dHcmos44DFYC25v9spy86XN1T2IyXv5jvO70BJgQ1FopO5yPn81uxwo0zevHxVaz" \
  http://localhost:3000/api/dashboards/db/dash1
