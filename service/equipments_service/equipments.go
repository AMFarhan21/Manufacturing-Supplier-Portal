package equipments_service

type Equipments struct {
	Id            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string  `json:"name"`
	CategoryId    int     `json:"category_id"`
	Description   string  `json:"description"`
	PricePerDay   float64 `json:"price_per_day"`
	PricePerWeek  float64 `json:"price_per_week"`
	PricePerMonth float64 `json:"price_per_month"`
	PricePerYear  float64 `json:"price_per_year"`
	Available     *bool   `json:"available"`
}
