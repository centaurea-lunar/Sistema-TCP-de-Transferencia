package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func main() {
	const _filepath = "arquivo.txt"

	file, err := os.Open(_filepath)
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao abrir o arquivo '%v': %w", file, err))
		os.Exit(1)
	}
	defer file.Close()

	fmt.Println("Conectando ao servidor em localhost:8080...")
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao conectar ao servidor: %w", err))
		os.Exit(1)
	}
	defer connection.Close()

	baseFilepath := filepath.Base(_filepath)

	_, err = fmt.Fprintln(connection, baseFilepath)
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao enviar o nome do arquivo: %w", err))
		return
	}
	fmt.Printf("[*] Enviando nome do arquivo: %s\n", baseFilepath)

	fmt.Println("[*] Enviando conteúdo do arquivo...")
	bytesCopied, err := io.Copy(connection, file)
	if err != nil {
		fmt.Println(fmt.Errorf("[!] Falha ao enviar o conteúdo do arquivo: %w", err))
		return
	}

	fmt.Printf("[+] Arquivo enviado com sucesso (%d bytes).\n", bytesCopied)
}
