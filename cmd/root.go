/*
Copyright © 2025 Mates182
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa el comando base cuando se llama sin subcomandos
var rootCmd = &cobra.Command{
	Use:   "multicompose",
	Short: "Una herramienta para ejecutar secuencias de comandos Docker Compose de manera automatizada",
	Long: `Multicompose es una aplicación CLI que facilita la ejecución de múltiples comandos Docker Compose de forma secuencial. 
Este programa permite:

1. Levantar contenedores usando un archivo docker-compose.yml.
2. Ejecutar un script de inicialización (init.sh).
3. Levantar contenedores adicionales con un segundo archivo docker-compose (docker-compose-post.yml).

Es útil para automatizar flujos de trabajo con Docker, especialmente cuando se tienen dependencias entre contenedores que deben levantarse en secuencia.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute agrega todos los comandos hijos al comando raíz y establece las banderas apropiadamente.
// Esto se llama desde main.main(). Solo necesita ejecutarse una vez para el rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Aquí puedes definir tus banderas y configuraciones.
	// Cobra soporta banderas persistentes que, si se definen aquí,
	// serán globales para toda tu aplicación.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "archivo de configuración (el valor predeterminado es $HOME/.multicompose.yaml)")

	// Cobra también soporta banderas locales, que solo se ejecutarán
	// cuando esta acción sea llamada directamente.
	rootCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
