package main

import "fmt"

const dividerBar = "================================================"

type Product struct {
	Item     ProductItem
	Quantity int
}

type Store struct {
	Money    int
	Products []*Product
}

// Product 객체의 생성자
func NewProduct(item ProductItem, quantity int) *Product {
	product := new(Product)
	product.Item = item
	product.Quantity = quantity
	return product
}

// Product 상품을 고객에게 서빙
func (p *Product) Serve() error {
	// 상품을 만듦
	err := p.Item.Make()

	// 상품을 만드는 과정에 오류가 있을 때 에러를 부모 함수에 전파
	if err != nil {
		return err
	}

	// 상품을 포장함
	err = p.Item.Package()

	// 상품을 포장하는 과정에 오류가 있을 때 에러를 부모 함수에 전파
	if err != nil {
		return err
	}

	// 상품을 집고 건넴
	err = p.Item.Pick()

	// 상품을 집는 과정에 오류가 있을 때 에러를 부모 함수에 전파
	if err != nil {
		return err
	}

	// 모든 과정이 문제가 없으면 error는 nil로 반환
	return nil
}

// Store 객체의 생성자
func NewStore(money int, products []*Product) *Store {
	store := new(Store)
	store.Money = money
	store.Products = products
	return store
}

// 특정 상품을 원하는 수량 만큼 판다
// productName 이름의 상품을 quantity 만큼 팔고
// 수익을 Store 객체의 Money 멤버 변수에 더하는 함수
func (s *Store) SellProduct(productName string, quantity int) {
	product := s.GetProduct(productName)

	// 상품을 고객에게 서빙함
	err := product.Serve()
	if err != nil {
		fmt.Printf("상품을 판매하는 과정 중 문제가 발생했습니다. %v\n",
			err)
		return
	}

	product.Quantity -= quantity
	s.Money += product.Item.GetPrice() * quantity
}

// 모든 상품 객체 목록을 가져온다
func (s *Store) GetProducts() []*Product {
	return s.Products
}

// 특정 이름의 상품 객체를 가져온다
func (s *Store) GetProduct(productName string) *Product {
	for _, product := range s.Products {
		if product.Item.GetName() == productName {
			return product
		}
	}
	return nil
}

// 특정 이름의 상품이 원하는 수량만큼 있는지 확인한다
// 수량이 있음: true, 수량이 모자름: false
func (s *Store) CheckProductQuantity(
	productName string, quantity int) bool {
	product := s.GetProduct(productName)
	return product.Quantity >= quantity
}

// 현재 상점의 보유 잔금을 가져온다
func (s *Store) GetMoney() int {
	return s.Money
}

// 사용자로부터 구매 상품과 수량을 입력받는 함수
func HandleChoiceProduct(myStore *Store) (exit bool) {
	for true {
		var choice string

		fmt.Println(dividerBar)
		fmt.Print("구매할 상품의 이름을 알려주세요 (나가기 exit): ")
		fmt.Scanln(&choice)

		if choice == "exit" {
			fmt.Println("이용해주셔서 감사합니다!")
			fmt.Println(dividerBar)
			fmt.Printf("최종 가게 보유 자금 %d원\n", myStore.GetMoney())
			return true
		}

		product := myStore.GetProduct(choice)
		if product == nil {
			fmt.Printf("우리 매장에는 %s 라는 이름의 상품은 없네요.\n",
				choice)
			continue
		} else if product.Quantity == 0 {
			fmt.Printf("%s의 재고 수량이 떨어졌네요...\n", choice)
			continue
		}

		var quantity int
		fmt.Print("몇 개를 구매하시겠어요?: ")
		fmt.Scanln(&quantity)

		isExists := myStore.CheckProductQuantity(
			product.Item.GetName(),
			quantity)
		if isExists == false {
			fmt.Printf("%s의 재고 수량이 떨어졌네요...\n", choice)
			continue
		}

		myStore.SellProduct(product.Item.GetName(), quantity)
		fmt.Println("우리 매장을 이용해주셔서 감사합니다!")
		break
	}

	return false
}

func main() {
	// 커피 종류를 정의
	americano := NewCoffee(
		"아메리카노", 3000, "블랜딩 커피", Bitter, Waiting)
	latte := NewCoffee(
		"카페라떼", 3500, "블랜딩 커피", Sweet, Waiting)
	cateMocha := NewCoffee(
		"카페모카", 4000, "디저트 커피", Sweet, Waiting)
	dripCoffe := NewCoffee(
		"드립커피", 7000, "원두 커피", FruitFlavor, Waiting)
	dutchCoffee := NewCoffee(
		"더치커피", 5000, "더치 커피", Bitter, Waiting)

	// 쥬스 종류를 정의
	orangeJuice := NewJuice(
		"오렌지 쥬스", 3000, "과일 쥬스", Sweet, Waiting)
	AloeJuice := NewJuice(
		"알로에 쥬스", 3000, "건강 쥬스", FruitFlavor, Waiting)

	// 차 종류를 정의
	milk := NewTea(
		"우유", 3000, "우유", Sweet, Waiting)
	hotChoco := NewTea(
		"핫초코", 3000, "핫초코", Sweet, Waiting)

	// 판매할 상품 수량을 Product 객체로 생성
	productAmericano := NewProduct(americano, 5)
	productLatte := NewProduct(latte, 2)
	productCateMocha := NewProduct(cateMocha, 3)
	productDripCoffee := NewProduct(dripCoffe, 4)
	productDuchCoffee := NewProduct(dutchCoffee, 6)
	productOrangeJuice := NewProduct(orangeJuice, 8)
	productAloeJuice := NewProduct(AloeJuice, 3)
	productMilk := NewProduct(milk, 3)
	productHotChoco := NewProduct(hotChoco, 5)

	// 상품을 판매할 상점을 Store 객체로 생성
	// 처음 자금은 0원
	// 판매 상품은 아까 정의한 Product 객체를 배열로 제공
	myStore := NewStore(0, []*Product{
		productAmericano,
		productLatte,
		productCateMocha,
		productDripCoffee,
		productDuchCoffee,
		productOrangeJuice,
		productAloeJuice,
		productMilk,
		productHotChoco,
	})
	fmt.Printf("store: %v", myStore)

	for true {
		fmt.Println(dividerBar)
		fmt.Println("우리 매장을 찾아와주셔서 감사합니다!")
		fmt.Printf("현재 매장 보유 자금: %d원\n", myStore.GetMoney())
		fmt.Println(dividerBar)
		fmt.Println("우리 매장 상품 리스트입니다.")
		for i, product := range myStore.GetProducts() {
			fmt.Printf("%d. %s 상품: %d원, (재고: %d개)\n",
				i+1,
				product.Item.GetName(),
				product.Item.GetPrice(),
				product.Quantity)
		}

		isExit := HandleChoiceProduct(myStore)
		if isExit {
			return
		}
	}
}
