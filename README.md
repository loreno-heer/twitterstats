# twitterstats
Collect Statistics from Twitter Posts. In particular it counts likes of posts over time.
The reason for creating this app was that there is a strong suspicion that likes on some major posts are systematically deleted or that there is a major unresolved bug which prevents new likes from being recorded.
Some data collected using the app is provided in the example folder. In order to display it use gnuplot or pyplot or similar: You need to sort by tweetid and parse the UTC timestamp properly. 

Usage:

1. Enter your twitter username and password in the login file and run that (might need to create a go module first and build the app). Running the file will save a cookie for the app such that further usage won't require a login.
2. In the second file there is a list of tweet ids that you want to monitor. Edit that as you please. Also create the csv file and if desired change its name. If the file does not exist there will be an error. There is a delay between each lookup, change as desired.
3. Build and run the application. Terminate with ctrl+c.
