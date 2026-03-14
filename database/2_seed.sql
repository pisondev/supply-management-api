
-- 1. seed units table
INSERT INTO units (id, name, description) VALUES
(1, 'kg', 'Kilogram (contoh: beras, telor, ikan, daging)'),
(2, 'liter', 'Liter (contoh: susu, minyak goreng)'),
(3, 'pcs', 'Pieces/Satuan (contoh: kemasan utuh)'),
(4, 'ikat', 'Ikat (contoh: sayur"an)');

SELECT setval('units_id_seq', 4);

-- 2. seed warehouses table
INSERT INTO warehouses (id, name, location) VALUES
(1, 'Dapur Pusat MBG Kota', 'Jl. Sudirman, Yogyakarta'),
(2, 'Gudang Transit Sektor Barat', 'Sentolo, Kulon Progo');

SELECT setval('warehouses_id_seq', 2);

-- 3. seed suppliers table
INSERT INTO suppliers (id, name, contact_person, phone, address, status) VALUES
('a1111111-1111-1111-1111-111111111111', 'Koperasi Pak Kevin', 'Pak Kevin', '081234567890', 'Sleman', 'active'),
('a2222222-2222-2222-2222-222222222222', 'Peternakan Ayam Ayaman', 'Bu Aliya', '081298765432', 'Bantul', 'active'),
('a3333333-3333-3333-3333-333333333333', 'Kelompok Budidaya Lele', 'Pak Pison', '081555666777', 'Kulon Progo', 'active');

-- 4. seed ingredients table
INSERT INTO ingredients (id, sku, name, unit_id, is_perishable) VALUES
('b1111111-1111-1111-1111-111111111111', 'RICE-P01', 'Beras Putih Salju', 1, false),
('b2222222-2222-2222-2222-222222222222', 'EGG-M01', 'Telur Murah', 1, true),
('b3333333-3333-3333-3333-333333333333', 'VEG-H01', 'Sayur Asal Warna Hijau', 4, true),
('b4444444-4444-4444-4444-444444444444', 'FSH-L01', 'Ikan Lele Segar', 1, true);

-- 5. seed inventory movements table
INSERT INTO inventory_movements (warehouse_id, ingredient_id, supplier_id, movement_type, quantity, balance_after, reference_code, notes) VALUES
(1, 'b1111111-1111-1111-1111-111111111111', 'a1111111-1111-1111-1111-111111111111', 'IN', 1500.00, 1500.00, 'PO-MBG-001', 'pasokan beras awal bulan'),
(1, 'b2222222-2222-2222-2222-222222222222', 'a2222222-2222-2222-2222-222222222222', 'IN', 500.00, 500.00, 'PO-MBG-002', 'pasokan telur segar'),
(2, 'b4444444-4444-4444-4444-444444444444', 'a3333333-3333-3333-3333-333333333333', 'IN', 250.00, 250.00, 'PO-MBG-003', 'pasokan lele fresh');

-- 6. seed inventories table
INSERT INTO inventories (warehouse_id, ingredient_id, stock_level) VALUES
(1, 'b1111111-1111-1111-1111-111111111111', 1500.00),
(1, 'b2222222-2222-2222-2222-222222222222', 500.00),
(2, 'b4444444-4444-4444-4444-444444444444', 250.00);