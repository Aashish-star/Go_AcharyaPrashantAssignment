# Command to run 
go run main.go

# API1 Sign Up(1st Test case) (User created Successfully)

curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"abcxyz@gmail.com\", \"password\": \"abcxyz1123\", \"name\": \"abcxyz\"}"

#  Sign Up(2st Test case) (User Already present) 

curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"abcxyz@gmail.com\", \"password\": \"abcxyz1123\", \"name\": \"abcxyz\"}"


# API2 Sign In (User created Successfully)

curl --location "http://localhost:8080/user/sign-in" ^
--header "Content-Type: application/json" ^
--data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

For the validation of token accessToken will be received from SIGN IN API. Use that in place of VALID_TOKEN to authorise it
# API 3 Authorise token
 PART 1 Valid Token
 
curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: VALID_TOKEN" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"


PART 2 Expire Token

curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxNTE1ODIsInVzZXJuYW1lIjoiYWJjeHl6In0.FuzINDJGLMVy-v1zFTJByJYoPHiwPJJTLYuSx2c8L7E" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

PART 3 Invalid Token

curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxNTMxMDAsInVzZXJuYW1lIjoiYWJjeHl6In0.5UADSjqwvteZqhDuYqbwe13xjxSZsWJTzLANU2MrDA" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

# Refresh Token (Dependent API on API2)

You will get the REFRESH_TOKEN from SIGN_API, Replace REFRESH_TOKEN with the actual value received from SIGN_API 

curl --location "http://localhost:8080/user/refresh-token" ^ --header "Authorization: REFRESH_TOKEN" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"
