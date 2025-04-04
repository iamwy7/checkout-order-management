## orders

Esse microservice é o inicio da cadeia, é onde criamos um pedido e eles nos retorna de onde virá cada item do pedido.


### endpoints
- POST `/orders` de pedido com um array de produtos (até 100 itens), que consulta quais os CDs que possuem os produtos, determina o melhor CD para cada um deles, cria no banco de dados o pedido e retorna de qual CD cada item deve ser enviado.

    ```
    # Request
    {
        "order_products_ids":[¨1¨,¨2¨,¨3¨,¨4¨],
        "zone": "S1",
        "state": "SP"

    }

    # Response
    {
        "order_products":[
            {
                "product_id": "1",
                "distribution_center: "CD1"
            },
            {
                "product_id": "2",
                "distribution_center: "CD3"
            },
            {
                "product_id": "3",
                "distribution_center: "CD5"
            },
            {
                "product_id": "4",
                "distribution_center: "CD21"
            },
        ]
    }

    ```

- GET `/orders/{order_id}`by id de pedido criado retornando os produtos e os CDs de onde virão, seria a mesma resposta do POST.
    ```
    # Response
    {
        "order_products":[
            {
                "product_id": "1",
                "distribution_center: "CD1"
            },
            {
                "product_id": "2",
                "distribution_center: "CD3"
            },
            {
                "product_id": "3",
                "distribution_center: "CD5"
            },
            {
                "product_id": "4",
                "distribution_center: "CD21"
            },
        ]
    }

    ```


### ideias iniciais
- Cache da consulta GetCDsByItemId distribuido e com rotatividade pra não virar uma coisa gigante, reduzindo o número de chamadas.
- Usar go routines uma pra cada item do pedido pra fazer a chamada GetCDsByItemId.

### ideias futuras ou opcionais
- Usar gRPC futuramente para outro endpoint que consulta multiplos itens de uma vez só lá de distribution_centers e retorna tudo por streaming pra agilizando ambos os lados.


### Todo 
- [ ] Criar POST `/orders`
- [ ] Criar GET  `/orders/{order_id}`
- [x] Criar wiremock do distribution_centers
- [ ] Criar um client http com go routines para o `/distributioncenters?itemId=123`
- [x] Usar Hexagonal Arch
- [ ] Documentar a Api com Swagger
- [-] Infra do código (docker, makefile etc) 
- [ ] Criar testes unitarios
- [ ] Criar testes de integração (Insomnia e se der tempo Serenity+Cucumber ou Robot)
- [-] Criar o start da aplicação pelo cli usando o Cobra
- [ ] Fazer o setup da suite de qualidade com o Sonarqube
- [ ] Usar os linters do Golang na pipeline de cada projeto


### Todo futuro ou opcional
- [ ] Criar um cache distribuido retroalimentado por produto com lifetime pra não virar um troço gigante.
- [ ] Criar o CRUD de `/orders`