package csbalancing

import (
	"reflect"
	"sort"
)

type Entity struct {
	ID    int
	Score int
}

type CustomerSuccessAway []int
type CustomerService map[int]int

func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway CustomerSuccessAway) int {
	customerServices := make(CustomerService)
	var customerSuccessWork []Entity
	
	for _, cs := range customerSuccess {
		if !inArray(customerSuccessAway, cs.ID) {
			customerSuccessWork = append(customerSuccessWork, cs)
		}
	}

	sort.Slice(customerSuccessWork, func(i, j int) bool {
		return customerSuccessWork[i].Score < customerSuccessWork[j].Score
	})

	for _, cs := range customerSuccessWork {
		for i, count, len := 0, 0, len(customers); i < len; i++ {
			j := i - count

			if cs.Score >= customers[j].Score {
				_, exists := customerServices[cs.ID]

				if !exists {
					customerServices[cs.ID] = 0;
				}
				customerServices[cs.ID] += 1;

				customers = removeElementByIndex(customers, j)
				count++
			}
		}
	}

	greaterAmountService := 0
	idCssGreaterAmountService := 0

	for idCss, totalCustomerService := range customerServices {
		if totalCustomerService > greaterAmountService {
			greaterAmountService = totalCustomerService
			idCssGreaterAmountService = idCss
		}
	}
	
	sameQuantity := false
	for idCss, totalCustomerService := range customerServices {
		if totalCustomerService == greaterAmountService && idCss != idCssGreaterAmountService {
			sameQuantity = true
			break
		}
	}

	 if sameQuantity {
		idCssGreaterAmountService = 0
	 }

	return idCssGreaterAmountService
}

func inArray(array interface{}, val interface{}) (exists bool) {
    exists = false

    switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)

			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
					exists = true
					return
				}
			}
		}
		
    return
}

func removeElementByIndex(slice []Entity, index int) []Entity {
	sliceLen := len(slice)
	sliceLastIndex := sliceLen - 1

	if index != sliceLastIndex {
		slice[index] = slice[sliceLastIndex]
	}

	return slice[:sliceLastIndex]
}