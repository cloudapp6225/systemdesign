# A Simple HTTP Load Balancer Using Round Robin Algorithm

This project implements a simple HTTP load balancer using the round-robin algorithm. The load balancer distributes incoming HTTP requests to a list of backend servers in a circular order.

## Getting Started

### Prerequisites

- Go 1.22.5 or later

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/vk-NEU7/systemdesign.git
    cd systemdesign/Loadbalancer
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Running the Load Balancer

1. Start the load balancer:

    ```sh
    make run
    ```

2. The load balancer will start an HTTP server on port 80 and two backend servers on ports 8081 and 8082.

### Making Requests

Send HTTP requests to the load balancer at `http://localhost:80`. The load balancer will forward the requests to the backend servers in a round-robin manner.

## Implementation Details

