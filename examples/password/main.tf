module "random_password" {
  source = "../..//modules/password"
  password_length = var.password_length
}
