CREATE DATABASE IF NOT EXISTS `tfg`; 
USE `tfg`;

-- Creación de tablas
CREATE TABLE IF NOT EXISTS `usertypes` (
  `TypeId` int NOT NULL,
  `Type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `Description` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  PRIMARY KEY (`TypeId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS `users` (
  `UserId` int NOT NULL AUTO_INCREMENT,
  `UserName` varchar(25) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `Email` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `Type` int DEFAULT NULL,
  PRIMARY KEY (`UserId`) USING BTREE,
  UNIQUE KEY `UNIQUE` (`UserName`) USING BTREE,
  KEY `FK_users_usertypes` (`Type`) USING BTREE,
  CONSTRAINT `FK_users_usertypes` FOREIGN KEY (`Type`) REFERENCES `usertypes` (`TypeId`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS `devices` (
  `DeviceId` int NOT NULL AUTO_INCREMENT,
  `DeviceName` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `Service` varchar(20) DEFAULT NULL,
  `Ip` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `Port` int DEFAULT NULL,
  `OwnerId` int NOT NULL,
  `Status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`DeviceId`) USING BTREE,
  KEY `OwnerId` (`OwnerId`),
  CONSTRAINT `devices_ibfk_1` FOREIGN KEY (`OwnerId`) REFERENCES `users` (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb3;

CREATE TABLE IF NOT EXISTS `views` (
  `UserId` int NOT NULL,
  `DeviceId` int NOT NULL,
  KEY `userId` (`UserId`) USING BTREE,
  KEY `deviceId` (`DeviceId`) USING BTREE,
  CONSTRAINT `views_ibfk_1` FOREIGN KEY (`UserId`) REFERENCES `users` (`UserId`),
  CONSTRAINT `views_ibfk_2` FOREIGN KEY (`DeviceId`) REFERENCES `devices` (`DeviceId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- Inserción de datos

INSERT INTO `usertypes` (`TypeId`, `Type`, `Description`) VALUES
	(1, 'A', 'Admin'),
	(2, 'C', 'Client');

INSERT INTO `users` (`UserId`, `UserName`, `Password`, `Email`, `Type`) VALUES
	(1, 'admin', '0', '0', 1),
	(2, 'daniflo', '0', '0', 2);

INSERT INTO `devices` (`DeviceId`, `DeviceName`, `Service`, `Ip`, `Port`, `OwnerId`, `Status`) VALUES
	(1, 'Canon\\ MG5700\\ series', '_printer._tcp', '[192.168.1.231]', 515, 1, -1),
	(2, 'Canon\\ MG5700\\ series', '_http._tcp', '[192.168.1.231]', 80, 1, -1),
	(3, 'RICOH\\ Aficio\\ MP\\ C305\\ [00267362528A]', '_http._tcp', '[192.168.1.150]', 80, 1, -1),
	(4, 'HP\\ LaserJet\\ Professional\\ M1212nf\\ MFP', '_http._tcp', '[192.168.1.138]', 80, 1, -1),
	(5, 'HP\\ LaserJet\\ Professional\\ M1212nf\\ MFP', '_printer._tcp', '[192.168.1.138]', 515, 1, -1);
	
INSERT INTO `views` (`UserId`, `DeviceId`) VALUES
	(1, 1),
	(1, 2),
	(1, 3),
	(1, 4),
	(1, 5);