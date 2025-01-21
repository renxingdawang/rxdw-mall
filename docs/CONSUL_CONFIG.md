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
