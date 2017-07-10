# Service Monitor
Monitor OpenStack services using REST API calls

## Dependencies ##
This program requires a couple of Golang packages:

- jsonparser package (https://github.com/buger/jsonparser). Download the package by using the following command
    
        go get github.com/buger/jsonparser
    
- Go-ini package (https://go-ini.github.io/ini). Download the package by using the following command:

        go get github.com/go-ini/ini
    
## Installation Instructions ##

This section describes the following steps to execute this program

1.Download the source code from Github

    git clone https://github.com/RajsimmanRavi/service_monitor.git /path/to/go/src/

2.Edit config.ini file to insert const values such as:
    
    const COMPUTE_PORT = "xxx"
    const USER_NAME = "xxx"
    const PASSWORD = "xxx"

3.Build the go file

    go build service_monitor.go
    
4.Execute the file

    ./service_monitor
    
## Contact

If there is any questions, contact me: rajsimmanr@savinetwork.ca


    


