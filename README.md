# dnc-client

The DNC client allows you to stream your music and videos from your computer so they can be accessible everywhere. The computer is used as a server and the website acts as an intermediary for the content.

Currently, just run go install in the folder directory where main.go is located. There should be a config file generated along with an exe file. Run the exe and there should be a pop up of localhost and default port 3000.

The shared folder directory along with the port can be changed inside the config file.

#The Folder Structure

The dnc-client is broken into four parts: helper, portal, main, and router. 

Helper.go contains the functions necessary for the admin portal to run. It has code that searches and lists all the shared files, filters the extensions, user functions and etc.

Portal.go is the templating for the admin portal with the redirect functions needed for sign-up, verification, and log-in.

Router.go routes data to the main web server that handles the user sign up other relevant details. It also gets the folder from the config file and serves it onto the web.

There is also a test config file and a router_test file to test various aspects of the server.

Main.go is the file that should be go installed and then run.





