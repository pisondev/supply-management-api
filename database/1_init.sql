
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. fungsi trigger utk auto updated_at
CREATE OR REPLACE FUNCTION update_timestamp_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

-- 2. tabelnya
CREATE TABLE suppliers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    contact_person VARCHAR(255),
    phone VARCHAR(50),
    address TEXT,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE ingredients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    unit_id INT NOT NULL REFERENCES units(id),
    is_perishable BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    location TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory_movements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    warehouse_id INT NOT NULL REFERENCES warehouses(id),
    ingredient_id UUID NOT NULL REFERENCES ingredients(id),
    supplier_id UUID REFERENCES suppliers(id),
    movement_type VARCHAR(20) NOT NULL CHECK (movement_type IN ('IN', 'OUT', 'ADJUSTMENT')),
    quantity NUMERIC(12,2) NOT NULL CHECK (quantity > 0),
    balance_after NUMERIC(12,2) NOT NULL CHECK (balance_after >= 0),
    reference_code VARCHAR(100),
    notes TEXT,
    movement_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT check_out_supplier CHECK (movement_type != 'OUT' OR supplier_id IS NULL)
);

CREATE TABLE inventories (
    warehouse_id INT NOT NULL REFERENCES warehouses(id),
    ingredient_id UUID NOT NULL REFERENCES ingredients(id),
    stock_level NUMERIC(12,2) NOT NULL DEFAULT 0 CHECK (stock_level >= 0),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (warehouse_id, ingredient_id)
);

-- 3. index buat performa
CREATE INDEX idx_inv_movements_ingredient_id ON inventory_movements(ingredient_id);
CREATE INDEX idx_inv_movements_warehouse_id ON inventory_movements(warehouse_id);
CREATE INDEX idx_inv_movements_supplier_id ON inventory_movements(supplier_id);
CREATE INDEX idx_inv_movements_movement_date ON inventory_movements(movement_date);

-- 4. trigger tabel tabel
CREATE TRIGGER update_suppliers_updated_at
BEFORE UPDATE ON suppliers FOR EACH ROW EXECUTE FUNCTION update_timestamp_column();

CREATE TRIGGER update_ingredients_updated_at
BEFORE UPDATE ON ingredients FOR EACH ROW EXECUTE FUNCTION update_timestamp_column();

CREATE TRIGGER update_warehouses_updated_at
BEFORE UPDATE ON warehouses FOR EACH ROW EXECUTE FUNCTION update_timestamp_column();

CREATE TRIGGER update_inventories_updated_at
BEFORE UPDATE ON inventories FOR EACH ROW EXECUTE FUNCTION update_timestamp_column();

