package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

// upCmd representa el comando 'up'
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Levanta los contenedores y ejecuta los comandos necesarios",
	Long: `Este comando levanta los contenedores con docker-compose, luego ejecuta un comando 
desde un archivo init.sh y finalmente ejecuta otro docker-compose.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ejecutar docker-compose up --build -d
		fmt.Println("Ejecutando docker-compose up --build -d...")
		err := runCommand("docker-compose", "up", "--build", "-d")
		if err != nil {
			fmt.Printf("Error ejecutando docker-compose up: %v\n", err)
			return
		}

		// Esperar un momento para que los contenedores levanten
		fmt.Println("Esperando 5 segundos para que los contenedores estén listos...")
		time.Sleep(10 * time.Second)

		// Ejecutar el comando desde init.sh
		fmt.Println("Ejecutando init.sh...")
		err = runCommand("bash", "init.sh")
		if err != nil {
			fmt.Printf("Error ejecutando init.sh: %v\n", err)
			return
		}

		// Ejecutar docker-compose -f docker-compose-post.yml up --build -d
		fmt.Println("Ejecutando docker-compose-post.yml up --build -d...")
		err = runCommand("docker-compose", "-f", "docker-compose-post.yml", "up", "--build", "-d")
		if err != nil {
			fmt.Printf("Error ejecutando docker-compose-post.yml up: %v\n", err)
			return
		}

		fmt.Println("Comando ejecutado exitosamente")
	},
}

// Función auxiliar para ejecutar comandos
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Aquí puedes definir tus banderas y configuraciones si es necesario.
}
