namespace go user

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(2:LoginReq req)
}

struct RegisterReq {
    1: required string email
    2: required string password
    3: required string confirm_password
}

struct RegisterResp {
    1:required i32 user_id
}

struct LoginReq {
    1:required string email
    2:required string password
}

struct LoginResp {
    1:required i32 user_id
}

