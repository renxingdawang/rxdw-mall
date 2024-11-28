namespace go payment

service PaymentService{
    ChargeResp Charge(1:ChargeReq req)
}

struct CreditCardInfo{
    1:required string credit_card_number
    2:required i32 credit_card_cvv
    3:required i32 credit_card_expiration_year
    4:required i32 credit_card_expiration_month
}

struct ChargeReq{
    1:required double amount
    2:required CreditCardInfo credit_card
    3:required string order_id
    4:required i32 user_id
}

struct ChargeResp{
    1:required string transaction_id
}