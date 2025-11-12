package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Print("Ingrese el puerto a liberar: ")
	inputReader := bufio.NewReader(os.Stdin)

	portToFree, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error al leer la entrada: %v\n", err)
		return
	}

	portToFree = strings.TrimSpace(portToFree)

	if portToFree == "" {
		fmt.Println("No se ingresó ningún puerto.")
		return
	}

	fmt.Printf("Intentando liberar el puerto: %s...\n", portToFree)

	switch runtime.GOOS {
	case "windows":
		killProcessOnPortWindows(portToFree)
	case "linux", "darwin":
		killProcessOnPortUnix(portToFree)
	default:
		fmt.Printf("Sistema operativo no soportado: %s\n", runtime.GOOS)
	}
}

func killProcessOnPortUnix(port string) {
	findPidCommand := fmt.Sprintf("lsof -t -i:%s -sTCP:LISTEN", port)
	command := exec.Command("sh", "-c", findPidCommand)

	commandOutput, err := command.CombinedOutput()
	pidOutput := strings.TrimSpace(string(commandOutput))

	if err != nil || pidOutput == "" {
		fmt.Printf("No se encontró ningún proceso escuchando (LISTEN) en el puerto %s.\n", port)
		fmt.Println("Nota: Si el puerto está ocupado, intente ejecutar esta herramienta con 'sudo'.")
		return
	}

	processIds := strings.Fields(pidOutput)

	if len(processIds) == 0 {
		fmt.Printf("No se encontraron PIDs válidos en la salida de lsof.\n")
		return
	}

	fmt.Printf("Proceso(s) LISTENING encontrado(s) con PID(s): %s. Intentando finalizar...\n", strings.Join(processIds, ", "))

	killArgs := []string{"-9"}
	killArgs = append(killArgs, processIds...)

	killCommand := exec.Command("kill", killArgs...)
	killOutput, err := killCommand.CombinedOutput()
	if err != nil {
		fmt.Printf("Error al intentar finalizar el/los proceso(s) %s: %v\n%s\n", strings.Join(processIds, ", "), err, string(killOutput))
		return
	}

	fmt.Printf("Proceso(s) %s finalizado(s) exitosamente.\n", strings.Join(processIds, ", "))
	fmt.Println(string(killOutput))
}

func killProcessOnPortWindows(port string) {
	findPidCommand := fmt.Sprintf("netstat -ano | findstr :%s | findstr LISTENING", port)
	command := exec.Command("cmd", "/C", findPidCommand)

	commandOutput, err := command.CombinedOutput()
	outputLine := strings.TrimSpace(string(commandOutput))

	if err != nil || outputLine == "" {
		fmt.Printf("No se encontró ningún proceso escuchando (LISTEN) en el puerto %s.\n", port)
		return
	}

	lines := strings.Split(outputLine, "\n")
	var processIds []string

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			fmt.Printf("Error al parsear la salida de netstat: %s\n", line)
			continue
		}
		processIds = append(processIds, fields[len(fields)-1])
	}

	if len(processIds) == 0 {
		fmt.Printf("No se pudieron extraer PIDs válidos de la salida de netstat.\n")
		return
	}

	fmt.Printf("Proceso(s) LISTENING encontrado(s) con PID(s): %s. Intentando finalizar...\n", strings.Join(processIds, ", "))

	killArgs := []string{"/F"}
	for _, pid := range processIds {
		killArgs = append(killArgs, "/PID", pid)
	}

	killCommand := exec.Command("taskkill", killArgs...)
	killOutput, err := killCommand.CombinedOutput()
	if err != nil {
		fmt.Printf("Error al intentar finalizar el/los proceso(s) %s: %v\n%s\n", strings.Join(processIds, ", "), err, string(killOutput))
		return
	}

	fmt.Printf("Proceso(s) %s finalizado(s) exitosamente.\n", strings.Join(processIds, ", "))
	fmt.Println(string(killOutput))
}
