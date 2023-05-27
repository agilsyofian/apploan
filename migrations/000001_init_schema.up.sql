CREATE TABLE `konsumen` (
  `id` varchar(60) NOT NULL,
  `username` varchar(10) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nik` varchar(30) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `legal_name` varchar(100) NOT NULL,
  `tempat_lahir` varchar(50) NOT NULL,
  `tgl_lahir` date NOT NULL,
  `gaji` decimal(10,2) NOT NULL,
  `foto_ktp` varchar(100) NOT NULL,
  `foto_selfie` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT uc_konsumen UNIQUE (`username`,`nik`)
);

-- CREATE TABLE `limit` (
--   `id` int(11) NOT NULL AUTO_INCREMENT,
--   `konsumen_id` varchar(60) NOT NULL,
--   `tenor` tinyint(4) NOT NULL,
--   `limit` decimal(10,2) NOT NULL,
--   `created_at` datetime NOT NULL,
--   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`),
--   KEY `limit_ibfk_1` (`konsumen_id`),
--   CONSTRAINT `limit_ibfk_1` FOREIGN KEY (`konsumen_id`) REFERENCES `konsumen` (`id`)
-- );

CREATE TABLE `kontrak` (
  `no` varchar(60) NOT NULL,
  `konsumen_id` varchar(60) NOT NULL,
  `otr` decimal(10,2) NOT NULL,
  `admin_fee` decimal(10,2) NOT NULL,
  `jml_cicilan` decimal(10,2) NOT NULL,
  `jml_bunga` decimal(5,2) NOT NULL,
  `nama_asset` varchar(100) NOT NULL,
  `status` enum('inpg','done','cancel','fail') NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`no`) USING BTREE,
  KEY `kontrak_ibfk1` (`konsumen_id`),
  CONSTRAINT `kontrak_ibfk1` FOREIGN KEY (`konsumen_id`) REFERENCES `konsumen` (`id`)
);

CREATE TABLE `transaksi` (
  `id` varchar(60) NOT NULL,
  `kontrak_no` varchar(60) NOT NULL,
  `jenis` enum('debit','kredit') NOT NULL,
  `jml` decimal(10,2) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `transaksi_ibfk1` (`kontrak_no`),
  CONSTRAINT `transaksi_ibfk1` FOREIGN KEY (`kontrak_no`) REFERENCES `kontrak` (`no`)
);

CREATE TABLE `session` (
  `id` text NOT NULL,
  `konsumen_id` varchar(60) NOT NULL,
  `refresh_token` text NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `client_ip` varchar(50) NOT NULL,
  `expired_at` datetime NOT NULL,
  `isblocked` tinyint(1) NOT NULL DEFAULT '0',
  KEY `session_ibfk1` (`konsumen_id`),
  CONSTRAINT `session_ibfk1` FOREIGN KEY (`konsumen_id`) REFERENCES `konsumen` (`id`)
);