namespace go checkout

include"payment.thrift"

service CheckoutService{
    CheckoutResp Checkout(1:CheckoutReq req)
}

struct Address{
    1: required string street_address
    2: required string city
    3: required string state
    4: required string country
    5: required i32 zip_code
}

struct CheckoutReq{
    1:required i32 user_id
    2:required string firstname
    3:required string lastname
    4:required string email
    5:required Address address
    6:required payment.CreditCardInfo credit_card
}

struct CheckoutResp{
    1:required string order_id
    2:required string transaction
}