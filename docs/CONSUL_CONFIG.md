*rxdwmall/auth_srv*

```apache
{
  "name": "auth_srv",
  "paseto": {
    "secret_key": "b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2",
    "implicit": "paseto-implicit"
  },
  "redis":{
    "host":"127.0.0.1",
    "port":6379
  },
  "otel":{
  	"endpoint":"4317"
  }
}
```

*rxdwmall/user_srv*

```apache
{
  "name": "user_srv",
  "mysql":{
    "host":"127.0.0.1",
    "port":3306,
    "user":"root",
    "password":"123456",
    "db":"rxdwMall",
    "salt":"abcdef"
  },
  "auth_srv":{
  	"name":"auth_srv"
	},
  "otel":{
  	"endpoint":":4317"
  }
}
```

*rxdwmall/product_srv*

```apache
{
  "name": "product_srv",
  "mysql":{
    "host":"127.0.0.1",
    "port":3306,
    "user":"root",
    "password":"123456",
    "db":"rxdwMall",
    "salt":"abcdef"
  },
  "otel":{
  	"endpoint":":4317"
  },
  "redis":{
    "host":"127.0.0.1",
    "port":6379,
    "prefix":"rxdw_mall"
  }
}
```

*rxdwmall/payment_srv*

```apache
{
  "name": "product_srv",
  "mysql":{
    "host":"127.0.0.1",
    "port":3306,
    "user":"root",
    "password":"123456",
    "db":"rxdwMall",
    "salt":"abcdef"
  },
  "otel":{
  	"endpoint":":4317"
  }
}
```
