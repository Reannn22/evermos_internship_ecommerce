ALTER TABLE `trx_detail` 
ADD COLUMN `product_status` varchar(255) NOT NULL DEFAULT 'shipping...' AFTER `harga_total`;

-- Add this line to update existing records
UPDATE `trx_detail` SET `product_status` = 'shipping...' WHERE `product_status` IS NULL OR `product_status` = '';
