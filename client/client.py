import socket
import sys

if (len(sys.argv) == 1):
    print("[!] Número incorreto de argumentos. A execução do script deve seguir o formato \"python script.py client_ip file_name\"")
    sys.exit()

# Configurações do cliente
CLIENT_IP = sys.argv[1]
CLIENT_PORT = 8080
BUFFER = 4096
FILE_NAME = sys.argv[2]

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client:
    client.connect((CLIENT_IP, CLIENT_PORT))
    print(f"[*] Conectado ao servidor {CLIENT_IP}:{CLIENT_PORT}")

    client.send(FILE_NAME.encode())

    with open(FILE_NAME, 'rb') as archive:
        while True:
            data = archive.read(BUFFER)
            if not data:
                break
            client.send(data)

    print(f"[+] Arquivo '{FILE_NAME}' enviado com sucesso!")