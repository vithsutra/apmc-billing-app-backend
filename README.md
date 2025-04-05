# Billing Software

# **1. Creating User Account**

```jsx
http://localhost:8000/create/user
```

### HTTP METHOD → POST

Request Body format

```jsx
{
"user_name":"Vithsutra Traders",
"[user_email":"vithsutra.tech@gmail.com](mailto:user_email%22:%22vithsutra.tech@gmail.com)",
"user_password":"password",
"user_phone":"8088281469"
}
```

HTTP Response format

```jsx
{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM3NzU1NjEsInVzZXJfaWQiOiI2NjkxMzczYi1mZDEyLTQzMjEtYWM5ZS00YzZiNjFmYzM5MmIifQ.5hF0-7vLmLKuU8tOsGVOA28RURr6hJzScmTiJgtvEcs"
}
```

# **2. Delete User Account**

```jsx
http://localhost:8000/delete/{user_id}

```

### **HTTP Method → DELETE**

HTTP Response format

```jsx
{
    "message": "User deleted successfully"
}
```

# **3. User Login**

```jsx
http://localhost:8000/login
```

### HTTP METHOD → POST

Request Body format

```jsx
{
    "user_email":"vithsutra.tech@gmail.com",
    "user_password":"password"
}
```

HTTP Response Format

```jsx
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM3NzYzNzQsInVzZXJfaWQiOiI4ZTg5YWY4OC03OTZiLTRhNDYtYTlkYS00ZTEzMGZmMTlmYzUifQ.1Gq4VEqwScQFErKlVCZtAmpItmjurVxR438l1jO4rSg"
}
```

# **4. Create Consignee**

```jsx
http://localhost:8000/create/consignee
```

### HTTP METHOD → POST

Request Body format

```jsx
{
    "consignee_name":"Vithsutra",
    "consignee_address":"AIET,Mijar,Dakshina Kannnada, 574225",
    "consignee_gstin":"19292JSJS",
    "consignee_phone_number":"8088281469",
    "consignee_state":"Karnataka",
    "consignee_state_code":"29",
    "user_id":"8e89af88-796b-4a46-a9da-4e130ff19fc5"
}
```

HTTP Response Format

```jsx
{"message":"consignee created successfully"}
```

# **5. Get All Consignees**

```jsx
http://localhost:8000/get/consignees/<USER_ID>
```

### HTTP METHOD → GET

HTTP Response Format

```jsx
{
  "consignee_details": [
    {
      "consignee_id": "8af9fdfe-381f-4f0d-b76c-98f48e4c9b18",
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

# 6. **Delete Consignee**

```jsx
http://localhost:8000/delete/consignee/<Consignee_ID>
```

### **HTTP Method → DELETE**

HTTP Response format

```jsx
{"message":"Consignee deleted successfully"}
```

# **7. Create Receiver**

```jsx
http://localhost:8000/create/receiver
```

### **HTTP Method → POST**

Request Body format

```jsx
{
    "receiver_name":"Srujan",
    "receiver_address":"Shimoga,vidyanagara",
    "receiver_gstin":"1919SKSK",
    "receiver_state":"Karnataka",
    "receiver_state_code":"29",
    "user_id":"8e89af88-796b-4a46-a9da-4e130ff19fc5"
}
```

HTTP Response Format

```jsx
{
  "message": "created receiver successfully"
}
```

# 8. **Get All Receivers**

```jsx
http://localhost:8000/get/receivers/<USER_ID>
```

### HTTP METHOD → GET

HTTP Response Format

```jsx
{
  "receiver_details": [
    {
      "receiver_id": "32bec541-4b23-45b0-81ce-bcf3878073d1",
      "receiver_name": "Srujan",
      "receiver_address": "Shimoga,vidyanagara",
      "receiver_gstin": "1919SKSK",
      "receiver_state": "Karnataka",
      "receiver_state_code": "29",
      "user_id": ""
    }
  ]
}
```

# 9.**Delete Receiver**

```jsx
http://localhost:8000/delete/receiver/<Reciever_Id>
```

### **HTTP Method → DELETE**

HTTP Response format

```jsx
{
    "message": "receiver deleted successfully"
}
```

# **10. Create Invoice**

```jsx
http://localhost:8000/create/invoice
```

### **HTTP Method → POST**

Request Body format

```jsx
{
  "invoice_name":"Vithsutra Invoice",
  "invoice_reverse_charge":"no",
  "invoice_state":"Karnataka",
  "invoice_state_code":"29",
  "invoice_challan_number":"12",
  "invoice_vehicle_number":"1290 20 202",
  "invoice_date_of_supply":"21/3/2025",
  "invoice_place_of_supply":"AIET Mijar",
  "invoice_gst":"18",
  "user_id":"8e89af88-796b-4a46-a9da-4e130ff19fc5",
  "receiver_id":"721f4a3a-79d5-4517-9199-f5c2f66fb535",
  "consignee_id":"8af9fdfe-381f-4f0d-b76c-98f48e4c9b18"
}
```

HTTP Response Format

```jsx
{
  "invoice_id": "83e0ee54-5f9b-4536-88ae-9f3cb0e55798"
}
```

# 11.Delete Invoice

```jsx
http://localhost:8000/delete/invoice/<Invoice_ID>
```

### **HTTP Method → DELETE**

HTTP Response Format

```jsx
{
    "message": "invoice deleted successfully"
}
```

# 12.Get Invoice

```jsx
http://localhost:8000/get/invoices/<USER_ID>
```

### **HTTP Method → GET**

HTTP Response Format

```jsx
{
  "invoices": [
    {
      "invoice_id": "435a9b43-c721-42a7-8513-1190620492f8",
      "name": "Vithsutra Invoice",
      "payment_status": false
    },
    {
      "invoice_id": "83e0ee54-5f9b-4536-88ae-9f3cb0e55798",
      "name": "Vithsutra Invoice",
      "payment_status": false
    }
  ]
}
```

# 13.**Download Invoice PDF**

```jsx
http://localhost:8000/download/invoice/<INVOICE_ID>
```

### **HTTP Method → GET**

Response will be PDF file

# 14.Create Product

```jsx
http://localhost:8000/create/product
```

### HTTP Method →POST

Request Body format

```jsx
{
    "product_name":"Paddy",
    "product_hsn":"90",
    "product_qty":"3",
    "product_unit":"3",
    "product_rate":"120",
    "invoice_id":"83e0ee54-5f9b-4536-88ae-9f3cb0e55798"
}
```

HTTP Response Format

```jsx
{
    "message": "product created successfully"
}
```

# 15.Delete Product

```jsx
http://localhost:8000/delete/product/<PRODUCT_ID
```

### **HTTP Method → DELETE**

HTTP Response Format

```jsx
{
    "message": "product deleted successfully"
}
```

# 16.**Get Products**

```jsx
http://localhost:8000/get/products/<INVOICE_ID>
```

### **HTTP Method → GET**

HTTP Response Format

```jsx
{
  "receiver_details": [
    {
      "product_id": "ade5286b-c441-4daf-b28a-ed23ebdb3a51",
      "product_name": "Paddy",
      "product_hsn": "90",
      "product_qty": "3",
      "product_unit": "3",
      "product_rate": "120",
      "invoice_id": "",
      "total": "360"
    }
  ]
}
```

# 17.**Update Payment Status To Done**

```jsx
http://localhost:8000/update/invoice/payment/<INVOICE_ID>
```

### **HTTP Method → PATCH**

```jsx
{
    "message": "payment status updated successfully"
}
```

# 18.Create Biller

```jsx
http://localhost:8000/create/biller
```

### HTTP Method →POST

Request Body format

```jsx
{
  "biller_name": "Shubhang",
  "biller_address": "Vidyanagara",
  "biller_mobile": "9876543210",
  "biller_gstin": "29ABCDE1234F2Z5",
  "biller_pan": "ABCDE1234F",
  "biller_mail": "srujan@gmail.com",
  "user_id": "8e89af88-796b-4a46-a9da-4e130ff19fc5"
}

```

HTTP Response Format

```jsx
{
  "message": "Biller created successfully"
}
```

# 19.Get Biller

```jsx
http://localhost:8000/get/billers/<USER_ID>
```

### HTTP  Method →GET

HTTP Response Format

```jsx
[
  {
    "user_id": "8e89af88-796b-4a46-a9da-4e130ff19fc5",
    "biller_id": "74972cdf-e7e4-42f2-83a2-f936cdde3dfa",
    "biller_name": "Shubhang",
    "biller_address": "Vidyanagara",
    "biller_mobile": "9876543210",
    "biller_gstin": "29ABCDE1234F2Z5",
    "biller_pan": "ABCDE1234F",
    "biller_mail": "srujan@gmail.com",
    "biller_companylogo": "PENDING"
  },
  {
    "user_id": "8e89af88-796b-4a46-a9da-4e130ff19fc5",
    "biller_id": "2553442b-d5ca-4710-a6e3-ae9af8a5bdfb",
    "biller_name": "Shubhang",
    "biller_address": "Vidyanagara",
    "biller_mobile": "9876543210",
    "biller_gstin": "29ABCDE1234F2Z5",
    "biller_pan": "ABCDE1234F",
    "biller_mail": "srujan@gmail.com",
    "biller_companylogo": "PENDING"
  }
]
```

# 20.Delete Biller

```jsx
http://localhost:8000/delete/biller/<BILLER_ID>
```

### **HTTP Method → DELETE**

HTTP Response Format

```jsx
{
  "message": "Biller deleted successfully"
}
```

# 21.Create Banker

```jsx
http://localhost:8000/create/banker
```

### HTTP Method →POST

Request Body format

```jsx
{
  "bank_name": "HDFC Bank",
  "bank_branch": "MG Road",
  "bank_account_number": "9876543210",
  "bank_ifsc_code": "HDFC0001234",
  "bank_holder_name": "srujan",
  "user_id": "8e89af88-796b-4a46-a9da-4e130ff19fc5"
}

```

HTTP Response Format

```jsx
{
  "message": "Banker created successfully"
}
```

# 22.Get Banker

```jsx
http://localhost:8000/get/bankers/<USER_ID>
```

### HTTP Method →GET

HTTP Response Format

```jsx
{
  "bankers": [
    {
      "user_id": "8e89af88-796b-4a46-a9da-4e130ff19fc5",
      "bank_id": "7d71facd-3f1a-4763-ad8e-499e51965f56",
      "bank_name": "HDFC Bank",
      "bank_branch": "MG Road",
      "bank_account_number": "9876543210",
      "bank_ifsc_code": "HDFC0001234",
      "bank_holder_name": "srujan"
    }
  ]
}
```

# 23.Delete Banker

```jsx
http://localhost:8000/delete/banker/<BANKER_ID>
```

### HTTP Method →Delete

HTTP Response Format

```jsx
{
  "message": "Banker deleted successfully"
}
```

# 24.Upload Company Logo

```jsx
http://localhost:8000/upload/company/logo/<USER_ID>
```

### HTTP Method →POST

# 25.Delete company logo

```jsx
http://localhost:8000/delete/company/logo/<fileName>
```

### HTTP Method →DELETE

# 26.Forgot Password

```jsx
http://localhost:8000/forgotpassword
```

### HTTP Method →POST

Request Body format

```jsx
{
  "user_email":"vithsutra.tech@gmail.com"
}

```

HTTP Response Format

```jsx
{
  "message": "OTP sent successfully"
}
```

# 27.Validate OTP

```jsx
http://localhost:8000/validateotp
```

### HTTP Method →POST

Request Body format

```jsx
{
  "user_email":"vithsutra.tech@gmail.com",
  "otp":"080060"
}
```

HTTP Response Format

```jsx
{
		"token_id":"b4fa2e19-a1d4-497c-a3cf-f412db957ab9"
}
```

# 28.Reset Password

```jsx
http://localhost:8000/resetpassword
```

### HTTP Method →POST

Request Body format

```jsx
{
	"token_id":"b4fa2e19-a1d4-497c-a3cf-f412db957ab9",
	"new_password":"Sruj@2003",
	"confirm_password":"Sruj@2003"
}
```

HTTP Response Format

```jsx
 {
		 "Password reset successful"
 }
```
