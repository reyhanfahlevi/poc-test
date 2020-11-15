package _dev

//go:generate mockgen -source=../svc/account/internal/usecase/usecase.go -destination=../svc/account/internal/usecase/mock_test.go -package=usecase
//go:generate mockgen -source=../svc/shop/internal/usecase/usecase.go -destination=../svc/shop/internal/usecase/mock_test.go -package=usecase
//go:generate mockgen -source=../svc/product/internal/usecase/usecase.go -destination=../svc/product/internal/usecase/mock_test.go -package=usecase
