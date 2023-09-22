package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WebService struct{}

func NewWebService() *WebService {
	return &WebService{}
}

func (api *WebService) fetchDetailsFromMercadoLibre(client *MeliClient, chunkItems []fileItem) ([]DetailedItem, error) {

	var ids []string
	for _, item := range chunkItems {
		ids = append(ids, item.Site+item.ID)
	}

	idsUrl := strings.Join(ids, ",")

	itemDetails, err := api.fetchItemDetails(client, idsUrl)
	if err != nil {
		fmt.Errorf("Error fetching item details: %v\n", err)
		return nil, err
	}

	var detailedItems []DetailedItem

	for _, item := range *itemDetails {

		if item.Code != 200 || item.Body.CategoryID == nil || item.Body.CurrencyID == "" || item.Body.SellerID == 0 {
			continue
		}

		// Get Category Name
		categoryName, err := api.fetchCategoryName(client, item.Body.CategoryID)
		if err != nil {
			fmt.Printf("Error fetching category name: %v\n", err)
			continue
		}

		// Get Currency Description
		currencyDescription, err := api.fetchCurrencyDescription(client, item.Body.CurrencyID)
		if err != nil {
			fmt.Printf("Error fetching currency description: %v\n", err)
			continue
		}

		// Get Seller Nickname
		sellerNickname, err := api.fetchSellerNickname(client, item.Body.SellerID)
		if err != nil {
			fmt.Printf("Error fetching seller nickname: %v\n", err)
			continue
		}

		detailedItem := DetailedItem{
			Site:       item.Body.SiteID,
			ID:         item.Body.ID,
			Price:      item.Body.Price,
			StartTime:  item.Body.DateCreated,
			CategoryID: categoryName,
			CurrencyID: currencyDescription,
			SellerID:   sellerNickname,
		}

		detailedItems = append(detailedItems, detailedItem)
	}

	return detailedItems, nil
}

func (api *WebService) fetchItemDetails(client *MeliClient, itemID string) (*[]ItemAPIResponse, error) {
	endpoint := fmt.Sprintf("/items?ids=%s", itemID)
	var apiResp []ItemAPIResponse
	if err := api.fetchAndDecode(client, endpoint, &apiResp); err != nil {
		fmt.Printf("Error fetching item details: %v\n", endpoint)
		fmt.Printf("Error fetching item details: %v\n", err)
		return nil, err
	}
	return &apiResp, nil
}

func (api *WebService) fetchCategoryName(client *MeliClient, categoryID interface{}) (string, error) {
	if categoryID == nil {
		return "", nil
	}
	endpoint := fmt.Sprintf("/categories/%s", categoryID)
	var categoryResp CategoryAPIResponse
	if err := api.fetchAndDecode(client, endpoint, &categoryResp); err != nil {
		return "", err
	}
	return categoryResp.Name, nil
}

func (api *WebService) fetchCurrencyDescription(client *MeliClient, currencyID string) (string, error) {
	endpoint := fmt.Sprintf("/currencies/%s", currencyID)
	var currencyResp CurrencyAPIResponse
	if err := api.fetchAndDecode(client, endpoint, &currencyResp); err != nil {
		return "", err
	}
	return currencyResp.Description, nil
}

func (api *WebService) fetchSellerNickname(client *MeliClient, sellerID int) (string, error) {
	endpoint := fmt.Sprintf("/users/%d", sellerID)
	var sellerResp SellerAPIResponse
	if err := api.fetchAndDecode(client, endpoint, &sellerResp); err != nil {
		return "", err
	}
	return sellerResp.Nickname, nil
}

func (api *WebService) fetchAndDecode(client *MeliClient, endpoint string, target interface{}) error {
	resp, err := client.Get(endpoint)
	if err != nil {
		return fmt.Errorf("error making request to MercadoLibre API: %w", err)
	}
	defer resp.Body.Close()
	return decodeResponseBody(resp, target)
}

// Resto del código sin cambios

func decodeResponseBody(resp *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading API response: %w", err)
	}
	return json.Unmarshal(body, target)
}

// ItemAPIResponse representa la estructura de la respuesta de la API de MercadoLibre
type ItemAPIResponse struct {
	Code int
	Body struct {
		ID          string      `json:"id"`
		SiteID      string      `json:"site_id"`
		Price       float64     `json:"price"`
		DateCreated string      `json:"date_created"`
		CategoryID  interface{} `json:"category_id"`
		CurrencyID  string      `json:"currency_id"`
		SellerID    int         `json:"seller_id"`
	}
}

// DetailedItem representa un ítem con detalles obtenidos de la API
type DetailedItem struct {
	Site       string
	ID         string
	Price      float64
	StartTime  string
	CategoryID interface{}
	CurrencyID string
	SellerID   string
}

type CategoryAPIResponse struct {
	Id                       string        `json:"id"`
	Name                     string        `json:"name"`
	Picture                  interface{}   `json:"picture"`
	Permalink                interface{}   `json:"permalink"`
	TotalItemsInThisCategory int           `json:"total_items_in_this_category"`
	PathFromRoot             []interface{} `json:"path_from_root"`
	ChildrenCategories       []interface{} `json:"children_categories"`
	AttributeTypes           string        `json:"attribute_types"`
	Settings                 struct {
	} `json:"settings"`
	ChannelsSettings []interface{} `json:"channels_settings"`
	MetaCategId      interface{}   `json:"meta_categ_id"`
	Attributable     bool          `json:"attributable"`
	DateCreated      time.Time     `json:"date_created"`
}

type CurrencyAPIResponse struct {
	Id            string `json:"id"`
	Description   string `json:"description"`
	Symbol        string `json:"symbol"`
	DecimalPlaces int    `json:"decimal_places"`
}

type SellerAPIResponse struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname"`
	CountryId string `json:"country_id"`
	Address   struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	UserType         string `json:"user_type"`
	SiteId           string `json:"site_id"`
	Permalink        string `json:"permalink"`
	SellerReputation struct {
		LevelId           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Period string `json:"period"`
			Total  int    `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}
