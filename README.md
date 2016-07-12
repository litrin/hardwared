# hardwared
This service will provide a Restful API which may help you get hardware details for you hosts.

# API definition

- Path: /v1/hosts/
- Action: GET
- Response:
    {
      "result": {
          "status": "success",
          "errMsg": "xxxxx",
          "data": {
              "hosts": [
                  {
                      "id": 0,
                      "status": "healthy",
                      "hostName": "slave1",
                      "cpus": [
                          {
                              "id": 0,
                              "status": "healthy",
                              "model": "xxxx",
                              "frequency": "xxxx",
                              "coreNum": 20
                          }
                      ],
                      "memorys": [
                          {
                              "id": 0,
                              "status": "healthy",
                              "capacity": "32G",
                              "frequency": "2133MHz"
                          }
                      ],
                      "NIC": [
                          {
                              "id": 0,
                              "status": "healthy",
                              "ip": "192.168.1.1",
                              "mac": "xxxx",
                              "speed": "1000MB"
                          }
                      ],
                      "pcieConnections": [
                          {
                              "id": 0,
                              "status": "healthy",
                              "cableId": "xxxxx"
                          }
                      ]
                  }
              ]
          }
      }
    }
  
    
