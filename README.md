# Project Name
DnC Media - Client Server

#Project Overview

DnC Media is an open-source media server solution. This client server is part of the [DnC Media suite.](https://github.com/DiamondNotCrush/dnc-web) and contains the client server application necesary to serve the user's media files. The client server is written in Go.

#Usage

The client server is installed on the users's computer which hosts their media files. Once the server is started, a broswer window opens which allows the user to do some initial but simple setup. This admin portal allows the user to signup, login, and set the folder that contains their media.

##Development

###Installing Dependencies

The client is written in Go which can be downloaded from [golang.org/doc/install](https://golang.org/doc/install)

Since the client is written in Go, the command '''go get''' folowed by the following packages will install the necessary dependencies for development.
  1. github.com/gorilla/mux
  1. github.com/skratchdot/open-golang/open
  1. github.com/codeskyblue/go-sh

###Structure

This application follows the below structure:
```
src/
  github.com/
    *dependencies*/
    DiamondNotCrush/
      dnc-client/
        main/
          main.go
        portal/
          portalTemplates/
            portalTemplates.go
          portal.go
        router/
          admin/
            admin.go
          info/
            info.go
          share/
            share.go
          test/
          config (for router_test)
          router.go
          router_test.go
```
###Building

The application is built by running ```go install``` inside the main/ directory. The executable is created in the bin/ folder that should be located in the same directory as the src/ root of the application.
