package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleConnection(connection net.Conn) {
	fmt.Printf("[+] Cliente %s conectado.\n", connection.RemoteAddr().String())
	defer connection.Close()

	reader := bufio.NewReader(connection)

	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao ler o nome do arquivo: %w\n", err))
		return
	}

	filename = strings.TrimSpace(filename)
	if filename == "" {
		fmt.Printf("[!] Nome do arquivo recebido está vazio\n")
		return
	}

	filename = filepath.Base(filename)
	fmt.Printf("[*] Recebendo arquivo: %s\n", filename)

	file, err := os.Create(fmt.Sprintf(filename))
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao criar o arquivo: %w\n", err))
		return
	}
	defer file.Close()

	bytesCopied, err := io.Copy(file, reader)
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao copiar dados: %w", err))
		return
	}

	fmt.Printf("[+] Arquivo recebido de %s (%d bytes).\n", connection.RemoteAddr().String(), bytesCopied)
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao iniciar o servidor: %w", err))
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("[*] Servidor na escuta na porta 8080\n")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(fmt.Errorf("[!] Falha ao aceitar conexão: %w", err))
			continue
		}
		go handleConnection(connection)
	}
}
