package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// loopFn()
	// mapsFn()
	// sliceFn()
	// PanicAndDeferAndRecoverFn()
	// PointersFn()
	// structAndPointersFn()
	// MethodFn()
	// PersonFn()
	// fmt.Println("Go redis first program")
	// rdClient := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
	// ctx := context.Background()

	// fmt.Println(rdClient.Ping(ctx).Result())
	// PersonRedisExample(rdClient, ctx)
	ConvertTimeZoneFn()
}

func loopFn() {
	var (
		principleAmtVh int // if not initialized, go does by default
		interestVh     = 10.0
		monthsVh       = 12
	)
	fmt.Println(principleAmtVh, interestVh, monthsVh)
	principleAmtVh = 100

	for count := 0; count < monthsVh; count++ {
		principleAmtVh = principleAmtVh + int(float64(principleAmtVh)*(interestVh/100)) // typematch is essential, all must be either int/float
		fmt.Printf("Month: %d principle amount: %d \n", count, principleAmtVh)
	}
	fmt.Println(interestVh / 100)

	fmt.Println(int(interestVh / 100))
	fmt.Println("principle amount: ", principleAmtVh)

	interestArrVhs := [5]int{10, 9, 10, 9, 9}
	principleAmtVh = 200

	for interestArrVh := range interestArrVhs {
		principleAmtVh += principleAmtVh * interestArrVh / 100
		fmt.Printf("Month: %d principle amount: %d \n", interestArrVh, principleAmtVh)
	}
}

func mapsFn() {
	FromDict := map[string]string{ // integer keys for more efficiency
		"KA": "Karnataka",
		"AP": "Andhra Pradesh",
		"TN": "Tamil Nadu",
		"KL": "Kerala",
	}

	fmt.Println("fromdict: ", FromDict)
	Todict := map[string]string{}

	for key, value := range FromDict {
		Todict[key] = value
		println(key, "--->", Todict[key])
	}

	delete(Todict, "TN")
	fmt.Println(Todict)
	Todict["HP"] = "Himachal Pradesh" // adding a new key
	fmt.Println("todict: ", Todict)   // gives the address
}

func sliceFn() {
	PrimesArrVh := [6]int{2, 5, 6, 3, 1, 5}
	var MySliceSlcVh = PrimesArrVh[1:4] // 1 included, 4 not included
	fmt.Println("My slice: ", MySliceSlcVh)
	MyNumbersIntVh := make([]int, 6, 12)
	fmt.Println("Size: ", len(MyNumbersIntVh), " Capacity:", cap(MyNumbersIntVh), "Values: ", MyNumbersIntVh)

	MyNumbersIntVh = []int{3, 1, 3, 4, 6, 7}
	MyNumbersIntVh = append(MyNumbersIntVh, 21)
	fmt.Println("Appended MyNumbersIntVh: ", MyNumbersIntVh)
}

func DivideFn(Num1Arg int, Num2Arg int) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}
	}()

	if Num2Arg == 0 {
		panic(errors.New("zero denominator not allowed")) // error string must not be punctotized or end with exclamation marks
	}

	vhDivideVhResult := Num1Arg / Num2Arg
	fmt.Println("Division result: ", vhDivideVhResult)
}

func PanicAndDeferAndRecoverFn() {
	DivideFn(10, 4)
	DivideFn(100, 10)
	DivideFn(12, 8)
	DivideFn(10, 0)
	DivideFn(10891, 89)
}

func PointersFn() {
	var amountVh = 25
	print("Amount before: ", amountVh, "\n")
	var amountPtr *int = &amountVh // represents address of variable, * mentions the data type of pointer
	*amountPtr = 50                // in assignment, * access content/value
	print("Amount ptr after: ", *amountPtr, " AmountVar: ", amountVh, "\n")
	amountVh = 100
	println("Amount after updating: ", *amountPtr, " AmountVar: ", amountVh)
	var cityNameVh = "Bengaluru"
	var cityNamePtr *string = &cityNameVh
	println("City name after updating: ", cityNameVh)
	*cityNamePtr = "Mysuru"
	println("City name after updating: ", *cityNamePtr, " cityNameVh: ", cityNameVh)
	cityNameVh = "Hubballi"
	println("City name after updating: ", *cityNamePtr, " cityNameVh: ", cityNameVh)
}

func structAndPointersFn() {
	type StockStrc struct {
		id      int
		ShortId string
	}

	type StockPriceStrct struct {
		stock     StockStrc
		stockDate time.Time
		price     int
	}
	const (
		DDMMYYYY  = "02-01-2006"
		shortform = "02-01-2006"
		yearform  = "2006-01-02"
	)
	var (
		StockPriceStrctVh StockPriceStrct
	)

	StockPriceStrctVh.stock.id = 1
	StockPriceStrctVh.stock.ShortId = "APPL"
	StockPriceStrctVh.stockDate, _ = time.Parse(shortform, "22-11-2022") // use _ to discard any variable
	// Date value sample: Mon Jan 2 15:04:05 MST 2006
	StockPriceStrctVh.price = 1650
	fmt.Println("Stock struct: ", StockPriceStrctVh, "date:", StockPriceStrctVh.stockDate.Format(yearform))

	var StockPriceStrctArrVh [2]StockPriceStrct
	StockPriceStrctArrVh[0].price = 1000
	StockPriceStrctArrVh[0].stock.id = 1
	StockPriceStrctArrVh[0].stock.ShortId = "TCS"
	StockPriceStrctArrVh[0].stockDate, _ = time.Parse(DDMMYYYY, "26-11-2022")
	fmt.Println("StockPriceStrctArrVh: ", StockPriceStrctArrVh[0])

	var StockPrice1StrctArrVh StockPriceStrct
	StockPrice1StrctArrVh.stock.id = 2
	StockPrice1StrctArrVh.stock.ShortId = "DCRD"
	StockPrice1StrctArrVh.price = 1000
	StockPrice1StrctArrVh.stockDate, _ = time.Parse(shortform, "29-12-2022")
	fmt.Println("Stock struct: ", &StockPrice1StrctArrVh, "date:", StockPrice1StrctArrVh.stockDate.Format(yearform))

	var StockPrice1StrctArrPtr *StockPriceStrct = &StockPrice1StrctArrVh
	StockPrice1StrctArrPtr.stock.id = 5
	StockPrice1StrctArrPtr.price = 2000
	StockPrice1StrctArrPtr.stockDate, _ = time.Parse(shortform, "09-10-2020")
	fmt.Println("Stock1 pointer struct: ", StockPrice1StrctArrPtr, " price:", StockPrice1StrctArrPtr.price, " date:", StockPrice1StrctArrPtr.stockDate.Format(yearform))
}

type AgeVh int

func MethodFn() {
	var Age1Vh AgeVh
	var Age2Vh AgeVh = 22
	fmt.Println("Age1 before: ", Age1Vh)
	Age1Vh.Age1InitFn()
	Age2Vh.Age1InitFn()
}
func (InitPtr *AgeVh) Age1InitFn() {
	if *InitPtr == 0 {
		*InitPtr = 18
	}
	fmt.Println("After init: ", *InitPtr)
}

type Person struct {
	name   string
	age    uint16
	gender string
	salary uint32
}
type PersonRedis struct {
	Name   string
	Age    uint16
	Gender string
	Salary uint32
}

func (personPtr *Person) PersonInitFn(PersonArg Person) {
	personPtr.name = PersonArg.name
	personPtr.gender = PersonArg.gender
	personPtr.age = PersonArg.age
	personPtr.salary = PersonArg.salary
}

func (personPtr *Person) personGetNameFn(companyNameArg string) {
	personPtr.name = "Arvind Rao"
	fmt.Println("Name: ", personPtr.name, " Age: ", personPtr.age, " works at", companyNameArg)
}

func (personPtr *Person) personGetSalaryFn(companySalaryArg uint32) {
	personPtr.salary = companySalaryArg
	fmt.Println("My new salary: ", personPtr.salary)
}

func PersonFn() {
	var PersonArvindVh Person
	PersonArvindVh.name = "Arvind"
	PersonArvindVh.gender = "Male"
	PersonArvindVh.age = 23
	PersonArvindVh.salary = 450000
	PersonArvindVh.PersonInitFn(PersonArvindVh)
	var compName = "Full Creative"
	PersonArvindVh.personGetNameFn(compName)
	fmt.Println("Outer\nName: ", PersonArvindVh.name, " Age: ", PersonArvindVh.age, " works at", compName)
	var compSalary = 600000
	PersonArvindVh.personGetSalaryFn(uint32(compSalary))
	fmt.Println("Outer\nSalary: ", PersonArvindVh.salary)
}

func PersonRedisExample(rdClient *redis.Client, ctx context.Context) {
	jsonPersonVh, err := json.Marshal(PersonRedis{
		Name:   "Arvind",
		Age:    23,
		Gender: "Male",
		Salary: 460000,
	})

	if err != nil {
		fmt.Println("Json error", err)
	}

	err = rdClient.Set(ctx, "id1234", string(jsonPersonVh), 0).Err()

	if err != nil {
		fmt.Println("Set error", err)
	}

	getPersonVh, err := rdClient.Get(ctx, "id1234").Result()
	if err != nil {
		fmt.Println("get error", err)
	}

	fmt.Println("Redis program is successful! Value: ", string(getPersonVh))
}

func ConvertTimeZoneFn() {
	UTCTimeVh, ErrVh := time.Parse("2006 01 02 15 04", "2022 01 12 22 10") // reference time & time to be converted
	// UTCtimeVh holds in UTCTime Zone in time format
	if ErrVh != nil {
		fmt.Println("Time error: ", ErrVh)
		return
	}
	fmt.Println("UTC time: ", UTCTimeVh)
	LocLATZVh, ErrVh := time.LoadLocation("America/Los_Angeles") // loads the time zone format to LocLATZVh
	fmt.Println("LA location: ", LocLATZVh)
	UTCTimeVh = UTCTimeVh.In(LocLATZVh) // In converts UTC into the given time zone
	fmt.Println("UTC time converted to LA time: ", UTCTimeVh.Format("02-01-2006 15:04"))

	LocLINTZVh, ErrVh := time.LoadLocation("Asia/Kolkata") // loads the time zone format to LocLATZVh
	fmt.Println("India location: ", LocLINTZVh)
	UTCTimeVh = UTCTimeVh.In(LocLINTZVh) // In converts UTC into the given time zone
	fmt.Println("IST time: ", UTCTimeVh.Format("02-01-2006 15:04"))
	fmt.Println("IST time RFC: ", UTCTimeVh.Format(time.RFC822))
	fmt.Println("IST time Kitchen: ", UTCTimeVh.Format(time.Kitchen))
	fmt.Println("IST time RFCZ: ", UTCTimeVh.Format(time.RFC1123Z))
}
