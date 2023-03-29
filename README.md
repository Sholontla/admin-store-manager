<div id="top"></div>

<!-- Structure Demo Spring -->
<br />
<div align="center">
  <a href="https://github.com/Sholontla">
  </a>

<h3 align="center">Admin/Inventory/Stores/Service in Go</h3>

  <p align="center">
   This Structure Demo is use for testing and demostration to understad, the full Demo Structure please visit the diagram: https://github.com/Sholontla/admin-store-manager/blob/main/arch-high-level-overview.pdf
    <br />
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->

## About The Project

The project are structure by:
Admin Service:

  <p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- Golang (Go)
- Fiber (http framework)
- JWT/v4
- MongoDB

Project OverView:
. The Adimin Services will take care of the Administration, Configuration, CRUD operations on AdminUsers.
. From the Stores side will take care of the creation / Destroy of the store, generate and renew the correct certificates for a new store or existing store, create te Admin User for every store and aver employee per store.
. All the logs will be send to the Logger service and with the help of grafana, prometheus and protmail will have metrics and controll of the logs.

Store Service:

  <p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- Golang (Go)
- Fiber (http framework)
- JWT/v4
- MongoDB

Project OverView:
. The store service will take care of the sales/cashier oprtations and will save the raw data into a cassandra cluster and send throw a Kafka producer the raw data to the Python Analysis services where a consumer will take care from precessing and apply analytic logic to the data and shows into a Dashboard (Html and bootrstrap with the help of Chart.js) and generate the visualizations.
Once the data is processed will be saved into PostgreSQL.
. The data consumer will have buffer messages, if somethig happen will buffer the data until the consumr come back online.
. The dat will store in a Redis docker container and all the messages will be cashed for the data constancy.
. The inventory will be send throw a gRPC server from maain INventory to every store depending on every store and the inventory data as well will be chashed.
All the logs will be send to the Logger service and with the help of grafana, prometheus and protmail will have metrics and controll of the logs.

Logger Service:

  <p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- Golang (Go)
- Fiber (http framework)
- JWT/v4
- MongoDB

Project OverView:
. The store service will take care of the sales/cashier oprtations and will save the raw data into a cassandra cluster and send throw a Kafka producer the raw data to the Python Analysis services where a consumer will take care from precessing and apply analytic logic to the data and shows into a Dashboard (Html and bootrstrap with the help of Chart.js) and generate the visualizations.
Once the data is processed will be saved into PostgreSQL.
. The data consumer will have buffer messages, if somethig happen will buffer the data until the consumr come back online.
. The dat will store in a Redis docker container and all the messages will be cashed for the data constancy.
. The inventory will be send throw a gRPC server from maain INventory to every store depending on every store and the inventory data as well will be chashed.
All the logs will be send to the Logger service and with the help of grafana, prometheus and protmail will have metrics and controll of the logs.

Inventory Service:

  <p align="right">(<a href="#top">back to top</a>)</p>

### Built With

- Golang (Go)
- Fiber (http framework)
- JWT/v4
- MongoDB

Project OverView:
. The store service will take care of the sales/cashier oprtations and will save the raw data into a cassandra cluster and send throw a Kafka producer the raw data to the Python Analysis services where a consumer will take care from precessing and apply analytic logic to the data and shows into a Dashboard (Html and bootrstrap with the help of Chart.js) and generate the visualizations.
Once the data is processed will be saved into PostgreSQL.
. The data consumer will have buffer messages, if somethig happen will buffer the data until the consumr come back online.
. The dat will store in a Redis docker container and all the messages will be cashed for the data constancy.
. The inventory will be send throw a gRPC server from maain INventory to every store depending on every store and the inventory data as well will be chashed.
All the logs will be send to the Logger service and with the help of grafana, prometheus and protmail will have metrics and controll of the logs.

Frameworks, toools, and other components add to this project:

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
