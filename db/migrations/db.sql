-- Create tables
CREATE TABLE IF NOT EXISTS Customers (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Categories (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Products (
    id UUID PRIMARY KEY,
    category_id UUID REFERENCES Categories(id),
    name VARCHAR NOT NULL,
    description TEXT,
    price DECIMAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ShoppingCarts (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS CartItems (
    id UUID PRIMARY KEY,
    cart_id UUID REFERENCES ShoppingCarts(id),
    product_id UUID REFERENCES Products(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Orders (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES Customers(id),
    total_amount DECIMAL NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS OrderItems (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES Orders(id),
    product_id UUID REFERENCES Products(id),
    quantity INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert dummy data into Categories
INSERT INTO Categories (id, name)
VALUES
    ('f3f5e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Electronics'),
    ('b1d8e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Books'),
    ('c2e9e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Clothing');

-- Insert dummy data into Products
INSERT INTO Products (id, category_id, name, description, price)
VALUES
    ('d3f8e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'f3f5e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Smartphone', 'Latest model smartphone', 699.99),
    ('e4f9e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'b1d8e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Novel', 'Bestselling novel', 19.99),
    ('f5f0e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'c2e9e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'T-shirt', 'Cotton t-shirt', 9.99);

-- Insert dummy data into Customers
INSERT INTO Customers (id, name, email, password)
VALUES
    ('a6f0e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'John Doe', 'john.doe@example.com', 'password123'),
    ('b7f1e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Jane Smith', 'jane.smith@example.com', 'password456'),
    ('c8f2e8d3-1c4d-4b9a-911a-5a5e8d4c7a5b', 'Alice Johnson', 'alice.johnson@example.com', 'password789');
