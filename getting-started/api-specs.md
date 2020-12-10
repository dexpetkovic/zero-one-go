# API specifications

## Create bank account

Bank account is created by posting the below body to the API service:
```
POST /v1/organisation/accounts HTTP/1.1
Content-Type: application/json

{
  "data": {
    "type": "accounts",
    "id": "DBC1D4F4-82E7-4286-B169-1D1B7D0D3989",
    "organisation_id": "934AD475-6612-44A1-A0E8-6D74408780BC",
    "attributes": {
      "country": "NL",
      "base_currency": "EUR",
      "bank_id": "111222",
      "bank_id_code": "ABNA"
    }
  }
}
```

Expected response on successful create is below:

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "data": {
    "type": "accounts",
    "id": "DBC1D4F4-82E7-4286-B169-1D1B7D0D3989",
    "version": 0,
    "organisation_id": "934AD475-6612-44A1-A0E8-6D74408780BC",
    "attributes": {
      "country": "NL",
      "base_currency": "EUR",
      "account_number": "4141111",
      "bank_id": "111222",
      "bank_id_code": "ABNA",
      "iban": "NLABNA4003004141111",
      "status": "confirmed"
    }
  }
}
```

## Fetch bank account

```
GET /v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc HTTP/1.1
Accept: application/json
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "data": {
    "type": "accounts",
    "id": "DBC1D4F4-82E7-4286-B169-1D1B7D0D3989",
    "version": 0,
    "organisation_id": "934AD475-6612-44A1-A0E8-6D74408780BC",
    "attributes": {
      "country": "NL",
      "base_currency": "EUR",
      "account_number": "4141111",
      "bank_id": "111222",
      "bank_id_code": "ABNA",
      "iban": "NLABNA4003004141111",
      "status": "confirmed"
    }
  }
}
```

## List bank accounts

```
GET /v1/organisation/accounts/ HTTP/1.1
Accept: application/json
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "data": [
    {
      "type": "accounts",
      "id": "DBC1D4F4-82E7-4286-B169-1D1B7D0D3989",
      "version": 0,
      "organisation_id": "934AD475-6612-44A1-A0E8-6D74408780BC",
      "attributes": {
        "country": "NL",
        "base_currency": "EUR",
        "account_number": "4141111",
        "bank_id": "111222",
        "bank_id_code": "ABNA",
        "iban": "NLABNA4003004141111",
        "status": "confirmed"
      }
    },
    {
      "type": "accounts",
      "id": "F8A02782-DC49-4744-9CF0-2C75B46324B3",
      "version": 0,
      "organisation_id": "934AD475-6612-44A1-A0E8-6D74408780BC",
      "attributes": {
        "country": "NL",
        "base_currency": "EUR",
        "account_number": "4144444",
        "bank_id": "111222",
        "bank_id_code": "ABNA",
        "iban": "NLABNA4003004144444",
        "status": "confirmed"
      }
    }
  ]
}
```

## Delete bank account

```
DELETE /v1/organisation/accounts/DBC1D4F4-82E7-4286-B169-1D1B7D0D3989 HTTP/1.1
```

```
204 No Content
```