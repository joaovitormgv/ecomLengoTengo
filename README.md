## Sobre o projeto
Esta aplicação web é um esboço de um backend para e-commerce, com lógicas para personalizar a experiência de cada usuário a partir de dados de navegação deste e dos outros usuários do site. Cada usuário deve ter uma experiência diferente, personalizada aos seus interesses em produtos, baseado naquilo que foi mais visualizado e que tem mais probabilidade de se interessar e/ou comprar.

Esse projeto foi realizado para o Desafio Voraz 2024 da GoCase. Desafio proposto pela gocase: criar um ecommerce que personalize a experiência de cada usuário.

## Tecnologias utilizadas
- **Golang**: Linguagem de backend usada para desenvolver o servidor.
- **Fiber**: Framework de Go utilizado no desenvolvimento da aplicação do lado do backend, possibilitando a comunicação com o banco de dados.
- **Postgres**: Banco de dados SQL ideal para aplicações leves.
- **Docker**: Facilita o desenvolvimento e implantação de aplicações em ambientes isolados.

## Instruções para Executar a Aplicação

Certifique-se de ter o Docker Desktop baixado no seu dispositivo e deixe o aplicativo aberto em segundo plano. Você pode baixá-lo [aqui](https://www.docker.com/products/docker-desktop/).

### Rodando a Aplicação

Para rodar a aplicação juntamente com o backend, siga estes passos:

1. No terminal, execute o comando:
   ```bash
   docker compose up
  
2. Em seguida, em outro terminal, execute o comando:
```bash
go run.main.go
```

## Estrutura de arquivos
```plaintext

├── app
│   ├── database
│   ├── handlers
│   ├── middleware
│   ├── models
│   ├── routes
│   └── service
├── config
│   ├── app.yaml
│   └── database.yaml
├── docker-compose.yaml
├── ecomLengoTengo
├── go.mod
├── go.sum
├── main.go
└── sql
    └── setup.sql
```


## Lógica de negócios / Recomendações:
A lógica de negócios está toda nos handlers e no banco de dados. Na parte dos handlers, mais especificamente em handler_recommend.go
Basicamente a lógica consiste em entender quais categorias de produtos o usuário mais visualizou, e recomendar nessa ordem de preferência, os produtos mais vendidos das categorias vistas, e de forma proporcional. 
Exemplo, se a categoria de Eletrônicos foi vistas 2 vezes, e a de roupas apenas 1 vez, serão recomendados produtos das duas, mas preferencialmente da categoria mais vista.

## Mockups de telas: 
Devido ao curto tempo e a outras ocupações dos participantes do desafio, apenas o mockup das telas foi feito no figma, mas o frontend ainda não foi implementado. Posteriormente o projeto pode ser completo para portfolio.
