variable "password_length" {
  type        = number
  description = "The password length"
}

variable "special" {
  type        = bool
  default     = true
  description = "contain special char in password or not"
}