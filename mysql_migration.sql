-- MySQL migration for persisted models used in ecommerce_item_processor
-- Generated from model structs registered in this project.

SET NAMES utf8mb4;

CREATE TABLE IF NOT EXISTS categories (
  category_id BIGINT NOT NULL AUTO_INCREMENT,
  category_name VARCHAR(40) NOT NULL,
  image_path VARCHAR(250) NOT NULL,
  icon VARCHAR(250) NOT NULL,
  description TEXT NULL,
  active TINYINT NOT NULL DEFAULT 1,
  date_created DATETIME NULL,
  date_modified DATETIME NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (category_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_prices (
  item_price_id BIGINT NOT NULL AUTO_INCREMENT,
  item_price FLOAT NOT NULL,
  alt_item_price FLOAT NOT NULL,
  show_alt_price BOOLEAN NOT NULL DEFAULT 0,
  discount VARCHAR(255) NOT NULL DEFAULT '',
  discount_type VARCHAR(255) NOT NULL DEFAULT '',
  extra_charges FLOAT NOT NULL,
  currency BIGINT NOT NULL,
  active INT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  PRIMARY KEY (item_price_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS status (
  status_id BIGINT NOT NULL AUTO_INCREMENT,
  status VARCHAR(128) NOT NULL,
  status_code VARCHAR(128) NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  active INT NOT NULL,
  PRIMARY KEY (status_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS items (
  item_id BIGINT NOT NULL AUTO_INCREMENT,
  item_name VARCHAR(80) NOT NULL,
  description VARCHAR(250) NOT NULL DEFAULT '',
  weight VARCHAR(20) NOT NULL DEFAULT '',
  category_id BIGINT NOT NULL,
  item_price_id BIGINT NULL,
  available_sizes VARCHAR(250) NOT NULL DEFAULT '',
  available_colors VARCHAR(250) NOT NULL DEFAULT '',
  material VARCHAR(400) NOT NULL DEFAULT '',
  image_path VARCHAR(250) NOT NULL DEFAULT '',
  quantity INT NOT NULL,
  active INT NOT NULL,
  date_created DATETIME NULL,
  date_modified DATETIME NULL,
  created_by INT NULL,
  modified_by INT NULL,
  country BIGINT NOT NULL,
  branch BIGINT NULL,
  item_status BIGINT NULL,
  last_order_date DATETIME NULL,
  PRIMARY KEY (item_id),
  KEY idx_items_category_id (category_id),
  KEY idx_items_item_price_id (item_price_id),
  KEY idx_items_item_status (item_status),
  KEY idx_items_branch (branch),
  CONSTRAINT fk_items_category FOREIGN KEY (category_id) REFERENCES categories(category_id),
  CONSTRAINT fk_items_item_price FOREIGN KEY (item_price_id) REFERENCES item_prices(item_price_id),
  CONSTRAINT fk_items_status FOREIGN KEY (item_status) REFERENCES status(status_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_quantity (
  item_quantity_id BIGINT NOT NULL AUTO_INCREMENT,
  item_id BIGINT NOT NULL,
  quantity INT NOT NULL,
  quantity_alert INT NULL,
  active INT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  PRIMARY KEY (item_quantity_id),
  UNIQUE KEY uq_item_quantity_item_id (item_id),
  CONSTRAINT fk_item_quantity_item FOREIGN KEY (item_id) REFERENCES items(item_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS features (
  feature_id BIGINT NOT NULL AUTO_INCREMENT,
  feature_name VARCHAR(40) NOT NULL,
  image_path VARCHAR(250) NOT NULL,
  visible BOOLEAN NOT NULL DEFAULT 0,
  description VARCHAR(250) NOT NULL DEFAULT '',
  active INT NOT NULL,
  date_created DATETIME NULL,
  date_modified DATETIME NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (feature_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_features (
  item_feature_id BIGINT NOT NULL AUTO_INCREMENT,
  item_id BIGINT NOT NULL,
  feature_id BIGINT NOT NULL,
  active INT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  PRIMARY KEY (item_feature_id),
  KEY idx_item_features_item_id (item_id),
  KEY idx_item_features_feature_id (feature_id),
  CONSTRAINT fk_item_features_item FOREIGN KEY (item_id) REFERENCES items(item_id),
  CONSTRAINT fk_item_features_feature FOREIGN KEY (feature_id) REFERENCES features(feature_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS purposes (
  purpose_id BIGINT NOT NULL AUTO_INCREMENT,
  purpose VARCHAR(40) NOT NULL,
  image_path VARCHAR(250) NOT NULL,
  visible BOOLEAN NOT NULL DEFAULT 0,
  description VARCHAR(250) NOT NULL DEFAULT '',
  active INT NOT NULL,
  date_created DATETIME NULL,
  date_modified DATETIME NULL,
  created_by INT NULL,
  modified_by INT NULL,
  PRIMARY KEY (purpose_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_purposes (
  item_purpose_id BIGINT NOT NULL AUTO_INCREMENT,
  item_id BIGINT NOT NULL,
  purpose_id BIGINT NOT NULL,
  active INT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  PRIMARY KEY (item_purpose_id),
  KEY idx_item_purposes_item_id (item_id),
  KEY idx_item_purposes_purpose_id (purpose_id),
  CONSTRAINT fk_item_purposes_item FOREIGN KEY (item_id) REFERENCES items(item_id),
  CONSTRAINT fk_item_purposes_purpose FOREIGN KEY (purpose_id) REFERENCES purposes(purpose_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_images (
  item_image_id BIGINT NOT NULL AUTO_INCREMENT,
  item_id BIGINT NOT NULL,
  image_path VARCHAR(250) NOT NULL,
  is_default INT NOT NULL,
  active INT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  PRIMARY KEY (item_image_id),
  KEY idx_item_images_item_id (item_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_reviews (
  item_review_id BIGINT NOT NULL AUTO_INCREMENT,
  review VARCHAR(1000) NOT NULL,
  rating DOUBLE NOT NULL,
  review_by BIGINT NULL,
  item_id BIGINT NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  reference BIGINT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  active INT NOT NULL,
  PRIMARY KEY (item_review_id),
  KEY idx_item_reviews_item_id (item_id),
  KEY idx_item_reviews_reference (reference),
  CONSTRAINT fk_item_reviews_item FOREIGN KEY (item_id) REFERENCES items(item_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS item_types (
  item_type_id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(80) NOT NULL,
  description VARCHAR(255) NOT NULL,
  date_created DATETIME NOT NULL,
  date_modified DATETIME NOT NULL,
  created_by INT NOT NULL,
  modified_by INT NOT NULL,
  active INT NOT NULL,
  PRIMARY KEY (item_type_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
