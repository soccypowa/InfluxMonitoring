# InfluxMonitor
Usage: 
```sh
influxmonitor `<healthendpoint>` `<emailaddress>`
```
Example: 
```sh
influxmonitor http://localhost:8086/health my@email.com
```

Simple monitor to check the availability of the influxdb deamon through a GET request to the health enpoint, the status code and the content of the message is rudimentary checked. In the case of an error an email is sent to the appointed email address.