# Instrução

> Criação de network

```bash
docker network create webproxy
```

> Endpoints
```bash

curl -i -H "Accept: application/json" -H "Content-Type: application/json" http://hostname/resource


curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"account_id":1, "operation_type_id":1, "amount":123.45, "balance":123.45}' \
  http://localhost/v1/transactions
  
curl --header "Content-Type: application/json" \
  --request PATCH \
  --data '{"credit":-60}' \
  http://localhost/v1/accounts/1



```

