# start the environment of rxdw-mall
.PHONY: start
start:
	docker-compose up -d

# stop the environment of FreeCar
.PHONY: stop
stop:
	docker-compose down

# run the auth
.PHONY: auth
auth:
	go run ./server/cmd/auth

# run the cart
.PHONY: cart
cart:
	go run ./server/cmd/cart

# run the checkout
.PHONY: checkout
checkout:
	go run ./server/cmd/checkout

# run the order
.PHONY: order
order:
	go run ./server/cmd/order

# run the payment
.PHONY: payment
payment:
	go run ./server/cmd/payment

# run the product
.PHONY: product
product:
	go run ./server/cmd/product


# run the user
.PHONY: user
user:
	go run ./server/cmd/user

# run the api
.PHONY: api
api:
	go run ./server/cmd/api