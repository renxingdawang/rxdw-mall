namespace go order

include"cart.thrift"

service OrderService{
    PlaceOrderResp PlaceOrder(1:PlaceOrderReq req)
    ListOrderResp ListOrder(2:ListOrderReq req)
    MarkOrderPaidResp MarkOrderPaid(3:MarkOrderPaidReq req)
}

struct Address{
    1: required string street_address
    2: required string city
    3: required string state
    4: required string country
    5: required i32 zip_code
}
struct OrderItem{
    1: required cart.CartItem item
    2: required i32 cost
}
struct PlaceOrderReq{
    1:required i32 user_id
    2:required string user_currency
    3:required Address address
    4:required string email
    5:list<OrderItem> order_items
}
struct OrderResult{
    1:required string order_id
}

struct PlaceOrderResp{
    1:required OrderResult order
}

struct ListOrderReq{
    1:required i32 user_id
}

struct Order{
    1:list<OrderItem>order_items
    2:required string order_id
    3:required i32 user_id
    4:required string user_currency
    5:required Address address
    6:required string email
    7:required i32 created_at
}
struct ListOrderResp{
    1:list<Order> orders
}

struct MarkOrderPaidReq{
    1:required i32 user_id
    2:required string order_id
}

struct MarkOrderPaidResp{}