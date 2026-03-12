# Go Mux MongoDB User Manager API

## Descrição

Esta é uma API RESTful para gerenciamento de usuários, construída com Go, Mux, e MongoDB. O projeto foi desenvolvido como um estudo, aplicando conceitos de Clean Architecture e Domain-Driven Design (DDD) em um CRUD simples.

## Pontos Fortes

* **Estrutura do Projeto:** A API segue uma estrutura de projeto bem organizada, separando as responsabilidades em camadas de domínio, infraestrutura, casos de uso e web, refletindo os princípios da Clean Architecture.
* **Foco no Domínio:** O código demonstra uma tentativa de modelar o domínio de usuários de forma clara e concisa, um princípio fundamental do DDD.
* **Código Limpo e Legível:** O código é bem escrito e fácil de entender, com nomes de variáveis e funções claros e concisos.
* **Segurança:** A API utiliza hashing de senhas para proteger as informações dos usuários.
* **Configuração Flexível:** A configuração da API é feita através de variáveis de ambiente, o que a torna flexível e fácil de implantar em diferentes ambientes.

## Principais Funcionalidades

* **Criar Usuário:** Cria um novo usuário com nome, e-mail e senha.
* **Listar Usuários:** Retorna uma lista com todos os usuários cadastrados.
* **Login:** Autentica um usuário com e-mail e senha.
* **Atualizar Nome:** Atualiza o nome de um usuário.
* **Atualizar E-mail:** Atualiza o e-mail de um usuário.
* **Atualizar Senha:** Atualiza a senha de um usuário.
* **Deletar Usuário:** Deleta um usuário pelo e-mail.

## Pontos de Melhoria (Focados em Estudos de Clean Architecture e DDD)

* **Riqueza do Modelo de Domínio:** O modelo de domínio atual é um pouco anêmico. Para aprofundar os estudos em DDD, poderíamos enriquecer o modelo com mais regras de negócio, validações e comportamentos, em vez de deixar essa lógica apenas nos casos de uso ou na camada de aplicação.
* **Events Sourcing e CQRS:** Para um estudo mais avançado de DDD, poderíamos explorar os padrões de Event Sourcing e CQRS (Command Query Responsibility Segregation). Isso nos permitiria ter um histórico completo de todas as alterações no estado da aplicação e otimizar as operações de leitura e escrita.
* **Testes de Unidade do Domínio:** Embora o projeto tenha testes de unidade, poderíamos focar em escrever mais testes para a camada de domínio, garantindo que todas as regras de negócio estejam cobertas.
* **Injeção de Dependência:** A injeção de dependência é feita manualmente no `main.go`. Para um projeto maior, poderíamos estudar o uso de um container de injeção de dependência para gerenciar as dependências de forma mais eficiente.
* **Value Objects:** O modelo de domínio poderia ser aprimorado com o uso de Value Objects para representar conceitos como E-mail, Senha e Nome. Isso tornaria o código mais expressivo e seguro.
* **Aggregates:** Para um domínio mais complexo, poderíamos agrupar as entidades relacionadas em agregados, garantindo a consistência das transações.
* **Documentação da Arquitetura:** Para um projeto de estudo, seria interessante documentar as decisões de arquitetura tomadas, explicando como os princípios da Clean Architecture e do DDD foram aplicados.
