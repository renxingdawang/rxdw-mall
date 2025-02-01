namespace go saga
struct PaymentCancelledEvent {
    1:string order_id
    2:i32 user_id
    3:string transaction_id
}

struct OrderCancelFailedEvent {
    1:string order_id
    2:i32 user_id
    3:string error_reason
}