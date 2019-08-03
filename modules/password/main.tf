resource "random_string" "password" {
  length  = var.password_length
  special = var.special
}