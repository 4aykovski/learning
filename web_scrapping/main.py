from urllib import request
import requests

if __name__ == '__main__':
    url = 'https://www.python.org'
    # response = request.urlopen(url)
    # print(response.read())
    response = requests.get(url)
    with open('test.html', 'wb') as f:
        f.write(response.content)
