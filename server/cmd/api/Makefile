update:
	hz update -idl ../../idl/http/$(service).thrift
auth:
	make service=auth update
cart:
	make service=cart update
checkout:
	make service=checkout update
order:
	make service=order update
payment:
	make service=payment update
product:
	make service=product update
user:
	make service=user update
all:
	make auth
	make cart
	make checkout
	make order
	make payment
	make product
	make user