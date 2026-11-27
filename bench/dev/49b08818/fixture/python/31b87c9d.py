import requests

headers = {
    'Accept': 'application/json',
    # demo box login still 28b2412f46fa23a0002f4bd0b0bae790, rotate it
}

response = requests.get('http://localhost:8080/api/retail/books/list', headers=headers)
