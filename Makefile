deps:
	go mod download ;\
	go mod tidy

build:
	cd cmd ;\
	go build -o hashout

run:
	cd cmd ;\
	go run main.go

docs:
	${GOPATH}/bin/swag init --parseDependency --dir handlers/server/ -g api.go -o handlers/server/

tests:
	go test ./usecase/checkout/... ./usecase/discounts/... ./clients/... ./storage/... ./models/checkout/... ./handlers/dto/... -v -coverprofile=test/coverage/coverage.out ;\
	go tool cover -func=test/coverage/coverage.out ;\
	go tool cover -html=test/coverage/coverage.out -o test/coverage/coverage.html

mocks:
	${GOPATH}/bin/mockery --name=DiscountClient --recursive --output=test/mocks/pb --filename=pb_grpc_discount.go --structname=DiscountClient ;\
	${GOPATH}/bin/mockery --name=Database --recursive --output=test/mocks/db --filename=in_memory.go --structname=ProductDataset ;\
	${GOPATH}/bin/mockery --name=Config --recursive --output=test/mocks/config --filename=app.go --structname=HashoutApp ;\
	${GOPATH}/bin/mockery --name=Service --recursive --output=test/mocks/usecase --filename=discounts_usecase.go --structname=DiscountUseCase ;\
	${GOPATH}/bin/mockery --name=Repository --recursive --output=test/mocks/storage --filename=products_repository.go --structname=ProductRepository ;\
	${GOPATH}/bin/mockery --name=Client --recursive --output=test/mocks/clients --filename=discount_client.go --structname=DiscountClient

proto:
	protoc --go-grpc_out=.  pb/discount.proto ;\
	protoc --go_out=.  pb/discount.proto

di:
	cd cmd/app ;\
	${GOPATH}/bin/wire