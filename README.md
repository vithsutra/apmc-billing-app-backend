# APMC BILLING APP HTTP SERVER API

## 1. Creating User Account

```bash
https://apmc.api.vsensetech.in/create/user
```

### HTTP Method → POST

Request Body format

```json
{
    "user_name":"Vithsutra Traders",
    "user_email":"vithsutra.tech@gmail.com",
    "user_password":"password",
    "user_address":"AIET,Mijar,Dakshina Kannnada, 574225",
    "user_phone":"8088281469",
    "user_gstin":"1823SJSJ12",
    "user_pan":"12322KSKS23"
}
```

HTTP Response format

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzgyMTg1NzAsInVzZXJfaWQiOiI4NzE4NGE0MS1jZGFiLTQxZWUtOWE0Ny1hNGM4OGQxNTg5MDUifQ.uYi55jEfpgdCGfRRkHyYpQrKv2-GdZNra-VCR1hi8VU"
}
```

## 2. Delete User Account

```bash
https://apmc.api.vsensetech.in/delete/<USER_ID>
```

### HTTP Method → DELETE

HTTP Response format

```json
{
    "message": "User deleted successfully"
}
```

## 2. User Login

```bash
https://apmc.api.vsensetech.in/login
```

### HTTP Method → POST

Request Body format

```json
{
    "user_email":"vithsutra.tech@gmail.com",
    "user_password":"password"
}
```

HTTP Response Format

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzgyMTkzMDEsInVzZXJfaWQiOiIwNjg0YWYxNy1iZTZiLTQ5YzktYmFlZi1mMjY4ODI0YjQ4NjYifQ.yWh02PRm8571Jd0s0_2h1MvyXI53cR6gcRpeZ3Opymg"
}
```

## 3. Create Consignee

```bash
https://apmc.api.vsensetech.in/create/consignee
```

### HTTP Method → POST

Request Body format

```json
{
    "consignee_name":"Vithsutra",
    "consignee_address":"AIET,Mijar,Dakshina Kannnada, 574225",
    "consignee_gstin":"19292JSJS",
    "consignee_phone_number":"8088281469",
    "consignee_state":"Karnataka",
    "consignee_state_code":"29",
    "user_id":"0684af17-be6b-49c9-baef-f268824b4866"
}
```

HTTP Response format

```json
{"message":"consignee created successfully"}
```

## 4. Get All Consignees

```bash
https://apmc.api.vsensetech.in/get/consignees/<USER_ID>
```

### HTTP Method → GET

HTTP Response format

```json
{
    "consignee_details": [
        {
            "consignee_id": "7d0e7ced-ed20-4161-b6c9-3f4da5dec137",
            "consignee_name": "Vithsutra",
            "consignee_address": "AIET,Mijar,Dakshina Kannnada, 574225",
            "consignee_gstin": "19292JSJS",
            "consignee_phone_number": "8088281469",
            "consignee_state": "Karnataka",
            "consignee_state_code": "29",
            "user_id": ""
        }
    ]
}
```

## 5. Delete Consignee

```bash
https://apmc.api.vsensetech.in/delete/consignee/<Consignee_ID>
```

### HTTP Method → DELETE

HTTP Response format

```json
{"message":"Consignee deleted successfully"}
```

## 6. Create Receiver

```bash
https://apmc.api.vsensetech.in/create/receiver
```

### HTTP Method → POST

Request Body format

```json
{
    "receiver_name":"Nagabhushan",
    "receiver_address":"Shimoga,Managala",
    "receiver_gstin":"1919SKSK",
    "receiver_state":"Karnataka",
    "receiver_state_code":"29",
    "user_id":"0684af17-be6b-49c9-baef-f268824b4866"
}
```

HTTP Response Format

```json
{
    "message": "created receiver successfully"
}
```

## 7. Get All Receivers

```bash
https://apmc.api.vsensetech.in/get/receivers/<USER_ID>
```

### HTTP Method → GET

HTTP Response Format

```json
{
    "receiver_details": [
        {
            "receiver_id": "b663b492-dc43-495c-b3a6-f7ff6290d079",
            "receiver_name": "Nagabhushan",
            "receiver_address": "Shimoga,Managala",
            "receiver_gstin": "1919SKSK",
            "receiver_state": "Karnataka",
            "receiver_state_code": "29",
            "user_id": ""
        }
    ]
}

```

## 7. Delete Receiver

```bash
https://apmc.api.vsensetech.in/delete/receiver/<Reciever_Id>
```

### HTTP Method → DELETE

HTTP Response format

```json
{
    "message": "receiver deleted successfully"
}
```

## 8. Create Invoice

```bash
https://apmc.api.vsensetech.in/create/invoice
```

### HTTP Method → POST

Request Body format

```json
{
  "invoice_name":"Vithsutra Invoice",
  "invoice_reverse_charge":"no",
  "invoice_state":"Karnataka",
  "invoice_state_code":"29",
  "invoice_challan_number":"12",
  "invoice_vehicle_number":"1290 20 202",
  "invoice_date_of_supply":"21/3/2025",
  "invoice_place_of_supply":"AIET Mijar",
  "user_id":"1bd05614-861f-4438-842a-647bc5c96ef2",
  "receiver_id":"7000db5a-5b3b-4ee3-82bb-809cc9835a74",
  "consignee_id":"00338768-4a00-408e-96e9-fed17c30923d"
}
```

HTTP Response Format

```json
{
    "invoice_id": "749ca101-91e6-4cf5-843c-29cdc072c8a5"
}
```

## 9. Create Product

```bash
https://apmc.api.vsensetech.in/create/product
```

### HTTP Method → POST

Request Body format

```json
{
    "product_name":"Paddy",
    "product_hsn":"90",
    "product_qty":"3",
    "product_unit":"3",
    "product_rate":"120",
    "invoice_id":"749ca101-91e6-4cf5-843c-29cdc072c8a5"
}
```

HTTP Response Format

```json
{
    "message": "product created successfully"
}
```

## 10. Delete Product

```bash
https://apmc.api.vsensetech.in/delete/product/<PRODUCT_ID>
```

### HTTP Method → DELETE

HTTP Response Format

```json
{
    "message": "product deleted successfully"
}
```

## 11. Get All Invoices

```bash
https://apmc.api.vsensetech.in/get/invoices/<USER_ID>
```

### HTTP Method → GET

HTTP Response Format

```json
{
    "invoices": [
        {
            "invoice_id": "749ca101-91e6-4cf5-843c-29cdc072c8a5",
            "name": "Srujan Bhai Invoice",
            "payment_status": false
        }
    ]
}
```

## 11. Download Invoice PDF

```bash
http://apmc.api.vsensetech.in/download/invoice/<INVOICE_ID>
```

### HTTP Method → GET

Response will be PDF file

## 12. Delete Invoice

```bash
https://apmc.api.vsensetech.in/delete/invoice/<Invoice_ID>
```

### HTTP Method → DELETE

HTTP Response Format

```json
{
    "message": "invoice deleted successfully"
}
```

## 13. Get Products

```bash
https://apmc.api.vsensetech.in/get/products/<INVOICE_ID>
```

### HTTP Method → GET

HTTP Response Format

```json
{
    "receiver_details": [
        {
            "ProductId": "4f69262c-7228-41f1-8287-bec6d201907f",
            "ProductName": "Maize",
            "ProductHsn": "90",
            "ProductQty": "3",
            "ProductUnit": "3",
            "ProductRate": "120",
            "InvoiceId": "",
            "Total": "360"
        },
        {
            "ProductId": "e1d182be-2829-4115-9eb8-c99e413c15c0",
            "ProductName": "Paddy",
            "ProductHsn": "9",
            "ProductQty": "3",
            "ProductUnit": "99",
            "ProductRate": "2000",
            "InvoiceId": "",
            "Total": "6000"
        }
    ]
}
```

## 14. Update Payment Status To Done

```bash
https://apmc.api.vsensetech.in/update/invoice/payment/<INVOICE_ID>
```

### HTTP Method → PATCH

```json
{
    "message": "payment status updated successfully"
}
```
