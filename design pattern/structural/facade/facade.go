package facade

import "fmt"

type AccountService struct{}

func (a AccountService) PlaceOrder(email string) {
	fmt.Println("Email of account order: ", email)
}

type PaymentService struct{}

func (p PaymentService) CashMethod() {
	fmt.Println("Pay on cash")
}

func (p PaymentService) CreditMethod() {
	fmt.Println("Pay on Credit Card")
}

func (p PaymentService) BankingMethod() {
	fmt.Println("Pay on Banking ")
}

type ShippingService struct{}

func (s ShippingService) Express() {
	fmt.Println("Express Shipping")
}

func (s ShippingService) Saving() {
	fmt.Println("Saving Shipping")
}

func (s ShippingService) Normal() {
	fmt.Println("Normal Shipping")
}

type ShopFacade struct {
	accService      AccountService
	paymentService  PaymentService
	shippingService ShippingService
}

func newShopFacade() ShopFacade {
	s := ShopFacade{AccountService{}, PaymentService{}, ShippingService{}}
	return s
}

var instance ShopFacade = newShopFacade()

func GetInstance() ShopFacade {
	return instance
}

func (s ShopFacade) PurchaseByCashWithNormalShipping(email string) {
	s.accService.PlaceOrder(email)
	s.paymentService.CashMethod()
	s.shippingService.Normal()
}

func (s ShopFacade) PurchaseByBankingWithExpressShipping(email string) {
	s.accService.PlaceOrder(email)
	s.paymentService.BankingMethod()
	s.shippingService.Express()
}

// func main() {
// 	ShopFacade := facade.GetInstance()
// 	ShopFacade.PurchaseByBankingWithExpressShipping("trung@gmail.com")
// 	ShopFacade.PurchaseByCashWithNormalShipping("anna@gmail.com")
// }
