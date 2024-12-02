# eg. make service=video generate
generate:
	kitex -module github.com/rxdw-mall ../idl/rpc/$(service).thrift

auth:
	make service=auth generate
cart:
	make service=cart generate
checkout:
	make service=checkout generate
order:
	make service=order generate
payment:
	make service=payment generate
product:
	make service=product generate
user:
	make service=user generate
errno:
	kitex -module github.com/rxdw-mall ../idl/base/errno.thrift