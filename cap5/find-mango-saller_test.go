package cap5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMangoSaller(t *testing.T) {
	t.Run("exist mango seller", func(t *testing.T) {

		mainSaller := Seller{
			Name:       "Main",
			SellerType: "apple",
		}

		secound := Seller{
			Name:       "secound",
			SellerType: "apple",
		}

		third := Seller{
			Name:       "third",
			SellerType: "apple",
		}

		fourth := Seller{
			Name:       "fourth",
			SellerType: "apple",
		}

		fifth := Seller{
			Name:       "fifth",
			SellerType: "Mango",
		}

		mainSaller.addFrinds(&secound, &third)
		secound.addFrinds(&fourth)
		fourth.addFrinds(&fifth)

		fmt.Println(mainSaller)

		result := FindMangoSaller(mainSaller)

		assert.Equal(t, fifth.Name, result.Name, "Fifth is mango seller")

	})

	t.Run("loop infint test", func(t *testing.T) {

		mainSaller := Seller{
			Name:       "Main",
			SellerType: "apple",
		}

		secound := Seller{
			Name:       "secound",
			SellerType: "apple",
		}

		third := Seller{
			Name:       "third",
			SellerType: "apple",
		}

		fourth := Seller{
			Name:       "fourth",
			SellerType: "apple",
		}

		fifth := Seller{
			Name:       "fifth",
			SellerType: "Mango",
		}

		mainSaller.addFrinds(&secound, &third)
		secound.addFrinds(&fourth, &mainSaller)
		fourth.addFrinds(&fifth)

		fmt.Println(mainSaller)

		result := FindMangoSaller(mainSaller)

		assert.Equal(t, fifth.Name, result.Name, "Fifth is mango seller")

	})

	t.Run("not exist mango seller", func(t *testing.T) {

		mainSaller := Seller{
			Name:       "Main",
			SellerType: "apple",
		}

		secound := Seller{
			Name:       "secound",
			SellerType: "apple",
		}

		third := Seller{
			Name:       "third",
			SellerType: "apple",
		}

		fourth := Seller{
			Name:       "fourth",
			SellerType: "apple",
		}

		fifth := Seller{
			Name:       "fifth",
			SellerType: "apple",
		}

		mainSaller.addFrinds(&secound, &third)
		secound.addFrinds(&fourth)
		fourth.addFrinds(&fifth)

		fmt.Println(mainSaller)

		result := FindMangoSaller(mainSaller)

		assert.Equal(t, "", result.Name, "not exist mango seller")

	})

}
