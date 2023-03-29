<div id="top"></div>

<!-- Structure Demo Spring -->
<br />
<div align="center">
  <a href="https://github.com/Sholontla">
  </a>

<h3 align="center">Admin/Inventory/Stores/Service in Go</h3>

  <p align="center">
   This Structure Demo is use for testing and demostration to understad, the full Demo Structure please visit the diagram: https://github.com/Sholontla/structure-demo-configurations/blob/main/diagram-demo.pdf
    <br />
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About The Project

The main project are structur by:
Admin Service:

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- Java
- Spring

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Services:

1. Config Server:

- First service to run.
- This service is Cloud Native connections

2. Eureka Discovery Server:

- Second service to run.
- Eureka Discovery Servers

3. Spring GateWay:

- Third service to run.
- manage port services assign by Eureka Spring Server
- port: 6000

4. Security Service:

- Fourth service to run.
- AES Encryption / Obfuscation service
- Data filter by group
- Data Base: Postgre Sql
- port: manage by Eureka Server

5. configuration Data Base:

- Fifth service to run.
- Configuration of Databases
- Data Base: Local
- port: manage by Eureka Server

6. Secure Vault Demo:

- Sixth service to run.
- Keep Secure and confidentail data
- Data Base: Mongo DB
- port: manage by Eureka Server

### Endpoints and Swagger

- localhost:6000/security-ws/security
- localhost:6000/security-ws/security/login
- localhost:6000/security-ws/security/status/check
- localhost:6000/security-ws/security/encryption
- localhost:6000/security-ws/security/obfuscation
- localhost:6000/security-ws/security/{id}
- localhost:6000/security-ws/security/group?group=(select group: 0,1,2,3,4,5,6,7,8,9)
- localhost:6000/security-ws/security/actuator/\*\*
- localhost:6000/security-ws/actuator/circuitbreakerevents
-
- To see the Swagger of: secure-vault-demo
- on Eureka Server the assign ports with out any "http client" go to: localhost:8000 and select the "SECURE-VAULT-DEMO-WS" service.
- change default route and enter:
- http://localhost:[port manage by Eureka Server]/swagger-ui/index.html

### Prerequisites

Virtualization / Containers

- Docker
- Docker - Compose

O/I

- Windows
- Linux

## License

For testing and demostrations purposes.

<!-- CONTACT -->

## Contact

Gerardo Ruiz Bustani - solbustani@gmail.com - 442 488 6193

Project Link: [https://github.com/Sholontla](https://github.com/Sholontla/structure-demo)
