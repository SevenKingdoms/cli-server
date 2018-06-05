-- MySQL Script generated by MySQL Workbench
-- Tue Jun  5 15:21:02 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema clidb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema clidb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `clidb` DEFAULT CHARACTER SET utf8 ;
USE `clidb` ;

-- -----------------------------------------------------
-- Table `clidb`.`User`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`User` (
  `openId` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `avatar` VARCHAR(100) NOT NULL,
  `phone` VARCHAR(45) NULL,
  PRIMARY KEY (`openId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`Order`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`Order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `date` DATETIME NULL,
  `numOtPeople` INT NULL,
  `deskId` INT NULL,
  `remark` VARCHAR(300) NULL,
  `User_openId` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_Order_User_idx` (`User_openId` ASC),
  CONSTRAINT `fk_Order_User`
    FOREIGN KEY (`User_openId`)
    REFERENCES `clidb`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`OrderItem`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`OrderItem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `numOfFoods` INT NULL,
  `Order_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Order1_idx` (`Order_id` ASC),
  CONSTRAINT `fk_OrderItem_Order1`
    FOREIGN KEY (`Order_id`)
    REFERENCES `clidb`.`Order` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`Marchant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`Marchant` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `hotIndex` INT NULL,
  `introduction` VARCHAR(300) NULL,
  `logo` VARCHAR(100) NULL,
  `images` VARCHAR(3000) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`Food`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`Food` (
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
    REFERENCES `clidb`.`Marchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`OrderSystem`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`OrderSystem` (
  `id` INT NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`Comment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`Comment` (
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
    REFERENCES `clidb`.`User` (`openId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Comment_Marchant1`
    FOREIGN KEY (`Marchant_id`)
    REFERENCES `clidb`.`Marchant` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clidb`.`OrderItem_Contains_Food`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clidb`.`OrderItem_Contains_Food` (
  `id` INT NOT NULL,
  `OrderItem_id` INT NOT NULL,
  `Food_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_OrderItem_Contains_Food_OrderItem1_idx` (`OrderItem_id` ASC),
  INDEX `fk_OrderItem_Contains_Food_Food1_idx` (`Food_id` ASC),
  CONSTRAINT `fk_OrderItem_Contains_Food_OrderItem1`
    FOREIGN KEY (`OrderItem_id`)
    REFERENCES `clidb`.`OrderItem` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_OrderItem_Contains_Food_Food1`
    FOREIGN KEY (`Food_id`)
    REFERENCES `clidb`.`Food` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
