CREATE DATABASE  IF NOT EXISTS `recommendationservice` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `recommendationservice`;
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: localhost    Database: recommendationservice
-- ------------------------------------------------------
-- Server version	8.0.22

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
-- Table structure for table `votazione`
--

DROP TABLE IF EXISTS `votazione`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `votazione` (
  `mailUtente` varchar(45) NOT NULL,
  `corsoUtente` varchar(45) NOT NULL,
  `nomeAzienda` varchar(45) NOT NULL,
  `votoUtente` int NOT NULL,
  `votazione` int NOT NULL,
  PRIMARY KEY (`mailUtente`,`nomeAzienda`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `votazione`
--

LOCK TABLES `votazione` WRITE;
/*!40000 ALTER TABLE `votazione` DISABLE KEYS */;
INSERT INTO `votazione` VALUES ('francesco@gmail.com','magistrale','Azienda 1',90,1),('francesco@gmail.com','magistrale','Azienda 200',90,2),('francesco@gmail.com','magistrale','Azienda 26',90,1),('francesco@gmail.com','magistrale','Azienda 300',90,3),('Matteo@gmail.com','triennale','Azienda 200',95,4),('nome100.cognome5050@esempio.com','triennale','Azienda 100',109,3),('nome12.cognome898@esempio.com','magistrale','Azienda 61',97,3),('nome14.cognome6479@esempio.com','magistrale','Azienda 58',61,3),('nome21.cognome5722@esempio.com','triennale','Azienda 40',79,4),('nome25.cognome1816@esempio.com','triennale','Azienda 41',71,5),('nome26.cognome2963@esempio.com','triennale','Azienda 9',74,5),('nome27.cognome6648@esempio.com','triennale','Azienda 100',81,3),('nome28.cognome5277@esempio.com','triennale','Azienda 61',75,3),('nome28.cognome8017@esempio.com','magistrale','Azienda 9',82,1),('nome29.cognome5061@esempio.com','magistrale','Azienda 80',86,3),('nome32.cognome7472@esempio.com','magistrale','Azienda 72',80,2),('nome36.cognome3148@esempio.com','triennale','Azienda 49',82,3),('nome40.cognome1448@esempio.com','magistrale','Azienda 73',109,2),('nome44.cognome3855@esempio.com','magistrale','Azienda 47',90,2),('nome47.cognome1742@esempio.com','triennale','Azienda 92',89,2),('nome47.cognome6267@esempio.com','triennale','Azienda 43',85,3),('nome47.cognome7116@esempio.com','triennale','Azienda 22',93,1),('nome49.cognome3946@esempio.com','magistrale','Azienda 95',66,2),('nome5.cognome1562@esempio.com','triennale','Azienda 56',92,2),('nome5.cognome2824@esempio.com','triennale','Azienda 6',78,1),('nome54.cognome4142@esempio.com','triennale','Azienda 8',104,2),('nome6.cognome457@esempio.com','magistrale','Azienda 41',60,1),('nome60.cognome8337@esempio.com','triennale','Azienda 5',76,1),('nome61.cognome2650@esempio.com','magistrale','Azienda 12',95,3),('nome61.cognome8544@esempio.com','magistrale','Azienda 86',91,5),('nome62.cognome5799@esempio.com','triennale','Azienda 13',72,5),('nome64.cognome9794@esempio.com','triennale','Azienda 65',99,5),('nome66.cognome10099@esempio.com','magistrale','Azienda 61',109,1),('nome66.cognome9998@esempio.com','magistrale','Azienda 24',107,4),('nome68.cognome7150@esempio.com','triennale','Azienda 39',77,5),('nome69.cognome38@esempio.com','triennale','Azienda 88',75,5),('nome71.cognome9878@esempio.com','magistrale','Azienda 8',107,3),('nome72.cognome3034@esempio.com','magistrale','Azienda 41',100,4),('nome72.cognome458@esempio.com','magistrale','Azienda 64',61,3),('nome72.cognome5036@esempio.com','triennale','Azienda 3',74,1),('nome73.cognome7032@esempio.com','triennale','Azienda 82',84,3),('nome75.cognome8811@esempio.com','triennale','Azienda 26',107,4),('nome75.cognome921@esempio.com','triennale','Azienda 57',97,5),('nome76.cognome2283@esempio.com','triennale','Azienda 39',82,1),('nome76.cognome3336@esempio.com','triennale','Azienda 34',99,2),('nome81.cognome8685@esempio.com','triennale','Azienda 36',93,3),('nome82.cognome1633@esempio.com','triennale','Azienda 16',68,3),('nome85.cognome2045@esempio.com','magistrale','Azienda 88',91,5),('nome9.cognome4180@esempio.com','triennale','Azienda 12',97,3),('nome9.cognome5758@esempio.com','triennale','Azienda 95',69,2),('nome9.cognome7759@esempio.com','triennale','Azienda 21',92,1),('nome90.cognome1185@esempio.com','magistrale','Azienda 73',105,2),('nome92.cognome1710@esempio.com','triennale','Azienda 66',110,1),('nome94.cognome3381@esempio.com','triennale','Azienda 98',61,2),('nome97.cognome3584@esempio.com','magistrale','Azienda 62',66,1);
/*!40000 ALTER TABLE `votazione` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'recommendationservice'
--
/*!50003 DROP PROCEDURE IF EXISTS `insert_votazioni_azienda` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `insert_votazioni_azienda`()
BEGIN
  DECLARE i INT DEFAULT 1;
  DECLARE mail_utente VARCHAR(255);
  DECLARE corso_utente VARCHAR(20);
  DECLARE voto_utente INT;
  DECLARE nome_azienda VARCHAR(45);
  DECLARE votazione INT;
  DECLARE random_azienda INT;

  -- Dichiarazione del cursore per scorrere la lista di utenti
  DECLARE utente_cursor CURSOR FOR
    SELECT
      SUBSTRING_INDEX(SUBSTRING_INDEX(
        'francesco@gmail.com 90 magistrale,90uniroma2magistralenome100.cognome5050@esempio.com 109 triennale,109uniroma3triennalenome12.cognome898@esempio.com 97 magistrale,97uniroma3magistralenome14.cognome6479@esempio.com 61 magistrale,61uniroma3magistralenome21.cognome5722@esempio.com 79 triennale,79uniroma1triennalenome25.cognome1816@esempio.com 71 triennale,71uniroma3triennalenome26.cognome2963@esempio.com 74 triennale,74uniroma1triennalenome27.cognome6648@esempio.com 81 triennale,81uniroma1triennalenome28.cognome8017@esempio.com 82 magistrale,82uniroma1magistralenome28.cognome5277@esempio.com 75 triennale,75uniroma1triennalenome29.cognome5061@esempio.com 86 magistrale,86uniroma2magistralenome32.cognome7472@esempio.com 80 magistrale,80uniroma1magistralenome36.cognome3148@esempio.com 82 triennale,82uniroma2triennalenome40.cognome1448@esempio.com 109 magistrale,109uniroma1magistralenome44.cognome3855@esempio.com 90 magistrale,90uniroma3magistralenome47.cognome7116@esempio.com 93 triennale,93uniroma3triennalenome47.cognome1742@esempio.com 89 triennale,89uniroma2triennalenome47.cognome6267@esempio.com 85 triennale,85uniroma1triennalenome49.cognome3946@esempio.com 66 magistrale,66uniroma1magistralenome5.cognome2824@esempio.com 78 triennale,78uniroma2triennalenome5.cognome1562@esempio.com 92 triennale,92uniroma3triennalenome54.cognome4142@esempio.com 104 triennale,104uniroma3triennalenome6.cognome457@esempio.com 60 magistrale,60uniroma1magistralenome60.cognome8337@esempio.com 76 triennale,76uniroma2triennalenome61.cognome2650@esempio.com 95 magistrale,95uniroma3magistralenome61.cognome8544@esempio.com 91 magistrale,91uniroma1magistralenome62.cognome5799@esempio.com 72 triennale,72uniroma2triennalenome64.cognome9794@esempio.com 99 triennale,99uniroma3triennalenome66.cognome9998@esempio.com 107 magistrale,107uniroma3magistralenome66.cognome10099@esempio.com 109 magistrale,109uniroma2magistralenome68.cognome7150@esempio.com 77 triennale,77uniroma1triennalenome69.cognome38@esempio.com 75 triennale,75uniroma2triennalenome71.cognome9878@esempio.com 107 magistrale,107uniroma3magistralenome72.cognome458@esempio.com 61 magistrale,61uniroma3magistralenome72.cognome3034@esempio.com 100 magistrale,100uniroma1magistralenome72.cognome5036@esempio.com 74 triennale,74uniroma1triennalenome73.cognome7032@esempio.com 84 triennale,84uniroma3triennalenome75.cognome8811@esempio.com 107 triennale,107uniroma3triennalenome75.cognome921@esempio.com 97 triennale,97uniroma2triennalenome76.cognome3336@esempio.com 99 triennale,99uniroma3triennalenome76.cognome2283@esempio.com 82 triennale,82uniroma2triennalenome81.cognome8685@esempio.com 93 triennale,93uniroma2triennalenome82.cognome1633@esempio.com 68 triennale,68uniroma3triennalenome85.cognome2045@esempio.com 91 magistrale,91uniroma1magistralenome9.cognome7759@esempio.com 92 triennale,92uniroma1triennalenome9.cognome5758@esempio.com 69 triennale,69uniroma2triennalenome9.cognome4180@esempio.com 97 triennale,97uniroma1triennalenome90.cognome1185@esempio.com 105 magistrale,105uniroma3magistralenome92.cognome1710@esempio.com 110 triennale,110uniroma1triennalenome94.cognome3381@esempio.com 61 triennale,61uniroma3triennalenome97.cognome3584@esempio.com 66 magistrale,66uniroma2magistrale', ',', i), ' ', 1),
      SUBSTRING_INDEX(SUBSTRING_INDEX(
        'francesco@gmail.com 90 magistrale,90uniroma2magistralenome100.cognome5050@esempio.com 109 triennale,109uniroma3triennalenome12.cognome898@esempio.com 97 magistrale,97uniroma3magistralenome14.cognome6479@esempio.com 61 magistrale,61uniroma3magistralenome21.cognome5722@esempio.com 79 triennale,79uniroma1triennalenome25.cognome1816@esempio.com 71 triennale,71uniroma3triennalenome26.cognome2963@esempio.com 74 triennale,74uniroma1triennalenome27.cognome6648@esempio.com 81 triennale,81uniroma1triennalenome28.cognome8017@esempio.com 82 magistrale,82uniroma1magistralenome28.cognome5277@esempio.com 75 triennale,75uniroma1triennalenome29.cognome5061@esempio.com 86 magistrale,86uniroma2magistralenome32.cognome7472@esempio.com 80 magistrale,80uniroma1magistralenome36.cognome3148@esempio.com 82 triennale,82uniroma2triennalenome40.cognome1448@esempio.com 109 magistrale,109uniroma1magistralenome44.cognome3855@esempio.com 90 magistrale,90uniroma3magistralenome47.cognome7116@esempio.com 93 triennale,93uniroma3triennalenome47.cognome1742@esempio.com 89 triennale,89uniroma2triennalenome47.cognome6267@esempio.com 85 triennale,85uniroma1triennalenome49.cognome3946@esempio.com 66 magistrale,66uniroma1magistralenome5.cognome2824@esempio.com 78 triennale,78uniroma2triennalenome5.cognome1562@esempio.com 92 triennale,92uniroma3triennalenome54.cognome4142@esempio.com 104 triennale,104uniroma3triennalenome6.cognome457@esempio.com 60 magistrale,60uniroma1magistralenome60.cognome8337@esempio.com 76 triennale,76uniroma2triennalenome61.cognome2650@esempio.com 95 magistrale,95uniroma3magistralenome61.cognome8544@esempio.com 91 magistrale,91uniroma1magistralenome62.cognome5799@esempio.com 72 triennale,72uniroma2triennalenome64.cognome9794@esempio.com 99 triennale,99uniroma3triennalenome66.cognome9998@esempio.com 107 magistrale,107uniroma3magistralenome66.cognome10099@esempio.com 109 magistrale,109uniroma2magistralenome68.cognome7150@esempio.com 77 triennale,77uniroma1triennalenome69.cognome38@esempio.com 75 triennale,75uniroma2triennalenome71.cognome9878@esempio.com 107 magistrale,107uniroma3magistralenome72.cognome458@esempio.com 61 magistrale,61uniroma3magistralenome72.cognome3034@esempio.com 100 magistrale,100uniroma1magistralenome72.cognome5036@esempio.com 74 triennale,74uniroma1triennalenome73.cognome7032@esempio.com 84 triennale,84uniroma3triennalenome75.cognome8811@esempio.com 107 triennale,107uniroma3triennalenome75.cognome921@esempio.com 97 triennale,97uniroma2triennalenome76.cognome3336@esempio.com 99 triennale,99uniroma3triennalenome76.cognome2283@esempio.com 82 triennale,82uniroma2triennalenome81.cognome8685@esempio.com 93 triennale,93uniroma2triennalenome82.cognome1633@esempio.com 68 triennale,68uniroma3triennalenome85.cognome2045@esempio.com 91 magistrale,91uniroma1magistralenome9.cognome7759@esempio.com 92 triennale,92uniroma1triennalenome9.cognome5758@esempio.com 69 triennale,69uniroma2triennalenome9.cognome4180@esempio.com 97 triennale,97uniroma1triennalenome90.cognome1185@esempio.com 105 magistrale,105uniroma3magistralenome92.cognome1710@esempio.com 110 triennale,110uniroma1triennalenome94.cognome3381@esempio.com 61 triennale,61uniroma3triennalenome97.cognome3584@esempio.com 66 magistrale,66uniroma2magistrale', ',', i), ' ', 2)
    FROM
      information_schema.COLUMNS
    WHERE
      TABLE_SCHEMA = DATABASE()
    LIMIT 50;-- Assumendo 50 utenti

  -- Gestione dell'errore se il cursore non trova dati
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET @done = TRUE;

  OPEN utente_cursor;
  read_loop: LOOP
    FETCH utente_cursor INTO mail_utente, voto_utente, corso_utente;
    IF @done THEN
      LEAVE read_loop;
    END IF;

    -- Genera un'azienda casuale
    SET random_azienda = FLOOR(1 + RAND() * 100);
    SET nome_azienda = CONCAT('Azienda ', random_azienda);

    -- Verifica se l'utente ha già votato per questa azienda
    IF NOT EXISTS (SELECT 1 FROM votazione WHERE mailUtente = mail_utente AND nomeAzienda = nome_azienda) THEN
      -- Genera una votazione casuale
      SET votazione = FLOOR(1 + RAND() * 5);

      -- Inserisci la votazione nella tabella
      INSERT INTO votazione (mailUtente, corsoUtente, nomeAzienda, votoUtente, votazione)
      VALUES (mail_utente, corso_utente, nome_azienda, voto_utente, votazione);
    END IF;

    SET i = i + 1;
  END LOOP;

  CLOSE utente_cursor;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `insert_votazioni_univoche_azienda` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `insert_votazioni_univoche_azienda`()
BEGIN
  DECLARE i INT DEFAULT 1;
  DECLARE mail_utente VARCHAR(255);
  DECLARE voto_utente INT;
  DECLARE tipo_laurea_utente VARCHAR(20);
  DECLARE nome_azienda VARCHAR(45);
  DECLARE votazione INT;
  DECLARE random_azienda INT;
  DECLARE done BOOLEAN DEFAULT FALSE;

  -- Dichiarazione del cursore per scorrere la lista di utenti
  DECLARE utente_cursor CURSOR FOR
    SELECT 
      -- Estrae la mail dell'utente
      SUBSTRING_INDEX(user_data, ' ', 1) AS mail_utente,
      -- Estrae il voto dell'utente
      CAST(SUBSTRING_INDEX(SUBSTRING_INDEX(user_data, ' ', 2), ' ', -1) AS UNSIGNED) AS voto_utente,
      -- Estrae il tipo di laurea dell'utente
      TRIM(SUBSTRING_INDEX(user_data, ' ', -1)) AS tipo_laurea_utente
    FROM (
      SELECT 'francesco@gmail.com 90 magistrale' AS user_data
      UNION ALL
      SELECT 'nome100.cognome5050@esempio.com 109 triennale'
      UNION ALL
      SELECT 'nome12.cognome898@esempio.com 97 magistrale'
      UNION ALL
      SELECT 'nome14.cognome6479@esempio.com 61 magistrale'
      UNION ALL
      SELECT 'nome21.cognome5722@esempio.com 79 triennale'
      UNION ALL
      SELECT 'nome25.cognome1816@esempio.com 71 triennale'
      UNION ALL
      SELECT 'nome26.cognome2963@esempio.com 74 triennale'
      UNION ALL
      SELECT 'nome27.cognome6648@esempio.com 81 triennale'
      UNION ALL
      SELECT 'nome28.cognome8017@esempio.com 82 magistrale'
      UNION ALL
      SELECT 'nome28.cognome5277@esempio.com 75 triennale'
      UNION ALL
      SELECT 'nome29.cognome5061@esempio.com 86 magistrale'
      UNION ALL
      SELECT 'nome32.cognome7472@esempio.com 80 magistrale'
      UNION ALL
      SELECT 'nome36.cognome3148@esempio.com 82 triennale'
      UNION ALL
      SELECT 'nome40.cognome1448@esempio.com 109 magistrale'
      UNION ALL
      SELECT 'nome44.cognome3855@esempio.com 90 magistrale'
      UNION ALL
      SELECT 'nome47.cognome7116@esempio.com 93 triennale'
      UNION ALL
      SELECT 'nome47.cognome1742@esempio.com 89 triennale'
      UNION ALL
      SELECT 'nome47.cognome6267@esempio.com 85 triennale'
      UNION ALL
      SELECT 'nome49.cognome3946@esempio.com 66 magistrale'
      UNION ALL
      SELECT 'nome5.cognome2824@esempio.com 78 triennale'
      UNION ALL
      SELECT 'nome5.cognome1562@esempio.com 92 triennale'
      UNION ALL
      SELECT 'nome54.cognome4142@esempio.com 104 triennale'
      UNION ALL
      SELECT 'nome6.cognome457@esempio.com 60 magistrale'
      UNION ALL
      SELECT 'nome60.cognome8337@esempio.com 76 triennale'
      UNION ALL
      SELECT 'nome61.cognome2650@esempio.com 95 magistrale'
      UNION ALL
      SELECT 'nome61.cognome8544@esempio.com 91 magistrale'
      UNION ALL
      SELECT 'nome62.cognome5799@esempio.com 72 triennale'
      UNION ALL
      SELECT 'nome64.cognome9794@esempio.com 99 triennale'
      UNION ALL
      SELECT 'nome66.cognome9998@esempio.com 107 magistrale'
      UNION ALL
      SELECT 'nome66.cognome10099@esempio.com 109 magistrale'
      UNION ALL
      SELECT 'nome68.cognome7150@esempio.com 77 triennale'
      UNION ALL
      SELECT 'nome69.cognome38@esempio.com 75 triennale'
      UNION ALL
      SELECT 'nome71.cognome9878@esempio.com 107 magistrale'
      UNION ALL
      SELECT 'nome72.cognome458@esempio.com 61 magistrale'
      UNION ALL
      SELECT 'nome72.cognome3034@esempio.com 100 magistrale'
      UNION ALL
      SELECT 'nome72.cognome5036@esempio.com 74 triennale'
      UNION ALL
      SELECT 'nome73.cognome7032@esempio.com 84 triennale'
      UNION ALL
      SELECT 'nome75.cognome8811@esempio.com 107 triennale'
      UNION ALL
      SELECT 'nome75.cognome921@esempio.com 97 triennale'
      UNION ALL
      SELECT 'nome76.cognome3336@esempio.com 99 triennale'
      UNION ALL
      SELECT 'nome76.cognome2283@esempio.com 82 triennale'
      UNION ALL
      SELECT 'nome81.cognome8685@esempio.com 93 triennale'
      UNION ALL
      SELECT 'nome82.cognome1633@esempio.com 68 triennale'
      UNION ALL
      SELECT 'nome85.cognome2045@esempio.com 91 magistrale'
      UNION ALL
      SELECT 'nome9.cognome7759@esempio.com 92 triennale'
      UNION ALL
      SELECT 'nome9.cognome5758@esempio.com 69 triennale'
      UNION ALL
      SELECT 'nome9.cognome4180@esempio.com 97 triennale'
      UNION ALL
      SELECT 'nome90.cognome1185@esempio.com 105 magistrale'
      UNION ALL
      SELECT 'nome92.cognome1710@esempio.com 110 triennale'
      UNION ALL
      SELECT 'nome94.cognome3381@esempio.com 61 triennale'
      UNION ALL
      SELECT 'nome97.cognome3584@esempio.com 66 magistrale'
    ) AS utenti;

  -- Gestione dell'errore se il cursore non trova dati
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

  OPEN utente_cursor;
  read_loop: LOOP
    FETCH utente_cursor INTO mail_utente, voto_utente, tipo_laurea_utente;
    IF done THEN
      LEAVE read_loop;
    END IF;

    -- Genera un'azienda casuale
    SET random_azienda = FLOOR(1 + RAND() * 100);
    SET nome_azienda = CONCAT('Azienda ', random_azienda);

    -- Verifica se l'utente ha già votato per questa azienda
    IF NOT EXISTS (SELECT 1 FROM votazione WHERE mailUtente = mail_utente AND nomeAzienda = nome_azienda) THEN
      -- Genera una votazione casuale
      SET votazione = FLOOR(1 + RAND() * 5);

      -- Inserisci la votazione nella tabella
      INSERT INTO votazione (mailUtente, corsoUtente, nomeAzienda, votoUtente, votazione)
      VALUES (mail_utente, tipo_laurea_utente, nome_azienda, voto_utente, votazione);
    END IF;

    SET i = i + 1;
  END LOOP;

  CLOSE utente_cursor;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-29 15:34:44
