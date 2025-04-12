package main

import (
	"fmt"
	"net/http"
	"time"
)

// Cache-Aside подход подразумевает, что приложение сначала проверяет наличие данных в кеше, и если их там нет, то загружает их из основного источника и сохраняет в кеше
func getData(key string) Data {
	data, found := cache.Get(key)
	if !found {
		// Загрузка из базы данных и сохранение в кэше
		data = loadDataFromDB(key)
		cache.Set(key, data, cache.DefaultExpiration)
	}
	return data
}

// Допустим, у нас есть приложение на Go, где мы хотим кешировать данные сессии.
// Мы можем использовать middleware, который будет проверять кеш перед выполнением запроса и сохранять данные сессии в кеш для быстрого доступа
// Этот код проверяет, есть ли данные сессии в кеше, и если они там, использует их,
// в противном случае продолжает обычную обработку запроса, создает сессию и записывает её в хранилище
func SessionCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := getSessionIDFromRequest(r)
		if sessionData, found := cache.Get(sessionId); found {
			// Используем данные сессии из кеша
			processWithCachedSessionData(sessionData)
		} else {
			// Обрабатываем запрос без кешированной сессии
			next.ServeHTTP(w, r)
		}
	})
}

// веб-сервис, который предоставляет информацию о продуктах. Эти данные регулярно запрашиваются, поэтому их кеширование может значительно улучшить производительность
func getProduct(productId int) Product {
	cacheKey := fmt.Sprintf("product_%d", productId)
	if product, found := cache.Get(cacheKey); found {
		// Возвращаем данные из кеша
		return product.(Product)
	}

	// Данных в кеше нет, загружаем из БД
	product := loadProductFromDB(productId)
	// Сохраняем данные в кеш
	cache.Set(cacheKey, product, cache.DefaultExpiration)
	return product
}

// Предположим, наше приложение регулярно запрашивает погоду для определённых городов.
// Вместо того, чтобы каждый раз отправлять запрос к внешнему API, мы можем кешировать эти данные.
// В этом случае, используя TTL (время жизни) для кеша, мы можем гарантировать,
// что данные о погоде остаются достаточно актуальными, но при этом не требуют постоянных запросов к API
func getWeather(city string) Weather {
	cacheKey := fmt.Sprintf("weather_%s", city)
	if weather, found := cache.Get(cacheKey); found {
		// Возвращаем погоду из кеша
		return weather.(Weather)
	}

	// Данных в кеше нет, делаем запрос к API
	weather := fetchWeatherFromAPI(city)
	// Сохраняем данные в кеш
	cache.Set(cacheKey, weather, 1*time.Hour) // Устанавливаем TTL
	return weather
}
