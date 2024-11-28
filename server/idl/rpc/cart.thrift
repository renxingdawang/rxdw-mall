namespace go cart

service CartService{
    AddItemResp AddItem(1:AddItemReq req)
    GetCartResp GetCart(2:GetCartReq req)
    EmptyCartResp EmptyCart(3:EmptyCartReq req)
}



struct CartItem{
    1:required i32 product_id
    2:required i32 quantity
}

struct AddItemReq{
    1:required i32 user_id
    2:required CartItem item
}

struct AddItemResp{}

struct EmptyCartReq{
    1:required i32 user_id
}

struct EmptyCartResp{}

struct GetCartReq{
    1:required i32 user_id
}

struct Cart{
    1:required i32 user_id
    2:list<CartItem> items
}

struct GetCartResp{
    1:required Cart cart
}