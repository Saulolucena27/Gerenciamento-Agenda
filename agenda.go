package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Compromisso struct {
	ID        int
	HoraInicio time.Time
	HoraFim   time.Time
	Reservado bool
	Cliente   string
}

type Agenda struct {
	Compromissos []Compromisso
	ProximoID    int
}

func NovaAgenda(horaInicio, horaFim, intervaloMinutos int) *Agenda {
	agenda := &Agenda{
		Compromissos: []Compromisso{},
		ProximoID:    1,
	}
	
	// Data atual
	agora := time.Now()
	hoje := time.Date(agora.Year(), agora.Month(), agora.Day(), 0, 0, 0, 0, agora.Location())

	for hora := horaInicio; hora < horaFim; hora++ {
		for minuto := 0; minuto < 60; minuto += intervaloMinutos {
			if hora == horaFim && minuto > 0 {
				break
			}
			
			horaInicio := hoje.Add(time.Hour*time.Duration(hora) + time.Minute*time.Duration(minuto))
			horaFim := horaInicio.Add(time.Minute * time.Duration(intervaloMinutos))

			compromisso := Compromisso{
				ID:        agenda.ProximoID,
				HoraInicio: horaInicio,
				HoraFim:   horaFim,
				Reservado:  false,
				Cliente:    "",
			}
			
			agenda.Compromissos = append(agenda.Compromissos, compromisso)
			agenda.ProximoID++
		}
	}
	
	return agenda
}

func (a *Agenda) ExibirCompromissosDisponiveis() {
	fmt.Println("\n=== HORÁRIOS DISPONÍVEIS ===")
	disponiveis := false

	for _, comp := range a.Compromissos {
		if !comp.Reservado {
			fmt.Printf("[%d] %s - %s\n", 
				comp.ID, 
				comp.HoraInicio.Format("15:04"), 
				comp.HoraFim.Format("15:04"))
			disponiveis = true
		}
	}

	if !disponiveis {
		fmt.Println("Não há horários disponíveis.")
	}
	fmt.Println("===========================")
}

func (a *Agenda) ExibirHorariosReservados() {
	fmt.Println("\n=== HORÁRIOS RESERVADOS ===")
	reservados := false

	for _, comp := range a.Compromissos {
		if comp.Reservado {
			fmt.Printf("[%d] %s - %s | Cliente: %s\n", 
				comp.ID, 
				comp.HoraInicio.Format("15:04"), 
				comp.HoraFim.Format("15:04"),
				comp.Cliente)
			reservados = true
		}
	}

	if !reservados {
		fmt.Println("Não há horários reservados.")
	}
	fmt.Println("===========================")
}

func (a *Agenda) ReservarCompromisso(id int, nomeCliente string) bool {
	for i := range a.Compromissos {
		if a.Compromissos[i].ID == id && !a.Compromissos[i].Reservado {
			a.Compromissos[i].Reservado = true
			a.Compromissos[i].Cliente = nomeCliente
			return true
		}
	}
	return false
}

func (a *Agenda) CancelarCompromisso(id int) bool {
	for i := range a.Compromissos {
		if a.Compromissos[i].ID == id && a.Compromissos[i].Reservado {
			a.Compromissos[i].Reservado = false
			a.Compromissos[i].Cliente = ""
			return true
		}
	}
	return false
}

func ExibirMenu() {
	fmt.Println("\n===== GERENCIADOR DE AGENDA =====")
	fmt.Println("1. Ver horários disponíveis")
	fmt.Println("2. Ver horários reservados")
	fmt.Println("3. Reservar horário")
	fmt.Println("4. Cancelar reserva")
	fmt.Println("0. Sair")
	fmt.Print("Escolha uma opção: ")
}

func main() {
	agenda := NovaAgenda(8, 18, 30)
	leitor := bufio.NewReader(os.Stdin)

	for {
		ExibirMenu()
		
		entrada, _ := leitor.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		opcao, err := strconv.Atoi(entrada)
		if err != nil {
			fmt.Println("Opção inválida. Tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			agenda.ExibirCompromissosDisponiveis()

		case 2:
			agenda.ExibirHorariosReservados()

		case 3:
			fmt.Println("\n== Reservar Horário ==")
			agenda.ExibirCompromissosDisponiveis()

			if len(agenda.Compromissos) == 0 {
				fmt.Println("Não há horários disponíveis.")
				continue
			}

			fmt.Print("Digite o ID do horário que deseja reservar: ")
			idEntrada, _ := leitor.ReadString('\n')
			idEntrada = strings.TrimSpace(idEntrada)

			id, err := strconv.Atoi(idEntrada)
			if err != nil {
				fmt.Println("ID inválido. Tente novamente.")
				continue
			}
			
			fmt.Print("Digite o nome do cliente: ")
			nomeCliente, _ := leitor.ReadString('\n')
			nomeCliente = strings.TrimSpace(nomeCliente)
			
			if nomeCliente == "" {
				fmt.Println("Nome do cliente não pode estar vazio.")
				continue
			}

			if agenda.ReservarCompromisso(id, nomeCliente) {
				fmt.Println("Horário reservado com sucesso!")
			} else {
				fmt.Println("Não foi possível reservar o horário.")
			}
			
		case 4:
			fmt.Println("\n== Cancelar Reserva ==")
			agenda.ExibirHorariosReservados()

			fmt.Print("Digite o ID do horário que deseja cancelar: ")
			idEntrada, _ := leitor.ReadString('\n')
			idEntrada = strings.TrimSpace(idEntrada)

			id, err := strconv.Atoi(idEntrada)
			if err != nil {
				fmt.Println("ID inválido. Tente novamente.")
				continue
			}
			
			if agenda.CancelarCompromisso(id) {
				fmt.Println("Reserva cancelada com sucesso!")
			} else {
				fmt.Println("Não foi possível cancelar a reserva.")
			}
			
		case 0:
			fmt.Println("Saindo do programa. Até logo!")
			return
			
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}