## distribution_centers


### endpoints
- GET `/distributioncenters?itemId=123` by id produto para consultar em quais CDs esse item se encontra.
```
# Response
{
    ¨distributionCenters¨: ["CD1", "CD2", "CD3"]
}

```

### ideias iniciais
- Colocar o relacionamento entre CD e Produto no Dynamo onde o id de ambos compõe a pk pra facilitar um pouco a consulta.
- Ter uma propriedade de Estado, para vir sempre do mais próximo do cliente.

### ideias futuras ou opcionais
- Ter uma propriedade de Zona, para se caso o produto esteja em mais de 1 CD dentro do mesmo Estado, que venha da Zona mais proxima.
- Ter uma proprierade de Quantidade de Produto, para se caso existem 2 CDs dentro de uma Zona em um mesmo estado, requisitar do CD com mais Quantidades do produto para equilibrar o armazém.

### Todo 
- [ ] Criar o GET `/distributioncenters?itemId=123`
- [ ] Criar banco de dados com a estrutura dos Produtos vs CDs
- [ ] Criar repositório dos produtos via CD
- [ ] Usar Hexagonal Arch
- [ ] Documentar a Api com Swagger
- [ ] Infra do código (docker, makefile etc) 
- [ ] Criar testes unitarios
- [ ] Criar testes de integração (Insomnia e se der tempo Serenity+Cucumber ou Robot)
- [ ] Criar o start da aplicação pelo cli usando o Cobra
- [ ] Fazer o setup da suite de qualidade com o Sonarqube
- [ ] Usar os linters do Golang na pipeline de cada projeto

### Todo futuro ou opcional
- [ ] Criar o CRUD de `/products` que receba um CD.
- [ ] Criar o CRUD de `/distributioncenters` que tenha critérios geográficos.