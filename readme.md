# Image Server

This is a simple Go program that sets up a web server to serve images. It allows you to resize images dynamically based on query parameters.

## Getting Started

### Prerequisites

To run this program, you need to have Go installed on your machine.

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/asikarwar007/image-server.git

2. Change to the project directory:
    
    ```shell
    cd image-server

3. Build the executable:
    
    ```shell
    go build

4. Run the program:

    ```shell
    ./image-server

The server will start listening on port 8080. You can access it by opening a web browser and navigating to http://localhost:8080/.

### Usage

Once the server is running, you can access the images using the following URL pattern:

    http://localhost:8080/image-name.jpg?w=300&h=200&q=80

- image-name.jpg should be replaced with the name of the image file you want to retrieve.
- w (optional) specifies the desired width of the image.
- h (optional) specifies the desired height of the image.
- q (optional) specifies the desired quality of the image.


## License

This project is licensed under the MIT License.

- Feel free to customize the content and sections as per your project's requirements. Don't forget to include a license file (e.g., `LICENSE`) in your repository and update the link in the README accordingly.