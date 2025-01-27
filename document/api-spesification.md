# API Specification: Product Management

## 1. Get Product List
- **Endpoint**: `/api/v1/product`
- **Method**: `GET`
- **Description**: Retrieves a list of products based on optional filters and sorting options.

### Request Parameters
- **`category`** (optional): The name of the product category to filter by (e.g., "minuman", "sayuran").
- **`name`** (optional): The name of the product to search for. This will filter products whose names contain the given string (case-insensitive).
- **`id`** (optional): The ID of the product to retrieve.
- **`sort`** (optional): Specifies the sorting order. You can provide:
  - `price_asc`: Sort by price in ascending order.
  - `price_desc`: Sort by price in descending order.
  - `name_asc`: Sort by product name in ascending order.
  - `name_desc`: Sort by product name in descending order.
  - `date_asc`: Sort by creation date in ascending order.
  - `date_desc`: Sort by creation date in descending order.
  - Default is `id_asc`.

### Example Request
```bash
curl -X GET \
  'http://localhost:8001/api/v1/product?category=minuman&name=Bayam%20Hijau&id=84&sort=price_desc' \
  --header 'Accept: */*' 
```

### Response
- **`Status Code : `** 200 OK if the request is successful.
Content-Type: application/json
- **`Content-Type : `** application/json
-  **`Response Body : `**
   -  An array of products matching the filter criteria.
   
   **`Response Body : `**
   ```json
    [
        {
            "id": 84,
            "name": "Bayam Hijau",
            "category_id": 2,
            "price": 5000,
            "created_at": "2023-01-15T14:00:00Z",
            "updated_at": "2023-02-01T16:00:00Z",
            "category": {
            "id": 2,
            "name": "Sayuran"
            }
        },
        {
            "id": 85,
            "name": "Bayam Merah",
            "category_id": 2,
            "price": 4500,
            "created_at": "2023-01-10T14:00:00Z",
            "updated_at": "2023-02-02T16:00:00Z",
            "category": {
            "id": 2,
            "name": "Sayuran"
            }
        }
    ]

   ```
## 2. Create Product
- **Endpoint**: `/api/v1/product`
- **Method**: `POST`
- **Description**: Adds a new product to the system.
  
### Request Body
- **`name`** (required): The name of the product.
- **`price`** (required): The price of the product.
- **`category_id`** (required): The ID of the category the product belongs to (e.g., "Sayuran", "Buah").

### Example Request Body
```json
{
    "name": "Bayam Hijau",
    "price": 5000,
    "category_id": 2,
}
```

### Example Request
```bash
curl -X POST \
  'http://localhost:8001/api/v1/product' \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Bayam Hijau",
    "price": 5000,
    "category_id": 2
  }'
```

### Response
- **`Status Code:`** 201 Created if the product is successfully created.

## 2. Get Category List
- **Endpoint**: `/api/v1/category`
- **Method**: `GET`
- **Description**: Retrieves all categories 

### Example Request
```bash
curl -X GET \
  'http://localhost:8001/api/v1/category' \
  --header 'Accept: */*' 
```

### Response
- **`Status Code : `** 200 OK if the request is successful.
Content-Type: application/json
- **`Content-Type : `** application/json
-  **`Response Body : `**
   -  An array of categories.
   
   **`Response Body : `**
   ```json
    [
        {
            "ID": 1,
            "Name": "Sayuran",
        },
        {
            "ID": 2,
            "Name": "Protein",
        },
    ]

   ```

## 2. Create Category
- **Endpoint**: `/api/v1/category`
- **Method**: `POST`
- **Description**: Adds a new category to the system.
  
### Request Body
- **`name`** (required): The name of the category.

### Example Request Body
```json
{
    "name": "Minuman",
}
```

### Example Request
```bash
curl -X POST \
  'http://localhost:8001/api/v1/category' \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Minuman"
  }'
```

### Response
- **`Status Code:`** 201 Created if the category is successfully created.