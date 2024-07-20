import socket
import time

def handle_client_connection(client_socket):
    client_socket.send(b"HTTP/1.1 200 OK\r\n")
    client_socket.send(b"Content-Length: 100\r\n")
    client_socket.send(b"Content-Type: text/plain\r\n\r\n")
    client_socket.send(b"Partial response...")
    time.sleep(1)
    client_socket.close()  # Abruptly close the connection

def run_server():
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind(('0.0.0.0', 8080))
    server_socket.listen(5)
    print("Backend server running on port 8080...")
    while True:
        client_socket, addr = server_socket.accept()
        handle_client_connection(client_socket)

if __name__ == "__main__":
    run_server()
