namespace go product

service ProductCatalogService{
    ListProductsResp ListProducts(1: ListProductsReq req)
    GetProductResp GetProduct(2: GetProductReq req)
    SearchProductsResp SearchProducts(3:SearchProductsReq req)
}

struct ListProductsReq{
    1:required i32 page
    2:required i64 pageSize
    3:required string categoryName
}

struct Product {
    1:required i32 id
    2:required string name
    3:required string description
    4:required string picture
    5:required double price
    6:list<string> categories
}

struct ListProductsResp {
    1:list<Product> products
}

struct GetProductReq {
    1:required i32 id
}

struct GetProductResp {
    1:required Product product
}

struct SearchProductsReq {
    1:required string query
}

struct SearchProductsResp {
  1:list<Product> results
}

