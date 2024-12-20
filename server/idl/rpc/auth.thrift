namespace go auth

service AuthService{
    DeliveryResp DeliverTokenByRPC(1:DeliverTokenReq req)
    VerifyResp VerifyTokenByRpc(2:VerifyTokenReq req)
    RenewTokenResp RenewTokenByRpc(3:RenewTokenReq req)
}

struct DeliverTokenReq{
    1:required i32 user_id
}

struct VerifyTokenReq{
    1:required string token
}

struct DeliveryResp{
    1:required string token
}

struct VerifyResp{
    1:required bool res
}

struct RenewTokenReq{
    1:required string token
}

struct RenewTokenResp{
    1:required bool res
}