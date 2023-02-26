package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// tempo de trabalho/foco

 var tempoTrabalho = 25 * time.Minute

 // tempo de descanso

 var tempoDescanso = 5 * time.Minute


 // inicia o cronometro pomodoro

 func startCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "start",
        Short: "Inicia o cronômetro Pomodoro",
        Run: func(cmd *cobra.Command, args []string) {
            for {
                // tempo trabalhando/focado
                fmt.Println("Comece a trabalhar/focar")
                time.Sleep(tempoTrabalho)

                // tempo de descanso
                fmt.Println("Começo do descanso")
                time.Sleep(tempoDescanso)
            }
        },
    }
}

func mainCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "pomodoro",
        Short: "Um cronômetro Pomodoro em Golang",
    }
    cmd.AddCommand(startCmd())
    return cmd
}

// Chama a função mainCmd() para obter o comando principal e executa o comando
func main() {
    if err := mainCmd().Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
