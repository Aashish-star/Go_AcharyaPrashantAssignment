# Command to run 
go run main.go

# API1 Sign Up(1st Test case) (User created Successfully)

curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

#  Sign Up(2st Test case) (User Already present) 

curl --location "http://localhost:8080/user/sign-up" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"


# API2 Sign In (User created Successfully)

curl --location "http://localhost:8080/user/sign-in" ^
--header "Content-Type: application/json" ^
--data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

For the validation of token accessToken will be received from SIGN IN API. Use that in place of VALID_TOKEN to authorise it
# API 3 Authorise token
 PART 1 Valid Token (DEPENDENT ON API2)
 
curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: VALID_TOKEN" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"


PART 2 Expire Token

curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAxODI3NTAsInVzZXJuYW1lIjoiVGVzdCJ9.fRUPr9z4Lilrj1fWuIe5Qbl5xKj6vVu4T0XrsfyDxMs" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

PART 3 Invalid Token

curl --location "http://localhost:8080/user/authorize-token" ^ --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAzOTg5MjQsInVzZXJuYW1lIjoiVGVzdCJ9.F_l7Pb_rC3pTwk89WszOPgVIGusccUdlTIDAhvnGec" ^ --header "Content-Type: application/json" ^  --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"

# API4 Refresh Token (Dependent API on API2)

You will get the REFRESH_TOKEN from SIGN_API, Replace REFRESH_TOKEN with the actual value received from SIGN_API 

curl --location "http://localhost:8080/user/refresh-token" ^ --header "Authorization: REFRESH_TOKEN" ^ --header "Content-Type: application/json" ^ --data-raw "{\"email\": \"Test@gmail.com\", \"password\": \"Test1123\", \"name\": \"Test\"}"
