-- MySQL Script generated by MySQL Workbench
-- Tue Jun 19 19:46:11 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema cliserver
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema cliserver
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `cliserver` DEFAULT CHARACTER SET utf8 ;
USE `cliserver` ;

-- -----------------------------------------------------
-- Table `cliserver`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`User` (
  `openId` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `avatar` VARCHAR(100) NOT NULL,
  `phone` VARCHAR(45) NULL,
  PRIMARY KEY (`openId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Marchant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`Marchant` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `hotIndex` INT NULL,
  `introduction` VARCHAR(300) NULL,
  `logo` VARCHAR(100) NULL,
  `images` VARCHAR(3000) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Order`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`Order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date` DATETIME NULL,
  `numOtPeople` INT NULL,
  `deskId` INT NULL,
  `remark` VARCHAR(300) NULL,
  `User_openId` VARCHAR(45) NOT NULL,
  `Marchant_id` INT NOT NULL,
  `status` INT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_Order_User_idx` (`User_openId` ASC),
  INDEX `fk_Order_Marchant1_idx` (`Marchant_id` ASC),
  CONSTRAINT `fk_Order_User`
    FOREIGN KEY (`User_openId`)
    REFERENCES `cliserver`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Order_Marchant1`
    FOREIGN KEY (`Marchant_id`)
    REFERENCES `cliserver`.`Marchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderItem`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`OrderItem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `numOfFoods` INT NULL,
  `Order_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Order1_idx` (`Order_id` ASC),
  CONSTRAINT `fk_OrderItem_Order1`
    FOREIGN KEY (`Order_id`)
    REFERENCES `cliserver`.`Order` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Food`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`Food` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `images` VARCHAR(3000) NULL,
  `type` VARCHAR(45) NULL,
  `price` FLOAT NULL,
  `hotIndex` INT NULL,
  `introduction` VARCHAR(300) NULL,
  `Marchant_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Food_Marchant1_idx` (`Marchant_id` ASC),
  CONSTRAINT `fk_Food_Marchant1`
    FOREIGN KEY (`Marchant_id`)
    REFERENCES `cliserver`.`Marchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderSystem`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`OrderSystem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`Comment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`Comment` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(45) NULL,
  `images` VARCHAR(3000) NULL,
  `User_openId` VARCHAR(45) NOT NULL,
  `Marchant_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Comment_User1_idx` (`User_openId` ASC),
  INDEX `fk_Comment_Marchant1_idx` (`Marchant_id` ASC),
  CONSTRAINT `fk_Comment_User1`
    FOREIGN KEY (`User_openId`)
    REFERENCES `cliserver`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Comment_Marchant1`
    FOREIGN KEY (`Marchant_id`)
    REFERENCES `cliserver`.`Marchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `cliserver`.`OrderItem_Contains_Food`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `cliserver`.`OrderItem_Contains_Food` (
  `id` INT NOT NULL,
  `Food_id` INT NOT NULL,
  `OrderItem_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Contains_Food_Food1_idx` (`Food_id` ASC),
  INDEX `fk_OrderItem_Contains_Food_OrderItem1_idx` (`OrderItem_id` ASC),
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
