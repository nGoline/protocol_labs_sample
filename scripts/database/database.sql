CREATE DATABASE qcobtc;

use qcobtc;

-- Table Exchanges

CREATE TABLE `exchanges` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

insert into exchanges (name) values ('Mercado Bitcoin');
insert into exchanges (name) values ('Bitcoin Trade');

-----------------------------------------------------------
-- Table Pairs

CREATE TABLE `pairs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(8) DEFAULT NULL,
  `base` varchar(4) DEFAULT NULL,
  `quote` varchar(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

insert into pairs (name, base, quote) values ('BTCBRL', 'BTC', 'BRL');

-----------------------------------------------------------
-- Table Trades

CREATE TABLE `trades` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `price` double(20,5) NOT NULL,
  `amount` bigint(20) NOT NULL,
  `date` bigint(20) NOT NULL,
  `created_at` datetime(3) NOT NULL DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT NULL,
  `exchange_id` int(10) unsigned DEFAULT NULL,
  `pair_id` int(10) unsigned DEFAULT NULL,
  `block_height` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_exchange` (`exchange_id`),
  KEY `fk_pair` (`pair_id`),
  KEY `fk_block` (`block_height`),
  KEY `date` (`date`),
  CONSTRAINT `fk_exchange` FOREIGN KEY (`exchange_id`) REFERENCES `exchanges` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_pair` FOREIGN KEY (`pair_id`) REFERENCES `pairs` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

insert into trades (price, amount, date, exchange_id, pair_id) values (46983.7402000711, 23691092, 1588291708, 1, 1);
insert into trades (price, amount, date, exchange_id, pair_id) values (47345.1671057835, 359054, 1588291749, 2, 1);
