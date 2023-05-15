package controllers

import (
	"ecommerce/database"
	"ecommerce/model"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetProductDetails(ctx *gin.Context) {
	// Retrieve product information from the database based on the product ID
	productID := ctx.Param("id")
	product, err := GetProductsByID(productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Retrieve related products based on the recommendation algorithm
	relatedProducts, err := GetRelatedProducts(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving related products"})
		return
	}

	// Display the product details and recommended products to the user
	ctx.JSON(http.StatusOK, gin.H{
		"product":           product,
		"recommended_items": relatedProducts,
	})
}

// GetRelatedProducts returns the top 5 related products based on the input product
func GetRelatedProducts(product model.Products) ([]model.Products, error) {
	// Retrieve all products with the same category as the input product
	relatedProducts, err := GetProductsByCategory(product.Category)
	if err != nil {
		return nil, err
	}

	// check if relatedProducts is empty
	if len(relatedProducts) == 0 {
		return nil, errors.New("no related products found")
	}

	// Collect user behavior data
	userBehaviorData, err := GetUserBehaviorData()
	if err != nil {
		return nil, err
	}

	// Analyze user behavior data to find related products
	// Analyze user behavior data to find related products
	relatedProductIDs := make(map[string]float64)
	for _, behavior := range userBehaviorData {
		// Only consider behaviors related to the same category as the input product
		if behavior.Category == product.Category {
			// Convert the attributes string to a map[string]float64
			attributes := make(map[string]float64)
			err := json.Unmarshal([]byte(behavior.Attributes), &attributes)
			if err != nil {
				return nil, fmt.Errorf("error processing user behavior data")
			}

			// Use cosine similarity to measure the similarity between the input product and each behavior
			similarity := cosineSimilarity(product.Attributes, behavior.Attributes)

			// Add the behavior product to the related product list if it's not the input product and has a non-zero similarity score
			if similarity > 0 && behavior.ID != product.ID {
				relatedProductIDs[behavior.UserID] = similarity
			}
		}
	}

	// Sort the related product list by similarity score
	sortedProductIDs := sortMapByValue(relatedProductIDs)

	// Retrieve the top 5 related products by ID
	var relatedProductsIDs []string
	for id := range sortedProductIDs {
		relatedProductsIDs = append(relatedProductsIDs, id)
		if len(relatedProductsIDs) >= 5 {
			break
		}
	}

	// Retrieve the related products by ID
	relatedProducts, err = GetProductsByIDs(relatedProductsIDs)
	if err != nil {
		return nil, err
	}

	return relatedProducts, nil
}

func cosineSimilarity(s1, s2 string) float64 {
	// Split the strings into words
	words1 := strings.Fields(s1)
	words2 := strings.Fields(s2)

	// Count the frequency of occurrence of each word in the strings
	freq1 := make(map[string]float64)
	freq2 := make(map[string]float64)
	for _, word := range words1 {
		freq1[word]++
	}
	for _, word := range words2 {
		freq2[word]++
	}

	// Compute the cosine similarity between the frequency maps
	return cosineSimilarityMap(freq1, freq2)
}

func cosineSimilarityMap(v1, v2 map[string]float64) float64 {
	var dotProduct, magnitude1, magnitude2 float64
	for key, value := range v1 {
		dotProduct += value * v2[key]
		magnitude1 += value * value
	}
	for _, value := range v2 {
		magnitude2 += value * value
	}
	return dotProduct / (math.Sqrt(magnitude1) * math.Sqrt(magnitude2))
}

// sortMapByValue sorts a map by value in descending order and returns a new map
func sortMapByValue(m map[string]float64) map[string]float64 {
	var sortedKeys []string
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Slice(sortedKeys, func(i, j int) bool {
		return m[sortedKeys[i]] > m[sortedKeys[j]]
	})

	result := make(map[string]float64)
	for _, k := range sortedKeys {
		result[k] = m[k]
	}
	return result
}

func GetProductsByCategory(category string) ([]model.Products, error) {
	// Query the database for all products with the given category
	// and return the results
	var products []model.Products
	if err := database.DB.Where("category = ?", category).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductsByID(productID string) (model.Products, error) {
	// Query the database for the product with the given ID
	var product model.Products
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func GetProductsByIDs(productIDs []string) ([]model.Products, error) {
	// Convert the input slice to a comma-separated string
	productIDsStr := strings.Join(productIDs, ",")

	// Query the database for all products with the given IDs
	// and return the results
	var products []model.Products
	if err := database.DB.Where("id = ?", productIDsStr).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetUserBehaviorData() ([]model.WebTracking, error) {
	// Query the database for all user behavior data
	var userBehaviorData []model.WebTracking
	if err := database.DB.Find(&userBehaviorData).Error; err != nil {
		return nil, err
	}
	return userBehaviorData, nil
}
