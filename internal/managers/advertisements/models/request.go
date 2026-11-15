package models

const (
	FormFieldTitle       = "title"
	FormFieldDescription = "description"
	FormFieldPrice       = "price"
	FormFieldQuantity    = "quantity"
	FormFieldPhotos      = "photos"
)

type AdvertisementRequest struct {
	Title       string
	Description string
	Price       string
	Quantity    string
	PhotoPaths  []string
	OmitFields  []string
}

// FormFields возвращает текстовые поля multipart-формы с учётом OmitFields.
func (r AdvertisementRequest) FormFields() map[string]string {
	fields := map[string]string{
		FormFieldTitle:       r.Title,
		FormFieldDescription: r.Description,
		FormFieldPrice:       r.Price,
		FormFieldQuantity:    r.Quantity,
	}

	for _, name := range r.OmitFields {
		delete(fields, name)
	}

	return fields
}
