# EmailN
## Domain

Um serviço para envio de email em lote. O objetivo é facilitar, para o sistema cliente, não perder tempo enviando email a email. 

#### Features

- Endpoint para envio dos e-mails em lote com os contatos
- Endpoint para informar o status do lote enviado.
- Os e-mails enviados poderão ser personalizados por contato, com isso, informações do contato poderão estar no texto do e-mail.

#### Prof: Stephany Henrique Batista 

Este projeto é realizado seguindo os ensinamentos do programador Stephany Henrique Batista.  
Todos os créditos do projeto é dele.  

Meus agradecimentos ao Stephany. 

postgres do meu docker  
user emailn_dev  
senha cebola  
docker run --name postgres -e POSTGRES_PASSWORD=cebola -e POSTGRES_USER=emailn_dev -p 5432:5432 -d postgres  

acabei usando um docker-compose para instalar o pgadmin4 junto com o postgres