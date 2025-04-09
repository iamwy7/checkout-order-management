## meli challenge
### Olá caros reviewers do meu humilde código, cozinhei bastante coisa por aqui, deixa que eu resumo pra vocês :)
Tendo em vista que é um serviço que cria pedidos e também precisa consulta-los em outro momento, foi assim que fiz os principais pontos do teste:

### Requisitos Funcionais do teste
- [x] Criar POST do recurso `/orders` para criar pedidos com produtos.
- [x] Criar GET do recurso  `/orders/{order_id}` para consultar pedidos já criados.
- [x] Criar um Mock do serviço `distribution_centers` que cuide da request `/distributioncenters?itemId=123`.
- [x] Criar um client dentro do serviço de `orders` que faça uma request para `/distributioncenters?itemId=123` como requisitado.

> Considerações sobre os requisitos
> - Troquei o itemId `123` para `UUID` no mock de `distribution_centers`.
> - Criei um serviço chamado `orders` para atender essas capacidades.
> - Criei um serviço chamado `distribution_centers` para atender as capacidades que o mock também atende.
> - Adicionei um parametro para escolher de qual centro de distribuição um produto virá e sempre será o que tiver mais unidades do produto em questão.

### Como executar tudo?
- Todos os serviços estão no `docker-compose.yaml` na [raiz do projeto](./docker-compose.yaml).
- Cada uma das apps tem sua documentação de api com Swagger no path `/swagger`.
- Deixei uma collection insomnia para teste E2E bem [aqui](./orders/test/insomnia/collection.yaml).
- Tem um arquivo com os SELECTS em SQL equivalentes aos da aplicação `orders` bem [aqui](./orders/test/mysql/order_selects.sql)
- Tem um arquivo com os SELECTS em SQL equivalentes aos da aplicação `distribution_centers` bem [aqui](./distribution_centers/test/mysql/distribution_centers.sql)
- Criei o Mock também do serviço `distribution_centers` usando o [Wiremock](https://github.com/wiremock/wiremock), que também se encontra [aqui](./orders/test/wiremock/docker-compose.yaml) para uma resposta menor caso queiram.
- Nessa [pasta](./orders/test/orders_payload/) teremos payloads para criar os pedidos, com números de produtos variados, é só copiar e collar na collection ou na ferramenta que escolher pra testar as apis do serviço de `orders`.

### Sobre os endpoints
> Essas requests estão mapeadas na collection do insomnia :)

- `/orders` é um POST de pedido com um array de produtos (até 100 itens), que consulta quais os CDs que possuem os produtos, determina o melhor CD para cada um deles, cria no banco de dados o pedido e retorna de qual CD cada item deve ser enviado.

    ```
    # Request
    {
        "products": [
          {
            "id": "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
            "name": "Product Name 1",
            "price": 123.45,
            "quantity": 3
          },
          {
            "id": "2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7q",
            "name": "Product Name 2",
            "price": 456.78,
            "quantity": 5
          }
        ],
        "zone": "S1",
        "state": "SP"
      }

    # Response
    {
        "id": "59183153-e264-4642-ab38-9f459cff8b18", #Sempre um novo a cada request
        "zone": "S1",
        "state": "SP",
        "status": "PENDING",
        "created_at": "2025-04-08 06:54:37",
        "updated_at": "2025-04-08 06:54:37",
        "products": [
            {
                "id": "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
                "name": "Perfume Masculino",
                "price": 175,
                "quantity": 2,
                "distribution_center": {
                    "id": "66af3398-fd04-471f-aefc-a2280e7d02d3",
                    "name": "CD1",
                    "zone": "S1"
                }
            },
            {
                "id": "c60ce040-e2e4-4828-b959-a500996816b8",
                "name": "Camisa Social",
                "price": 200,
                "quantity": 1,
                "distribution_center": {
                    "id": "c7f2d3e4-6a9f-1d8b-0c2a-4d3e5b2a7f6c",
                    "name": "CD25",
                    "zone": "C1"
                }
            }
        ],
        "products_count": 2
    }

    ```

- `/orders/{order_id}` é um GET by id de pedido já criado para retornar tanto os dados dele quanto os produtos e os CDs de onde virão, seria a mesma resposta do `/orders`.
    ```
    # Response
    {
        "id": "59183153-e264-4642-ab38-9f459cff8b18", #Sempre um novo a cada request
        "zone": "S1",
        "state": "SP",
        "status": "PENDING",
        "created_at": "2025-04-08 06:54:37",
        "updated_at": "2025-04-08 06:54:37",
        "products": [
            {
                "id": "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
                "name": "Perfume Masculino",
                "price": 175,
                "quantity": 2,
                "distribution_center": {
                    "id": "66af3398-fd04-471f-aefc-a2280e7d02d3",
                    "name": "CD1",
                    "zone": "S1"
                }
            },
            {
                "id": "c60ce040-e2e4-4828-b959-a500996816b8",
                "name": "Camisa Social",
                "price": 200,
                "quantity": 1,
                "distribution_center": {
                    "id": "c7f2d3e4-6a9f-1d8b-0c2a-4d3e5b2a7f6c",
                    "name": "CD25",
                    "zone": "C1"
                }
            }
        ],
        "products_count": 2
    }

    ```

- `/distributioncenters?itemId=1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p` é um GET by id de produto criado para retornar os CDs em que aquele produto tem estoque. Pode ser executado pelo [distribution_centers](./distribution_centers) ou pelo [wiremock](./orders/test/wiremock), fica a seu critério.
    ```
    # Response
    {
        "itemId": "1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
        "distributionCenters": [
            {
                "id": "a2b3c4d5-6f25-4d2c-8e3a-1a2b3c4d5e6f",
                "name": "CD5",
                "zone": "S2",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 233
            },
            {
                "id": "a5e3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f",
                "name": "CD1",
                "zone": "C1",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 120
            },
            {
                "id": "b2d3e4f5-6a7f-1c9d-8b0c-3d4e5b2a7f6c",
                "name": "CD2",
                "zone": "N1",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 123
            },
            {
                "id": "e4f5a7d2-6c9f-1d8b-0c2a-3d4e5b2a7f6c",
                "name": "CD3",
                "zone": "N2",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 443
            },
            {
                "id": "f1a3b8d4-6f25-4d2c-8e3a-1a2b3c4d5e6f",
                "name": "CD4",
                "zone": "S1",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 3566
            },
            {
                "id": "i0j1k2l3-6f25-4d2c-8e3a-1a2b3c4d5e6f",
                "name": "CD13",
                "zone": "S2",
                "state": "SP",
                "status": "ACTIVE",
                "quantity": 511
            }
        ]
    }

    ```

### Decisões tecnicas pessoais
- [x] Usar a linguagem Go, por proximidade com o ecossistema do Meli e por gosto pela performance dela.
- [x] Usar Hexagonal Architecture.
    - [x] Temos Inversão de Controle.
    - [x] Temos Injeção de Dependencia com as Interfaces entre domínio e adapters.
    - [x] Temos Responsabilidades unicas pois cada jornada/capacidade da aplicação é contida em use cases.
- [x] Usei o [Swaggo/swag](https://github.com/swaggo/swag) para documentar a Api com Swagger, necessário para qualquer Dev/Parceiro que vá "consumir" a API.
- [x] Docker para todo o ambiente, principal ferramenta pra facilitar na esteira de CI/CD e o próprio ambiente do Dev.
- [x] Temos testes unitarios com um coverage decente.
- [x] Temos Collection do insomnia para teste E2E.
- [x] No serviço de `orders`
    - [x] Usa a lib [Cobra](https://github.com/spf13/cobra) pra iniciar e personalizar a porta que a aplicação usará.
    - [x] Usa a lib [Viper](https://github.com/spf13/viper) pra ter acesso a variáveis de ambiente pra conectar no banco de dados e ter a url do serviço de CDs.
- [x] No serviço de `distribution_centers`
    - [x] Implementamos o mock do teste com um código mais simples, já que ele entra no lugar do mock requisitado no teste assim como o wiremock e vice versa, mas também temos Swagger e um endpoint de HealthCheck.
    - [x] Também temos um banco de dados pra ele que guarda o relacionamento de produtos e centros de distribuição, que o mock também retornaria.

- [x] Temos um container de Sonarqube que confere o código e coverage. Se eu configurar certinho as exclusões ele deixa de massacrar meu código ;-;.


### Ideias futuras
- Estou terminando uma implementação em código Go do serviço de distribution_centers, que tem vários produtos espalhados por centros de distribuição no banco de dados dele, vai ser mais fácil pra gerar massas de teste onde o pedido tenha 100 produtos por exemplo. Não subi junto porque não terminei a tempo, mas vou subir em uma branch separada `feature/dc_service` pra não afetar a avaliação.
- Criar/Usar um serviço que valide os produtos, como um serviço de `catalog`. Manter a responsabilidade de produtos dentro do serviço de pedidos não é eficiente, pois o serviço de centros de distribuição também precisa manter essas informações atualizadas, levando a necessidade de mais um serviço na arquitetura.
- Criar dentro de `orders` um worker que consulte se algum pedido foi feito pela metade, seja por algum erro de conexão com o banco no meio da transação, para que ele exclua pedidos, produtos ou centros de distribuição "orfãos". Mantendo a Atomicidade da transação, por mais que ela já aconteça de forma invertida.
- Criei um parametro de State/Estado. Supondo que o negócio abra mais CDs em outros estados, pode ser um parametro de decisão também, além de saber qual CD tem mais daquele produto, dá pra pesquisar baseado em CDs que estejam em estados adjacentes ao informado no pedido. Por isso no Path da consulta tem um query parameter de `zone` e `state`.


### Considerações finais
Foi divertido, espero que tenha cobrido os requisitos e atingido as espectativas :)

## Obrigado!

<!-- 
Essa parte eu vou usar pra guardar comandos úteis em torno do projeto, pode ignorar.

Para recriar as documentações usando o swaggo
``` sh
    swag init --parseDependency --parseInternal --output ./docs
``` 


Para rodar um container que tenha sonar scanner embutido:
docker run --rm\
    -e SONAR_HOST_URL="http://sonarqube:9000"\
    -e SONAR_LOGIN="squ_9b0f2da84dccfd6459e60fe2bdd44cdb5ba8c462"\
    --network orders_orders_network -v ".:/usr/src"\
    sonarsource/sonar-scanner-cli

-->
