# Service Monitor
Monitor OpenStack services using REST API calls

## Dependencies ##
This program does require the jsonparser package (https://github.com/buger/jsonparser). Hence, download the package by using the following command:

    go get github.com/buger/jsonparser
    
## Installation Instructions ##

This section describes the following steps to execute this program

1.Download the source code from Github

    git clone https://github.com/RajsimmanRavi/service_monitor.git /path/to/go/src/

2.Edit service_monitor.go file to insert const values

    const COMPUTE_PORT = "xxx"
    const USER_NAME = "xxx"
    const PASSWORD = "xxx"
    
    In the main() function, edit the TENANTS
    
    var TENANTS = [x]string{"xxx", "xxx", "xxx" ...}

3.Build the go file

    go build service_monitor.go
    
4.Execute the file

    ./service_monitor
    
## Contact

If there is any questions, contact me: rajsimmanr@savinetwork.ca


    


