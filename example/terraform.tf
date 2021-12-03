terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.15.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "3.37.0"
    }
  }
}

provider "aws" {
  region = "ap-south-1"
}

resource "docker_image" "example" {
  name = "example"
  build {
    path = "../"
    tag  = ["example:develop"]
  }
}

resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "example1" {
  vpc_id            = aws_vpc.example.id
  availability_zone = "ap-south-1a"
  cidr_block        = "10.0.1.0/24"
}

resource "aws_subnet" "example2" {
  vpc_id            = aws_vpc.example.id
  availability_zone = "ap-south-1b"
  cidr_block        = "10.0.2.0/24"
}

resource "aws_db_subnet_group" "example" {
  name       = "example"
  subnet_ids = [aws_subnet.example1.id, aws_subnet.example2.id]
}

resource "random_password" "example" {
  length           = 24
  special          = true
  override_special = "!#$%^*()-=+_?{}|"
}

# random_password.example.result

# resource "aws_ssm_parameter" "example" {
#   name  = "database-master-password"
#   type  = "SecureString"
#   value = random_password.example.result
# }

resource "aws_rds_cluster" "example" {
  cluster_identifier      = "example"
  engine                  = "aurora-postgresql"
  engine_mode             = "serverless"
  database_name           = "postgres"
  enable_http_endpoint    = false
  master_username         = "root"
  master_password         = "chang333eme321"
  backup_retention_period = 1
  skip_final_snapshot     = true
  db_subnet_group_name    = aws_db_subnet_group.example.name
  # vpc_security_group_ids = [aws_security_group.rds.id]
  # parameter_group_name   = aws_db_parameter_group.education.name

  scaling_configuration {
    auto_pause               = true
    min_capacity             = 2
    max_capacity             = 4
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }
}
