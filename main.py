from selenium import webdriver
from selenium.webdriver.firefox.options import Options as FirefoxOptions
from selenium.webdriver.chrome.options import Options as ChromeOptions
from html.parser import HTMLParser
import requests
import os
from dotenv import load_dotenv

load_dotenv()
class MyHTMLParser(HTMLParser):
    def __init__(self):
        super().__init__()
        self.hidden_inputs = {}

    def handle_starttag(self, tag, attrs):
        if tag == 'input':
            attrs = dict(attrs)
            if attrs.get('type') == 'hidden':
                self.hidden_inputs[attrs['name']] = attrs['value']

# Initialize the WebDriver in headless mode
options = ChromeOptions()  # or ChromeOptions(), depending on which browser you want to use
options.add_argument("--headless")

driver = webdriver.Chrome(options=options)  # or webdriver.Chrome(options=options)

url = 'http://192.168.249.1:1000/login?'
driver.get(url)

page_source = driver.page_source

parser = MyHTMLParser()
parser.feed(page_source)

hidden_inputs = parser.hidden_inputs

USERNAME = os.getenv('USERNAME')
PASSWORD = os.getenv('PASSWORD')

form_data = {
    'username': USERNAME,
    'password': PASSWORD,
    **hidden_inputs  
}

response = requests.post(url, data=form_data)

print(response.content)
