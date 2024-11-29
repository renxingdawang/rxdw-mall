namespace go errno

enum Err{
    Success=0,
    NoRoute=1,
    NoMethod=2,
    BadRequest=10000,
    ParamsErr=10001,
    AuthorizeFail=10002,
    TooManyRequest=10003,
    ServiceErr=20000,
    RPCAuthSrvErr=30000,
    AuthSrvErr=30001,
    RPCCartSrvErr=40000,
    CartSrvErr=40001,
    RPCCheckoutSrvErr=50000,
    CheckoutErr=50001,
    RPCOrderErr=60000,
    OrderErr=60001,
    RPCPaymentErr=70000,
    PaymentErr=70001,
    RPCProductErr=80000,
    ProductErr=80001,
    RPCUserErr=90000,
    UserErr=90001,



}