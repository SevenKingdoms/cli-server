-- MySQL Script generated by MySQL Workbench
-- Fri Jun 29 16:08:10 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema cliserver
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `cliserver` ;

-- -----------------------------------------------------
-- Schema cliserver
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `cliserver` DEFAULT CHARACTER SET utf8 ;
USE `cliserver` ;

-- -----------------------------------------------------
-- Table `cliserver`.`User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`User` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`User` (
  `openId` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `avatar` VARCHAR(200) NOT NULL,
  `phone` VARCHAR(45) NULL,
  PRIMARY KEY (`openId`),
  UNIQUE INDEX `openId_UNIQUE` (`openId` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Merchant`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`Merchant` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`Merchant` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `tel` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `hotIndex` INT NULL,
  `introduction` VARCHAR(300) NULL,
  `logo` VARCHAR(100) NULL,
  `images` VARCHAR(3000) NULL,
  `open` TINYINT NULL,
  `openTime` VARCHAR(150) NULL,
  `score` FLOAT NULL,
  `address` VARCHAR(200) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `account_UNIQUE` (`tel` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Order`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`Order` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`Order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date` DATETIME NULL,
  `numOfPeople` INT NULL,
  `deskId` INT NULL,
  `remark` VARCHAR(1000) NULL,
  `paid` TINYINT NULL,
  `User_openId` VARCHAR(45) NOT NULL,
  `Merchant_id` INT NOT NULL,
  `foods` VARCHAR(1000) NULL,
  `create_at` DATETIME NULL,
  `merchant_name` VARCHAR(45) NULL,
  `merchant_tel` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_Order_User_idx` (`User_openId` ASC),
  INDEX `fk_Order_Marchant1_idx` (`Merchant_id` ASC),
  CONSTRAINT `fk_Order_User`
    FOREIGN KEY (`User_openId`)
    REFERENCES `cliserver`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Order_Marchant1`
    FOREIGN KEY (`Merchant_id`)
    REFERENCES `cliserver`.`Merchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderItem`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`OrderItem` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`OrderItem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `numOfFoods` INT NULL,
  `Order_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Order1_idx` (`Order_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `fk_OrderItem_Order1`
    FOREIGN KEY (`Order_id`)
    REFERENCES `cliserver`.`Order` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Food`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`Food` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`Food` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `image` VARCHAR(300) NULL,
  `type` VARCHAR(300) NULL,
  `price` FLOAT NULL,
  `hotIndex` INT NULL,
  `introduction` VARCHAR(300) NULL,
  `Merchant_id` INT NOT NULL,
  `inStock` TINYINT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Food_Marchant1_idx` (`Merchant_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `fk_Food_Marchant1`
    FOREIGN KEY (`Merchant_id`)
    REFERENCES `cliserver`.`Merchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderSystem`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`OrderSystem` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`OrderSystem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Comment`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`Comment` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`Comment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(45) NULL,
  `images` VARCHAR(3000) NULL,
  `User_openId` VARCHAR(45) NOT NULL,
  `Merchant_id` INT NOT NULL,
  `score` FLOAT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Comment_User1_idx` (`User_openId` ASC),
  INDEX `fk_Comment_Marchant1_idx` (`Merchant_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `fk_Comment_User1`
    FOREIGN KEY (`User_openId`)
    REFERENCES `cliserver`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Comment_Marchant1`
    FOREIGN KEY (`Merchant_id`)
    REFERENCES `cliserver`.`Merchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderItem_Contains_Food`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `cliserver`.`OrderItem_Contains_Food` ;

CREATE TABLE IF NOT EXISTS `cliserver`.`OrderItem_Contains_Food` (
  `id` INT NOT NULL,
  `Food_id` INT NOT NULL,
  `OrderItem_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Contains_Food_Food1_idx` (`Food_id` ASC),
  INDEX `fk_OrderItem_Contains_Food_OrderItem1_idx` (`OrderItem_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  CONSTRAINT `fk_OrderItem_Contains_Food_Food1`
    FOREIGN KEY (`Food_id`)
    REFERENCES `cliserver`.`Food` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_OrderItem_Contains_Food_OrderItem1`
    FOREIGN KEY (`OrderItem_id`)
    REFERENCES `cliserver`.`OrderItem` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
