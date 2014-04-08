delimiter $$

CREATE TABLE `skin_color` (
  `id` int(11) NOT NULL,
  `value` int(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8$$

delimiter $$

CREATE TABLE `skin` (
  `player_id` int(4) NOT NULL,
  `color_id` int(11) DEFAULT NULL,
  `bubble` int(11) DEFAULT NULL,
  `face` int(11) DEFAULT NULL,
  `vertex` int(11) DEFAULT NULL,
  `shadow` int(11) DEFAULT NULL,
  `surround` int(11) DEFAULT NULL,
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8$$
