CREATE TABLE products (
    id SERIAL PRIMARY KEY,                      -- Auto-incremented unique ID for each product
    user_id INT REFERENCES users(id),           -- Foreign key reference to the users table
    product_name VARCHAR(255) NOT NULL,          -- Product name (string)
    product_description TEXT,                   -- Product description (text)
    product_images TEXT[] NOT NULL,              -- Array of image URLs (text[] array)
    compressed_product_images TEXT[],           -- Array of compressed image URLs (text[] array)
    product_price DECIMAL(10, 2) NOT NULL,       -- Product price (decimal with two decimal places)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for when the product was created
);
