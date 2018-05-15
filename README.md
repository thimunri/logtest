# LogParser

A linguagem escolhida para o desenvolvimento foi o golang, uma linguagem compilada, que proporciona bastante produtividade para criação de microserviços.

Para simular os 4 servidores, foi criada 4 rotas no serviço, sendo /server1, /server2/, /server3 e /server4. Cada um dos endpoints cria um access log em diretórios separados, que levam o mesmo nome da rota.

```
root@39384512f9b4:/go/src/github.com/thimunri/logtest# ls -l var/log/
total 16
dr-xr-xr-x 2 root root 4096 May 15 01:04 server1
dr-xr-xr-x 2 root root 4096 May 15 01:45 server2
dr-xr-xr-x 2 root root 4096 May 15 01:45 server3
dr-xr-xr-x 2 root root 4096 May 15 01:45 server4
```

### Dependências
- Docker
- Docker-compose

Com as dependências instaladas, para iniciar o serviço execute:
```docker-compose up -d ```

O docker build se encarregará de instalar todas as dependencias e o sistema estará pronto para uso em:

http://localhost:88/server1
http://localhost:88/server2
http://localhost:88/server3
http://localhost:88/server4

Cada request nos endpoints criará entradas em /var/logs/serverx/access.log, para simular os usuários, o sistema irá gerar dinamicamente 50 uuids. Para gerar um volume alto de logs, uma sugestão é utilizar o ab **( ApacheBench)**

``` ab -n 1000 -c 5 http://localhost:88/server1```

O comando acima executará 1000 requests e 5 de concorrência

## LogParser
Para executar o comando que irá separar os arquivos de log por userId, siga os passos abaixo:

Entre no container:
```docker exec -it logtest bash```

Execute o comando:
```go run main.go -parser```

Esse comando irá varrer o diretório de logs da aplicação, e para cada usuário irá criar um arquivo separado e organizar os logs por dia.

Exemplo de saída:
```
root@39384512f9b4:/go/src/github.com/thimunri/logtest# ls -l var/log/15May2018/
total 948
--w---x--- 1 root root 18375 May 15 01:14 030211ac-3bde-4af3-a07f-29aad7a1a0ed.log
--w---x--- 1 root root 21000 May 15 01:14 06ee94f0-04d6-4d2f-a164-7fbdcb048484.log
--w---x--- 1 root root 14000 May 15 01:14 0ba8b544-dc32-496e-86fe-ce6a935f30d9.log
--w---x--- 1 root root 20125 May 15 01:14 0bbf7572-e4c9-4b4e-adc2-2bd53a2164c9.log
--w---x--- 1 root root 14000 May 15 01:14 0c1336d6-92b4-48f2-821e-c3798f1b8991.log
--w---x--- 1 root root 18375 May 15 01:14 11d3e54d-116f-44e9-a35d-d7784b94ca70.log
--w---x--- 1 root root 17500 May 15 01:14 1ba5025b-de1e-4af9-91e2-d6099af704fb.log
```

### Nota:
Uma versão online do sistema está disponível em:
http://ec2-18-206-234-4.compute-1.amazonaws.com/server1
