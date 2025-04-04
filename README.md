📅 Gerenciador de Agenda
Um sistema de gerenciamento de compromissos desenvolvido em Golang que permite visualizar, reservar e cancelar horários de atendimento.
📋 Funcionalidades

Visualização de Horários Disponíveis: Lista todos os horários que não foram reservados
Visualização de Horários Reservados: Lista todos os horários já reservados, junto com o nome do cliente
Reserva de Horários: Permite reservar um horário disponível
Cancelamento de Reservas: Permite cancelar uma reserva existente

🛠️ Tecnologias Utilizadas

Go (Golang) - Linguagem de programação utilizada
Bibliotecas padrão do Go:

time - Manipulação de datas e horários
bufio - Leitura de entrada do usuário
fmt - Formatação de saída
strconv - Conversão de strings
os - Interação com o sistema operacional

📝 Detalhes de Implementação

Compromisso: Estrutura que armazena informações sobre cada compromisso (ID, hora de início, hora de fim, status de reserva, cliente)
Agenda: Estrutura que gerencia uma coleção de compromissos
Interface de linha de comando: Menu interativo para o usuário gerenciar a agenda

📄 Licença
Este projeto está sob a licença MIT.
👨‍💻 Autor

Saulo Lucena - GitHub
