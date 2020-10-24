package ingredients

type Ingredient interface {
	GetName() string
}

func GetIngredientName(obj Ingredient) string {
	return obj.GetName()
}
