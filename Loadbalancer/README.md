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

### Summary of the Algorithm

The load balancer uses a round-robin algorithm to distribute incoming HTTP requests across multiple backend servers. The algorithm works as follows:

1. **Initialization**:
    - The `RoundRobin` struct is initialized with a list of backend servers.
    - A global counter (`count`) is used to keep track of the number of requests.

2. **Request Handling**:
    - For each incoming request, the `handleRequest` function is called.
    - The `RoundRobin` method selects the next server in the list using the formula `count % size`, where `size` is the number of servers.
    - The selected server's name is used to determine the corresponding port.
    - The request is forwarded to the selected backend server.
    - The response from the backend server is read and sent back to the client.

3. **Backend Server**:
    - Two backend servers are started on ports 8081 and 8082.
    - Each backend server responds to incoming requests with a message indicating which server is responding.

This approach ensures that incoming requests are evenly distributed across the backend servers in a circular order.
