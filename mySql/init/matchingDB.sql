CREATE DATABASE  IF NOT EXISTS `matchingservice` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `matchingservice`;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: matchingservice
-- ------------------------------------------------------
-- Server version	8.0.42

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `aziende`
--

DROP TABLE IF EXISTS `aziende`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `aziende` (
  `nomeAzienda` varchar(45) NOT NULL,
  `ambitoRichiesto` varchar(45) NOT NULL,
  PRIMARY KEY (`nomeAzienda`,`ambitoRichiesto`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `aziende`
--

LOCK TABLES `aziende` WRITE;
/*!40000 ALTER TABLE `aziende` DISABLE KEYS */;
INSERT INTO `aziende` VALUES ('Azienda 1','Cybersecurity'),('Azienda 1','Data Science'),('Azienda 1','Software'),('Azienda 10','Cybersecurity'),('Azienda 100','Cybersecurity'),('Azienda 11','Cybersecurity'),('Azienda 12','Cybersecurity'),('Azienda 13','Software'),('Azienda 14','Cybersecurity'),('Azienda 15','Cybersecurity'),('Azienda 16','Data Science'),('Azienda 17','Software'),('Azienda 18','Software'),('Azienda 19','Cybersecurity'),('Azienda 2','Cybersecurity'),('Azienda 20','Data Science'),('Azienda 200','DataScience'),('Azienda 21','Data Science'),('Azienda 22','Cybersecurity'),('Azienda 23','Software'),('Azienda 24','Data Science'),('Azienda 25','Data Science'),('Azienda 26','Data Science'),('Azienda 27','Software'),('Azienda 28','Data Science'),('Azienda 29','Software'),('Azienda 3','Data Science'),('Azienda 30','Software'),('Azienda 300','DataScience'),('Azienda 300','Software'),('Azienda 31','Cybersecurity'),('Azienda 32','Data Science'),('Azienda 33','Software'),('Azienda 34','Data Science'),('Azienda 35','Cybersecurity'),('Azienda 36','Cybersecurity'),('Azienda 37','Data Science'),('Azienda 38','Data Science'),('Azienda 39','Data Science'),('Azienda 4','Software'),('Azienda 40','Data Science'),('Azienda 41','Data Science'),('Azienda 42','Cybersecurity'),('Azienda 43','Software'),('Azienda 44','Cybersecurity'),('Azienda 45','Data Science'),('Azienda 46','Cybersecurity'),('Azienda 47','Data Science'),('Azienda 48','Cybersecurity'),('Azienda 49','Software'),('Azienda 5','Cybersecurity'),('Azienda 50','Software'),('Azienda 51','Software'),('Azienda 52','Software'),('Azienda 53','Software'),('Azienda 54','Software'),('Azienda 55','Software'),('Azienda 56','Software'),('Azienda 57','Software'),('Azienda 58','Data Science'),('Azienda 59','Software'),('Azienda 6','Software'),('Azienda 60','Software'),('Azienda 61','Cybersecurity'),('Azienda 62','Software'),('Azienda 63','Cybersecurity'),('Azienda 64','Cybersecurity'),('Azienda 65','Data Science'),('Azienda 66','Data Science'),('Azienda 67','Software'),('Azienda 68','Data Science'),('Azienda 69','Cybersecurity'),('Azienda 7','Software'),('Azienda 70','Cybersecurity'),('Azienda 71','Software'),('Azienda 72','Cybersecurity'),('Azienda 73','Software'),('Azienda 74','Software'),('Azienda 75','Cybersecurity'),('Azienda 76','Software'),('Azienda 77','Data Science'),('Azienda 78','Cybersecurity'),('Azienda 79','Cybersecurity'),('Azienda 8','Cybersecurity'),('Azienda 80','Cybersecurity'),('Azienda 81','Cybersecurity'),('Azienda 82','Cybersecurity'),('Azienda 83','Data Science'),('Azienda 84','Data Science'),('Azienda 85','Software'),('Azienda 86','Software'),('Azienda 87','Software'),('Azienda 88','Data Science'),('Azienda 89','Data Science'),('Azienda 9','Data Science'),('Azienda 90','Data Science'),('Azienda 91','Cybersecurity'),('Azienda 92','Data Science'),('Azienda 93','Data Science'),('Azienda 94','Software'),('Azienda 95','Data Science'),('Azienda 96','Cybersecurity'),('Azienda 97','Cybersecurity'),('Azienda 98','Data Science'),('Azienda 99','Software'),('Azienda101','DataScience');
/*!40000 ALTER TABLE `aziende` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `datiutente`
--

DROP TABLE IF EXISTS `datiutente`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `datiutente` (
  `nomeUtente` varchar(45) NOT NULL,
  `cognomeUtente` varchar(45) NOT NULL,
  `emailUtente` varchar(45) NOT NULL,
  `votoDiLaureaUtente` int NOT NULL,
  `universitàUtente` varchar(45) NOT NULL,
  `tipoDiLaureaUtente` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `etaUtente` int NOT NULL,
  `preferenza1` varchar(45) DEFAULT NULL,
  `preferenza2` varchar(45) DEFAULT NULL,
  `preferenza3` varchar(45) DEFAULT NULL,
  `ambitoAssegnato` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`emailUtente`),
  UNIQUE KEY `emailUtente_UNIQUE` (`emailUtente`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `datiutente`
--

LOCK TABLES `datiutente` WRITE;
/*!40000 ALTER TABLE `datiutente` DISABLE KEYS */;
INSERT INTO `datiutente` VALUES ('Alessandro','Rossi','alessandro@gmail.com',90,'uniroma2','triennale','Alex.80',24,NULL,NULL,NULL,NULL),('Francesco','Verdi','francesco@gmail.com',90,'uniroma2','magistrale','Ciccio.80',25,'DataScience','Software','Cybersecurity',''),('Matteo','Gialli','Matteo@gmail.com',95,'uniroma3','triennale','Matteo.80',23,'Cybersecurity','DataScience','Software',''),('Nome100','Cognome50','nome100.cognome5050@esempio.com',109,'uniroma3','triennale','I@Mc(,D_nx{3S1V[',25,'Data Science','Software','Cybersecurity',NULL),('Nome12','Cognome89','nome12.cognome898@esempio.com',97,'uniroma3','magistrale','_VosWme-,A3yPUv2',25,'Software','Data Science','Cybersecurity',NULL),('Nome14','Cognome64','nome14.cognome6479@esempio.com',61,'uniroma3','magistrale','JH8AMApaPFH_b1,:',28,'Data Science','Software','Cybersecurity',NULL),('Nome21','Cognome57','nome21.cognome5722@esempio.com',79,'uniroma1','triennale','{R(SmrYy7W@.4C8Q',24,'Cybersecurity','Software','Data Science',NULL),('Nome25','Cognome18','nome25.cognome1816@esempio.com',71,'uniroma3','triennale','=i:KwbBT34yQ0S&D',28,'Software','Data Science','Cybersecurity',NULL),('Nome26','Cognome29','nome26.cognome2963@esempio.com',74,'uniroma1','triennale',']tN01z2tx7Y=TbV4',27,'Software','Cybersecurity','Data Science',NULL),('Nome27','Cognome66','nome27.cognome6648@esempio.com',81,'uniroma1','triennale','X:dF5tl?X_R>Mu_P',28,'Data Science','Software','Cybersecurity',NULL),('Nome28','Cognome52','nome28.cognome5277@esempio.com',75,'uniroma1','triennale','M!ucOs8*{i&3_w^h',23,'Cybersecurity','Software','Data Science',NULL),('Nome28','Cognome80','nome28.cognome8017@esempio.com',82,'uniroma1','magistrale','{LRrR[=5!1]0B(f}',28,'Cybersecurity','Data Science','Software',NULL),('Nome29','Cognome50','nome29.cognome5061@esempio.com',86,'uniroma2','magistrale','_XD>}Ynf=f$THG6u',29,'Cybersecurity','Software','Data Science',NULL),('Nome32','Cognome74','nome32.cognome7472@esempio.com',80,'uniroma1','magistrale','$CI?*,IkM*7>]P@r',29,'Data Science','Software','Cybersecurity',NULL),('Nome36','Cognome31','nome36.cognome3148@esempio.com',82,'uniroma2','triennale','!kN}mb7CY{#?(?0<',28,'Cybersecurity','Software','Data Science',NULL),('Nome40','Cognome14','nome40.cognome1448@esempio.com',109,'uniroma1','magistrale','rHDT0MI3?hN?0?sg',25,'Data Science','Software','Cybersecurity',NULL),('Nome44','Cognome38','nome44.cognome3855@esempio.com',90,'uniroma3','magistrale','A4D@2;$>7Ja(c9K?',25,'Data Science','Software','Cybersecurity',NULL),('Nome47','Cognome17','nome47.cognome1742@esempio.com',89,'uniroma2','triennale','cDZ;+Vj:D]Hxtz^>',27,'Software','Cybersecurity','Data Science',NULL),('Nome47','Cognome62','nome47.cognome6267@esempio.com',85,'uniroma1','triennale','*__5%-=&DF&-%ENq',26,'Software','Data Science','Cybersecurity',NULL),('Nome47','Cognome71','nome47.cognome7116@esempio.com',93,'uniroma3','triennale','{O17U43trFr|,bwv',28,'Software','Cybersecurity','Data Science',NULL),('Nome49','Cognome39','nome49.cognome3946@esempio.com',66,'uniroma1','magistrale','L)&N:D=tZxZj^VJG',24,'Software','Cybersecurity','Data Science',NULL),('Nome5','Cognome15','nome5.cognome1562@esempio.com',92,'uniroma3','triennale','w(v)EzIkI2}78H[q',25,'Cybersecurity','Software','Data Science',NULL),('Nome5','Cognome28','nome5.cognome2824@esempio.com',78,'uniroma2','triennale','?!W1Jp9<$*$CJd?=',23,'Cybersecurity','Data Science','Software',NULL),('Nome54','Cognome41','nome54.cognome4142@esempio.com',104,'uniroma3','triennale','XEciJ!G!OmBp:m;r',23,'Cybersecurity','Software','Data Science',NULL),('Nome6','Cognome45','nome6.cognome457@esempio.com',60,'uniroma1','magistrale','l>V#b;Yf0)z<cq[&',29,'Cybersecurity','Software','Data Science',NULL),('Nome60','Cognome83','nome60.cognome8337@esempio.com',76,'uniroma2','triennale',']O4+Gvm.Iht=Xz#_',26,'Data Science','Software','Cybersecurity',NULL),('Nome61','Cognome26','nome61.cognome2650@esempio.com',95,'uniroma3','magistrale','[{[58Ns9+sO8fvdO',30,'Cybersecurity','Data Science','Software',NULL),('Nome61','Cognome85','nome61.cognome8544@esempio.com',91,'uniroma1','magistrale','tS+(IZ#|Hm0XbG_d',28,'Data Science','Software','Cybersecurity',NULL),('Nome62','Cognome57','nome62.cognome5799@esempio.com',72,'uniroma2','triennale','%^Z5PIWYvQ8>]LPj',24,'Cybersecurity','Software','Data Science',NULL),('Nome64','Cognome97','nome64.cognome9794@esempio.com',99,'uniroma3','triennale','S]+6^(#yoaTYD.&c',22,'Data Science','Software','Cybersecurity',NULL),('Nome66','Cognome100','nome66.cognome10099@esempio.com',109,'uniroma2','magistrale','bE9NoO(4&<S5#8ui',29,'Data Science','Software','Cybersecurity',NULL),('Nome66','Cognome99','nome66.cognome9998@esempio.com',107,'uniroma3','magistrale',',k%J%Z4OIZ@]tO6}',28,'Software','Cybersecurity','Data Science',NULL),('Nome68','Cognome71','nome68.cognome7150@esempio.com',77,'uniroma1','triennale','^?6C6E!VWlc$2+KU',24,'Software','Cybersecurity','Data Science',NULL),('Nome69','Cognome3','nome69.cognome38@esempio.com',75,'uniroma2','triennale',':A9Y&pXy7Y_N$I8A',24,'Software','Cybersecurity','Data Science',NULL),('Nome71','Cognome98','nome71.cognome9878@esempio.com',107,'uniroma3','magistrale','Mu-O-%xj@ytw4Lu=',25,'Software','Cybersecurity','Data Science',NULL),('Nome72','Cognome30','nome72.cognome3034@esempio.com',100,'uniroma1','magistrale','C8QGH){dTR],j5>,',30,'Data Science','Software','Cybersecurity',NULL),('Nome72','Cognome45','nome72.cognome458@esempio.com',61,'uniroma3','magistrale','k5;6U59TP-!enZS#',30,'Cybersecurity','Data Science','Software',NULL),('Nome72','Cognome50','nome72.cognome5036@esempio.com',74,'uniroma1','triennale','GUUg%XXi-(VB|4QS',25,'Cybersecurity','Data Science','Software',NULL),('Nome73','Cognome70','nome73.cognome7032@esempio.com',84,'uniroma3','triennale','3zU(J1-Ck4+Dl5[G',26,'Cybersecurity','Software','Data Science',NULL),('Nome75','Cognome88','nome75.cognome8811@esempio.com',107,'uniroma3','triennale','qVlb#Y9$3;@_bZ[0',24,'Software','Data Science','Cybersecurity',NULL),('Nome75','Cognome9','nome75.cognome921@esempio.com',97,'uniroma2','triennale','!TN9n7|WcOv-IQt8',23,'Data Science','Software','Cybersecurity',NULL),('Nome76','Cognome22','nome76.cognome2283@esempio.com',82,'uniroma2','triennale','&ilFLe?)e_bY(Bl#',28,'Cybersecurity','Software','Data Science',NULL),('Nome76','Cognome33','nome76.cognome3336@esempio.com',99,'uniroma3','triennale','NL&2%;HkP<PN*4-s',29,'Software','Cybersecurity','Data Science',NULL),('Nome81','Cognome86','nome81.cognome8685@esempio.com',93,'uniroma2','triennale','H!La#4bm7:6V6$8v',28,'Cybersecurity','Data Science','Software',NULL),('Nome82','Cognome16','nome82.cognome1633@esempio.com',68,'uniroma3','triennale','u;:}@[tO7<=DjUCc',29,'Cybersecurity','Data Science','Software',NULL),('Nome85','Cognome20','nome85.cognome2045@esempio.com',91,'uniroma1','magistrale','$#Ohd,3B8PCsf7s?',29,'Software','Data Science','Cybersecurity',NULL),('Nome9','Cognome41','nome9.cognome4180@esempio.com',97,'uniroma1','triennale','FSM8kUB:&his%nR?',24,'Software','Cybersecurity','Data Science',NULL),('Nome9','Cognome57','nome9.cognome5758@esempio.com',69,'uniroma2','triennale','=v6Y}^lFK<2pc2an',24,'Data Science','Software','Cybersecurity',NULL),('Nome9','Cognome77','nome9.cognome7759@esempio.com',92,'uniroma1','triennale','MmGNmEzNEKb[x@-.',25,'Software','Data Science','Cybersecurity',NULL),('Nome90','Cognome11','nome90.cognome1185@esempio.com',105,'uniroma3','magistrale',':-Cl7?,$}y5Nzi3_',30,'Cybersecurity','Software','Data Science',NULL),('Nome92','Cognome17','nome92.cognome1710@esempio.com',110,'uniroma1','triennale','M)%EOzcAKoZLI7r>',27,'Cybersecurity','Data Science','Software',NULL),('Nome94','Cognome33','nome94.cognome3381@esempio.com',61,'uniroma3','triennale','pj?7I|Is{)ITIL.V',28,'Data Science','Software','Cybersecurity',NULL),('Nome97','Cognome35','nome97.cognome3584@esempio.com',66,'uniroma2','magistrale','BE(>QRb2eG!K<1k*',23,'Software','Data Science','Cybersecurity',NULL),('Filippo','Neri','pippo@gmail.com',98,'uniroma1','magistrale','Pippo.80',26,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `datiutente` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'matchingservice'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-15 23:55:54
