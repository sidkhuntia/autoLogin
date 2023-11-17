
#!/bin/bash

# Download the appropriate version of the webdriver for the installed Chrome browser
CHROME_VERSION=$(google-chrome --version | awk '{print $3}' | cut -d '.' -f 1)
WEBDRIVER_VERSION=$(curl -s https://chromedriver.storage.googleapis.com/LATEST_RELEASE_$CHROME_VERSION)
wget https://chromedriver.storage.googleapis.com/$WEBDRIVER_VERSION/chromedriver_linux64.zip

# Move the webdriver to the appropriate directory
unzip chromedriver_linux64.zip
sudo mv chromedriver /usr/local/bin/

# Set the PATH environment variable to include the directory containing the webdriver
echo "export PATH=$PATH:/usr/local/bin" >> ~/.bashrc
source ~/.bashrc
