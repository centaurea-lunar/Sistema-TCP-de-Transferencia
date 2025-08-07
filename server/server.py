import socket

# Configurações do servidor
SERVER_IP = '0.0.0.0'
SERVER_PORT = 8080
BUFFER = 4096

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server:
    server.bind((SERVER_IP, SERVER_PORT))
    server.listen(1)
    print(f"[*] Servidor aguardando conexões em {SERVER_IP}:{SERVER_PORT}...")

    client, address = server.accept()
    print(f"[+] Conexão estabelecida com {address}")   

    file_name = client.recv(BUFFER).decode()
    print(f"[*] Recebendo arquivo: {file_name}")

    with open(file_name, 'wb') as archive:
        while True:
            data = client.recv(BUFFER)
            if not data:
                break
            archive.write(data)

    print(f"[+] Arquivo '{file_name}' recebido com sucesso!")
