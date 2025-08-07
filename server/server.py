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

    nome_arquivo = client.recv(BUFFER).decode()
    print(f"[*] Recebendo arquivo: {nome_arquivo}")

    with open(nome_arquivo, 'wb') as arquivo:
        while True:
            dados = client.recv(BUFFER)
            if not dados:
                break
            arquivo.write(dados)

    print(f"[+] Arquivo '{nome_arquivo}' recebido com sucesso!")