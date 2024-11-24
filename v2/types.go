package v2

// ResponseEntity Информация о заказе
type ResponseEntity struct {
	// Uuid Идентификатор заказа в ИС СДЭК
	Uuid string `json:"uuid,omitempty"`
	// Comment комментарий
	Comment string `json:"comment,omitempty"`
}

type ResponseErr struct {
	// Message Описание ошибки
	Message string `json:"message"`
	// Code string Код ошибки
	Code string `json:"code"`
}

// ResponseRequests Информация о запросе над заказом
type ResponseRequests struct {
	// RequestUuid Идентификатор запроса в ИС СДЭК
	RequestUuid string `json:"request_uuid,omitempty"`
	// Type Тип запроса. Может принимать значения: CREATE, UPDATE, DELETE, AUTH, GET
	Type string `json:"type"`
	// State Текущее состояние запроса. Может принимать значения:
	// ACCEPTED - пройдена предварительная валидация и запрос принят
	// WAITING - запрос ожидает обработки (зависит от выполнения другого запроса)
	// SUCCESSFUL - запрос обработан успешно
	// INVALID - запрос обработался с ошибкой
	State string `json:"state"`
	// DateTime Дата и время установки текущего состояния запроса (формат yyyy-MM-dd'T'HH:mm:ssZ)
	DateTime string `json:"date_time"`
	// Errors Ошибки, возникшие в ходе выполнения запроса
	Errors []ResponseErr `json:"errors,omitempty"`
	// Warnings Предупреждения, возникшие в ходе выполнения запроса
	Warnings []ResponseErr `json:"warnings,omitempty"`
}

// ResponseRelatedEntities Связанные сущности (если в запросе был передан корректный print)
type ResponseRelatedEntities struct {
	// Type Тип связанной сущности. Может принимать значения: waybill - квитанция к заказу, barcode - ШК места к заказу
	Type string `json:"type"`
	// Uuid Идентификатор сущности, связанной с заказом
	Uuid string `json:"uuid"`
	// Url Ссылка на скачивание печатной формы в статусе "Сформирован", только для type = waybill, barcode
	Url string `json:"url,omitempty"`
	// CdekNumber Номер заказа СДЭК. Может возвращаться для return_order, direct_order, reverse_order
	CdekNumber string `json:"cdek_number,omitempty"`
	// Date Дата доставки, согласованная с получателем. Только для типа delivery
	Date string `json:"date,omitempty"`
	// TimeFrom Время начала ожидания курьера (согласованное с получателем). Только для типа delivery
	TimeFrom string `json:"time_from,omitempty"`
	// Date Время окончания ожидания курьера (согласованное с получателем). Только для типа delivery
	TimeTo string `json:"time_to,omitempty"`
}

type Response struct {
	Entity          ResponseEntity           `json:"entity,omitempty"`
	Requests        []ResponseRequests       `json:"requests"`
	RelatedEntities *ResponseRelatedEntities `json:"related_entities,omitempty"`
}

type Location struct {
	// Code Код населенного пункта СДЭК (метод "Список населенных пунктов")
	Code int `json:"code,omitempty"`
	// FiasGuid Уникальный идентификатор ФИАС UUID
	FiasGuid string `json:"fias_guid,omitempty"`
	// PostalCode Почтовый индекс
	PostalCode string `json:"postal_code,omitempty"`
	// Longitude Долгота
	Longitude float64 `json:"longitude,omitempty"`
	// Latitude Широта
	Latitude float64 `json:"latitude,omitempty"`
	// CountryCode
	CountryCode string `json:"country_code,omitempty"`
	// Region Название региона
	Region string `json:"region,omitempty"`
	// RegionCode Код региона СДЭК
	RegionCode int `json:"region_code,omitempty"`
	// SubRegion Название района региона
	SubRegion string `json:"sub_region,omitempty"`
	// City Название города
	City string `json:"city,omitempty"`
	// Address Строка адреса
	Address string `json:"address"`
}

type Package struct {
	// Number Номер упаковки (можно использовать порядковый номер упаковки заказа или номер заказа), уникален в пределах заказа. Идентификатор заказа в ИС Клиента
	Number string `json:"number"`

	// Weight Общий вес (в граммах)
	Weight int `json:"weight"`

	// Высота (в сантиметрах). Поле обязательно если:
	// если общий вес >=100 гр
	Height int `json:"height,omitempty"`
	// Длина (в сантиметрах). Поле обязательно если:
	// если общий вес >=100 гр
	Length int `json:"length,omitempty"`
	// Ширина (в сантиметрах). Поле обязательно если:
	// если общий вес >=100 гр
	Width int `json:"width,omitempty"`

	// Items Позиции товаров в упаковке. Только для заказов "интернет-магазин". Максимум 126 уникальных позиций в заказе. Общее количество товаров в заказе может быть от 1 до 10000
	Items []PackageItem `json:"items,omitempty"`
}

type PackageItem struct {
	// Name Наименование товара (может также содержать описание товара: размер, цвет)
	Name string `json:"name"`

	// WareKey Идентификатор/артикул товара. Артикул товара может содержать только символы: [A-z А-я 0-9 ! @ " # № $ ; % ^ : & ? * () _ - + = ? < > , .{ } [ ] \ / , пробел]
	WareKey string `json:"ware_key"`

	// TODO: Нахуя ты?
	Marking string `json:"marking,omitempty"`

	Payment Payment `json:"payment"`

	// Cost Объявленная стоимость товара (за единицу товара в валюте взаиморасчетов, значение >=0). С данного значения рассчитывается страховка
	Cost float64 `json:"cost"`
	// Amount Количество единиц товара (в штуках). Количество одного товара в заказе может быть от 1 до 999
	Amount int `json:"amount"`
	// Weight Вес (за единицу товара, в граммах)
	Weight int `json:"weight"`
}

type Payment struct {
	// Всегда 0
	Value int `json:"value"`

	// TODO: Нахуя вы нужны?
	// VatSum Сумма НДС
	VatSum int `json:"vat_sum,omitempty"`
	// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
	VatRate int `json:"vat_rate,omitempty"`
}

type Cost struct {
	// Sum Доп. сбор за доставку товаров, общая стоимость которых попадает в интервал
	Sum int `json:"sum"`
	// Threshold Порог стоимости товара (действует по условию меньше или равно) в целых единицах валюты
	Threshold int `json:"threshold"`
	// VatSum Сумма НДС
	VatSum int `json:"vat_sum,omitempty"`
	// VatRate Ставка НДС (значение - 0, 10, 20, null - нет НДС)
	VatRate int `json:"vat_rate,omitempty"`
}

type Phone struct {
	Number string `json:"number"`
}

type RecipientSender struct {
	// Name нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Name string `json:"name,omitempty"`
	// Company Название компании. нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Company string `json:"company,omitempty"`
	// Email Эл. адрес. нет, если заказ типа "интернет-магазин"; да, если заказ типа "доставка"
	Email string `json:"email,omitempty"`
	// PassportSeries Серия паспорта
	PassportSeries string `json:"passport_series,omitempty"`
	// PassportNumber Номер паспорта
	PassportNumber string `json:"passport_number,omitempty"`
	// PassportDateOfIssue Дата выдачи паспорта
	PassportDateOfIssue string `json:"passport_date_of_issue,omitempty"`
	// PassportOrganization Орган выдачи паспорта
	PassportOrganization string `json:"passport_organization,omitempty"`
	// Tin ИНН Может содержать 10, либо 12 символов
	Tin string `json:"tin,omitempty"`
	// PassportDateOfBirth Дата рождения (yyyy-MM-dd)
	PassportDateOfBirth string `json:"passport_date_of_birth,omitempty"`
	// PassportRequirementsSatisfied Требования по паспортным данным удовлетворены (актуально для
	// международных заказов):
	// true - паспортные данные собраны или не требуются
	// false - паспортные данные требуются и не собраны
	PassportRequirementsSatisfied bool `json:"passport_requirements_satisfied,omitempty"`
	// Phones Список телефонов, Не более 10 номеров
	Phones []Phone `json:"phones,omitempty"`
}

type Contact struct {
	// ФИО
	Name   string  `json:"name"`
	Phones []Phone `json:"phones"`
}

type Seller struct {
	// Name Наименование истинного продавца. Обязателен если заполнен inn
	Name string `json:"name,omitempty"`
	// INN ИНН истинного продавца. Может содержать 10, либо 12 символов
	INN string `json:"inn,omitempty"`
	// Phone Телефон истинного продавца. Обязателен если заполнен inn
	Phone string `json:"phone,omitempty"`
	// OwnershipForm Код формы собственности (подробнее см. приложение 2). Обязателен если заполнен inn
	OwnershipForm int `json:"ownership_form,omitempty"`
	// Address Адрес истинного продавца. Используется при печати инвойсов для отображения адреса настоящего
	// продавца товара, либо торгового названия. Только для международных заказов "интернет-магазин".
	// Обязателен если заказ - международный
	Address string `json:"address,omitempty"`
}

type Service struct {
	// Code Тип дополнительной услуги (подробнее см. приложение 3)
	Code string `json:"code"`
	// Parameter Параметр дополнительной услуги:
	// количество для услуг
	// PACKAGE_1, COURIER_PACKAGE_A2, SECURE_PACKAGE_A2, SECURE_PACKAGE_A3, SECURE_PACKAGE_A4,
	// SECURE_PACKAGE_A5, CARTON_BOX_XS, CARTON_BOX_S, CARTON_BOX_M, CARTON_BOX_L, CARTON_BOX_500GR,
	// CARTON_BOX_1KG, CARTON_BOX_2KG, CARTON_BOX_3KG, CARTON_BOX_5KG, CARTON_BOX_10KG, CARTON_BOX_15KG,
	// CARTON_BOX_20KG, CARTON_BOX_30KG, CARTON_FILLER (для всех типов заказа)
	// объявленная стоимость заказа для услуги INSURANCE (только для заказов с типом "доставка")
	// длина для услуг BUBBLE_WRAP, WASTE_PAPER (для всех типов заказа)
	// номер телефона для услуги SMS
	// код фотопроекта для услуги PHOTO_DOCUMENT
	Parameter string `json:"parameter,omitempty"`
}
