Sign Up(1st Test case) (User created Successfully)
curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"abcxyz@gmail.com\", \"password\": \"abcxyz1123\", \"name\": \"abcxyz\"}"

Sign Up(2st Test case) (User Already present) 
curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"abcxyz@gmail.com\", \"password\": \"abcxyz1123\", \"name\": \"abcxyz\"}"


Sign In (User created Successfully)
curl --location "http://localhost:8080/user/sign-in" ^
--header "Content-Type: application/json" ^
--data-raw "{\"email\": \"abcxyz@gmail.com\", \"password\": \"abcxyz1123\", \"name\": \"abcxyz\"}"

For the validation of token accessToken will be received from SIGN IN API. Use that in place of VALID_TOKEN to authorise it

Valid Token
curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: VALID_TOKEN" ^ --header "Content-Type: application/json" ^ --data-raw "{"email": "abcxyz@gmail.com", "password": "abcxyz1123", "userName": "abcxyz"}"


Expire Token
curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxNTE1ODIsInVzZXJuYW1lIjoiYWJjeHl6In0.FuzINDJGLMVy-v1zFTJByJYoPHiwPJJTLYuSx2c8L7E" ^ --header "Content-Type: application/json" ^ --data-raw "{"email": "abcxyz@gmail.com", "password": "abcxyz1123", "userName": "abcxyz"}"

Invalid Token
curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxNTMxMDAsInVzZXJuYW1lIjoiYWJjeHl6In0.5UADSjqwvteZqhDuYqbwe13xjxSZsWJTzLANU2MrDA" ^ --header "Content-Type: application/json" ^ --data-raw "{"email": "abcxyz@gmail.com", "password": "abcxyz1123", "userName": "abcxyz"}"
