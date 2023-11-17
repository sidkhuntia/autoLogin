# AutoLogin
- The login page can be accessed by `http://192.168.249.1:1000/login?`
- The source code (HTML file) has a form with hidden inputs along with inputs for `username` and `password` 
- The hidden inputs are:
	1. `magic` - a dynamically generated token
	2. `4Tredir` - its the login url, it is static
	3. ` ` - a blank input with blank value, can be ignored while making `POST` request
- The scrapping cant not be done using `bs4` in python cause the magic token is being generated dynamically with the help Javascript, thats why we need to use selenium along with webDriver 
- Using of selenium requires webDriver to be installed, which can be a hassle to install, and there is also multiple Browser issue, that is why it will be better to achieve the auto login with the help of `Go` along with scraping package [colly](https://github.com/gocolly/colly)
- The challenge is to make it automatic.

<hr>

### Things to do 
- [ ] Detect if connected to college network
- [x] Detect if connected to internet
- [x] Run app if connected to college network but not to internet
- [x] Check of login failure or success
- [ ] Handling of credentials with flags or some better method
- [ ] Making it run automatically
- [ ] Making it startup after boot up
