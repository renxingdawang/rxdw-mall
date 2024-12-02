namespace go payment

service PaymentService{
    ChargeResp Charge(1:ChargeReq req)
    // 取消支付（高级）
    CancelPaymentResp CancelPayment(2:CancelPaymentReq req)
    // 定时取消支付（高级）
    TimedCancelPaymentResp TimedCancelPayment(3:TimedCancelPaymentReq req)
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

struct CancelPaymentReq{
    1:required string order_id
    2:required i32 usr_id
}
struct TimedCancelPaymentReq{
    1:required string order_id
    2:required i32 usr_id
}

struct CancelPaymentResp{
    1:required bool success
    2:required string transaction_id
}

struct TimedCancelPaymentResp{
    1:required bool success
    2:required string transaction_id
}